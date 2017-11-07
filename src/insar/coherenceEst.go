package insar

import (
	"fmt"
	alluxio "github.com/Alluxio/alluxio-go"
	"github.com/Alluxio/alluxio-go/option"
	"log"
    "io/ioutil"
    "time"
)

func CoherenceFun(ip string, inM string, inS string, inFlat string, outCoh string, rLooks uint32, cLooks uint32, iCols uint32, iRows uint32) {
	fs := alluxio.NewClient(ip, 39999, time.Minute)

	fInMId, err := fs.OpenFile(inM, &option.OpenFile{})
	if err != nil {
		log.Fatal(err)
	}
    fInMReader, err := fs.Read(fInMId)
    if err != nil {
		log.Fatal(err)
    }
    fInMBytes, err := ioutil.ReadAll(fInMReader)
    if err != nil {
		log.Fatal(err)
    }
    fmt.Println(len(fInMBytes))
	defer fs.Close(fInMId)

	fInSId, err := fs.OpenFile(inS, &option.OpenFile{})
	if err != nil {
		log.Fatal(err)
	}
    fInSReader, err := fs.Read(fInSId)
    if err != nil {
		log.Fatal(err)
    }
    fInSBytes, err := ioutil.ReadAll(fInSReader)
    if err != nil {
		log.Fatal(err)
    }
    fmt.Println(len(fInSBytes))
	defer fs.Close(fInSId)

	fInFlatId, err := fs.OpenFile(inFlat, &option.OpenFile{})
	if err != nil {
		log.Fatal(err)
	}
    fInFlatReader, err := fs.Read(fInFlatId)
    if err != nil {
		log.Fatal(err)
    }
    fInFlatBytes, err := ioutil.ReadAll(fInFlatReader)
    if err != nil {
		log.Fatal(err)
    }
    fmt.Println(len(fInFlatBytes))
	defer fs.Close(fInFlatId)

	fOutCohId, err := fs.OpenFile(outCoh, &option.OpenFile{})
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close(fOutCohId)

	/*var newRow uint32 = iRows / rLooks*/
	/*var newCol uint32 = iCols / cLooks*/

	/*mData := make([]complex64, rLooks*iCols)*/
	/*sData := make([]complex64, rLooks*iCols)*/
	/*ftData := make([]float32, rLooks*iCols)*/
	/*outData := make([]complex64, newCol)*/

	/*var covar, ftTemp complex64*/
	/*var mCov, sCov float32*/
	//for i := 0; i < newRow; ++i {
	//}
	
}
