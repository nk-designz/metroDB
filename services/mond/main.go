package main

import (
	"fmt"
	"log"
	"os"

	mapdClient "github.com/nk-designz/metroDB/services/mapd/client"
)

type cluster []*mapdClient.Mapd

func main() {
	mapds := make(cluster, len(os.Args)-1)
	for index, hostname := range os.Args {
		if index != 0 {
			fmt.Println(index)
			mapds[index-1] = mapdClient.New(hostname)
		}
	}
	for {
		schedule(mapds)
	}
}

func schedule(mapds cluster) {
	for _, mapd := range mapds {
		mapd.Connect()
		log.Println(
			fmt.Print(mapd))
	}
}
