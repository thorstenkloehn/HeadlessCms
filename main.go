package main

import (
    "fmt"
    "github.com/thorstenkloehn/HeadlessCms/datenbank"
)


func main() {

    start := datenbank.Dbinfo()
    fmt.Println(start)


}