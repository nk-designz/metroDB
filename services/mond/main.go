package main

import (
	"fmt"
	"log"
	"os"

	mapdClient "github.com/nk-designz/metroDB/services/mapd/client"
)

func main() {
	hostname, _ := os.Hostname()
	log.Println(
		fmt.Println(
			mapdClient.New(hostname)))
}
