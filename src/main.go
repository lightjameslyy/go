package main

import (
	/*"fmt"*/
	/*"log"*/
	/*"time"*/

    /*alluxio "github.com/Alluxio/alluxio-go"*/
    /*"github.com/Alluxio/alluxio-go/option"*/
    "insar"
)

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

    insar.CoherenceFun("10.2.152.24", "/insar/fineReg/1.rmg", "/insar/fineReg/1.rmg", "/insar/flatPha/1-2.rmg", 
                        "/insar/coh/1-2.rmg", 4, 2, 18707, 21180) 

}
