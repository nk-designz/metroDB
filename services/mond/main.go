package main

import (
	"fmt"
)

func main() {
	logds := new(Logds)
	logds.Init()
	defer logds.Close()
	for logd := range logds.logds {
		fmt.Println(logd.Get(898))
	}
}
