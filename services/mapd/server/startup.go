package main

import(
	"log"
	"fmt"
	"os"
	"flag"
	"time"
	"strings"
	"math/rand"

	logd "github.com/nk-designz/metroDB/services/logd/client"
	mapdClient "github.com/nk-designz/metroDB/services/mapd/client"
)

func (mapd *Mapd) init() error {
	log.Println(`msg="initializing map deamon..."`)
	// Add Logger File to Logd
	file, err := os.OpenFile(
		fmt.Sprintf("%s/mapd.db", os.Getenv("MAPD_INDEX_PATH")),
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0600)
	mapd.index.disk = file
	log.Println(fmt.Sprintf(`msg="Persisting to File: %s"`, mapd.index.disk.Name()))
	err = mapd.retrivePersistentIndex()
	if err != nil {
		log.Println(fmt.Sprintf(`msg="Init blank map" err="%v"`, err))
		mapd.index.memory = map[string][]Replica{}
		err = nil
	}
	mapd.sheduleDiskSync()
	for k, v := range mapd.index.memory {
		log.Println(
			fmt.Sprintf(`msg="found key-value-pair" key="%s" offset="%x" logstore="%d"`, k, v[0].Offset, v[0].LogStore))
	}
	logdFlag := flag.String("logds", "logds-1.metrodb.cluster.local", "comma seperated list of logds")
	mapdFlag := flag.String("cluster", "mapd-1.metro.cluster.local", "comman seperated list of mapds")
	flag.Parse()
	logds := strings.Split(*logdFlag, ",")
	mapds := strings.Split(*mapdFlag, ",")
	mapd.logds = make([]Logds, len(logds))
	mapd.cluster = make(Cluster, len(mapds))
	for index, name := range logds {
		mapd.logds[index] = Logds{
			logd: logd.New(name),
			name: name,
			size: rand.Int63n(3)}
	}
	for index, name := range mapds {
		mapd.cluster[index] = mapdClient.New(name)
		for {
			if mapd.cluster[index].Connect() == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
	return err
}