package main

import (
        "fmt"
        "log"
        "time"

        alluxio "github.com/Alluxio/alluxio-go"
        "github.com/Alluxio/alluxio-go/option"

       )


func main() {
    fs := alluxio.NewClient("10.2.152.24", 39999, time.Second*5)
    ok, err := fs.Exists("/insar/1.txt", &option.Exists{})
    if err != nil {
        log.Fatal(err)
    
    }
    
    fmt.Println(ok)

}
