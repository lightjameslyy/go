package insar

import (
	"bytes"
	"fmt"
	alluxio "github.com/Alluxio/alluxio-go"
	"github.com/Alluxio/alluxio-go/option"
	"io/ioutil"
	"log"
	"math"
	"math/cmplx"
	"runtime"
	"time"
	"unsafe"
)

func getBuffer(client *alluxio.Client, path string) ([]byte, error) {
	id, err := client.OpenFile(path, &option.OpenFile{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close(id)
	readCloser, err := client.Read(id)
	if err != nil {
		log.Fatal(err)
	}
	defer readCloser.Close()
	return ioutil.ReadAll(readCloser)
}

func CompareComplex64(ip string, path1 string, path2 string, counts int) {
	fs := alluxio.NewClient(ip, 39999, time.Minute)
	buffer1, err := getBuffer(fs, path1)
	if err != nil {
		log.Fatal(err)
	}
	data1 := *(*[]complex64)(unsafe.Pointer(&buffer1))
	len1 := len(data1)
	if counts > len1 {
		log.Fatal("counts larger than file1 length!")
	}

	buffer2, err := getBuffer(fs, path2)
	if err != nil {
		log.Fatal(err)
	}
	data2 := *(*[]complex64)(unsafe.Pointer(&buffer2))
	len2 := len(data2)
	if counts > len2 {
		log.Fatal("counts larger than file2 length!")
	}
	if len1 != len2 {
		log.Fatal("file length different!")
	}
	fmt.Println(len1, len2)
	for i := 0; i < len1; i++ {
		/*
			if i < 67108870 {
				continue
			}
				if i < counts {
					fmt.Println(data1[i], data2[i])
				}
		*/
		if i%1000000 == 0 {
			fmt.Println(i)
		}
		diff := data1[i] - data2[i]
		if math.Abs(float64(real(diff))) > 10e-6 || math.Abs(float64(imag(diff))) > 10e-6 {
			fmt.Println(i, data1[i], data2[i])
		}
	}
}

func CoherenceFun(ip string, inM string, inS string, inFlat string, outCoh string, rLooks uint32, cLooks uint32, iCols uint32, iRows uint32) {
	// get alluxio client
	fs := alluxio.NewClient(ip, 39999, time.Minute)

	// get input buffers
	MBytes, err := getBuffer(fs, inM)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(MBytes))

	SBytes, err := getBuffer(fs, inS)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(SBytes))

	FlatBytes, err := getBuffer(fs, inFlat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(FlatBytes))

	ok, err := fs.Exists(outCoh, &option.Exists{})
	if err != nil {
		log.Fatal(err)
	}
	if true == ok {
		err := fs.Delete(outCoh, &option.Delete{})
		if err != nil {
			log.Fatal(err)
		}
	}
	outCohId, err := fs.CreateFile(outCoh, &option.CreateFile{})
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close(outCohId)
	/*
		var message = []byte("Greetings traveller!")
		n, err := fs.Write(outCohId, bytes.NewBuffer(message))
		if err != nil {
			log.Fatal(err)
		}
		n, err = fs.Write(outCohId, bytes.NewBuffer(message))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n)
	*/

	var newRow uint32 = iRows / rLooks
	var newCol uint32 = iCols / cLooks

	mDataAll := *(*[]complex64)(unsafe.Pointer(&MBytes))
	sDataAll := *(*[]complex64)(unsafe.Pointer(&SBytes))
	ftDataAll := *(*[]float32)(unsafe.Pointer(&FlatBytes))
	outDataAllBytes := make([]byte, 8*newRow*newCol)
	outDataAll := *(*[]complex64)(unsafe.Pointer(&outDataAllBytes))
	/*
		for i := 0; i < 10; i++ {
			fmt.Println(mData[i])
		}
	*/

	//var covar, ftTemp complex64
	//var mCov, sCov float32

	runtime.GOMAXPROCS(64)
	for i := uint32(0); i < newRow; i++ {
		//mData := mDataAll[i*rLooks*iCols:]
		//sData := sDataAll[i*rLooks*iCols:]
		//ftData := ftDataAll[i*rLooks*iCols:]
		//outData := outDataAll[i*newCol:]

		go func(mData []complex64, sData []complex64, ftData []float32, outData []complex64, newCol uint32) {
			//for j := uint32(0); j < newCol; j++ {
			for j := uint32(0); j < newCol; j++ {
				var covar complex64 = complex(0, 0)
				var mCov, sCov float32 = 0.0, 0.0
				for ii := uint32(0); ii < rLooks; ii++ {
					for jj := uint32(0); jj < cLooks; jj++ {
						cur := ii*iCols + (j*cLooks + jj)
						//fmt.Println(mData[cur], sData[cur], ftData[cur])
						covar += complex64(complex128(mData[cur]) * cmplx.Conj(complex128(sData[cur])) * complex(math.Cos(float64(ftData[cur])), -math.Sin(float64(ftData[cur]))))
						//mCov += float32(cmplx.Abs(complex128(mData[cur])))
						//sCov += float32(cmplx.Abs(complex128(sData[cur])))
						mCov += float32(real(mData[cur])*real(mData[cur]) + imag(mData[cur])*imag(mData[cur]))
						sCov += float32(real(sData[cur])*real(sData[cur]) + imag(sData[cur])*imag(sData[cur]))
						//fmt.Println(covar, mData[cur], mCov, sData[cur], sCov)
					}
				}
				mCov = float32(math.Sqrt(float64(mCov)))
				sCov = float32(math.Sqrt(float64(sCov)))
				//fmt.Printf("%v %v %v\n", covar, mCov, sCov)
				msCov := mCov * sCov
				if math.Abs(float64(msCov)) > 10e-6 {
					outData[j] = complex(real(covar)/msCov, imag(covar)/msCov)
				} else {
					outData[j] = complex(0, 0)
				}
				//fmt.Println(outData[j])
			}
		}(mDataAll[i*rLooks*iCols:], sDataAll[i*rLooks*iCols:], ftDataAll[i*rLooks*iCols:], outDataAll[i*newCol:], newCol)
	}

	fmt.Println(len(outDataAll))
	//n, err := fs.Write(outCohId, bytes.NewReader(*(*[]byte)(unsafe.Pointer(&outDataAll))))
	n, err := fs.Write(outCohId, bytes.NewReader(outDataAllBytes))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
