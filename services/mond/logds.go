package main

import (
	"fmt"
	"log"
	"os"

	mapdClient "github.com/nk-designz/metroDB/services/mapd/client"
)

func main() {
	log.Println(
		fmt.Println(
			func(name string, err error) { 
				mapdClient.New(name) 
			}(os.Hostname())))
}
