package main

import (
	"fmt"
	"math/cmplx"
	/*"log"*/ /*"time"*/ /*alluxio "github.com/Alluxio/alluxio-go"*/ /*"github.com/Alluxio/alluxio-go/option"*/
	"insar"
)

func TestComplex() {
	var c complex64 = 3.0 + 4.0i
	fmt.Printf("real: %f, imag: %f, value: %v, norm: %f\n", real(c), imag(c), c, cmplx.Abs(complex128(c)))
}

func main() {
	/*
		fs := alluxio.NewClient("10.2.152.24", 39999, time.Second*5)
		ok, err := fs.Exists("/insar/fineReg/1.rmg", &option.Exists{})
		if err != nil {
			log.Fatal(err)

		}
	*/
	//id, err := fs.CreateFile("/insar/2.txt", &option.CreateFile{})
	//if err != nil {
	//log.Fatal(err)
	//}
	//n, err := fs.Write(id, bytes.NewBuffer(make([]byte, 1024)))
	//if err != nil {
	//log.Fatal(err)
	//}
	//defer fs.Close(id)
	//fmt.Println(n)

	//fmt.Println(ok)

	insar.CoherenceFun("10.2.152.24", "/insar/fineReg/1.rmg", "/insar/fineReg/2.rmg", "/insar/flatPha/1-2.rmg",
		"/insar/coh/1-2.rmg", 4, 2, 18707, 21180)
	insar.CoherenceFun("10.2.152.24", "/insar/fineReg/1.rmg", "/insar/fineReg/4.rmg", "/insar/flatPha/1-4.rmg",
		"/insar/coh/1-4.rmg", 4, 2, 18707, 21180)
	insar.CoherenceFun("10.2.152.24", "/insar/fineReg/2.rmg", "/insar/fineReg/3.rmg", "/insar/flatPha/2-3.rmg",
		"/insar/coh/2-3.rmg", 4, 2, 18707, 21180)
	insar.CoherenceFun("10.2.152.24", "/insar/fineReg/2.rmg", "/insar/fineReg/4.rmg", "/insar/flatPha/2-4.rmg",
		"/insar/coh/2-4.rmg", 4, 2, 18707, 21180)
	insar.CoherenceFun("10.2.152.24", "/insar/fineReg/3.rmg", "/insar/fineReg/4.rmg", "/insar/flatPha/3-4.rmg",
		"/insar/coh/3-4.rmg", 4, 2, 18707, 21180)
	//insar.CompareComplex64("10.2.152.24", "/insar/coh/1-2.rmg", "/insar/linux_coh/1-2.rmg", 100)

	//TestComplex()

}
