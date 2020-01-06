package main

import (
	"fmt"
	"os"

	mapd "github.com/nk-designz/metroDB/services/mapd/client"
)

func main() {
	address := os.Args[1]
	mapd := mapd.New(address)
	if os.Args[2] == "set" && os.Args[5] == "safe" {
		mapd.setSafe(os.Args[3], []byte(os.Args[4]))
		fmt.Println("Seems ok")
	} else if os.Args[2] == "set" {
		mapd.Set(os.Args[3], []byte(os.Args[4]))
		fmt.Println("Seems ok")
	} else if os.Args[2] == "get" {
		fmt.Println(mapd.get(os.Args[3]))
	} else {
		fmt.Println("No valid command")
	}
}
