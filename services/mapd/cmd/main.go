package main

import (
	"fmt"
	"os"

	mapd "github.com/nk-designz/metroDB/services/mapd/client"
)

func main() {
	address := os.Args[1]
	mapdInstance := mapd.New(address)
	if os.Args[2] == "set" && os.Args[5] == "safe" {
		mapdInstance.SetSafe(
			os.Args[3],
			[]byte(os.Args[4]))
		fmt.Println("Seems ok")
	} else if os.Args[2] == "set" {
		mapdInstance.Set(
			os.Args[3],
			[]byte(os.Args[4]))
		fmt.Println("Seems ok")
	} else if os.Args[2] == "get" {
		fmt.Println(mapdInstance.Get(os.Args[3]))
	} else {
		fmt.Println("No valid command")
	}
}
