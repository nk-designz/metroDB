package main

import (
    "log"
    "net/http"
    "fmt"
)

var stackMap = map[string]*Blockchain{}
var port = "8756"

func main() {

    fmt.Println("Listening on port", port, "...")

    router := NewRouter().StrictSlash(true)
    log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v",port), router))
}
