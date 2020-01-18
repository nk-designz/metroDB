package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	logd "github.com/nk-designz/metroDB/services/logd/client"
	mapdClient "github.com/nk-designz/metroDB/services/mapd/client"
	pb "github.com/nk-designz/metroDB/services/mapd/pb"
)

type Replica struct {
	LogStore int
	Offset   int64
}

type Logds struct {
	logd *logd.Logd
	name string
	size int64
}

type cluster []*mapdClient.Mapd

type Mapd struct {
	index struct {
		memory map[string][]Replica
		disk   *os.File
		sync   struct {
			ticker *time.Ticker
			quit   chan struct{}
		}
	}
	logds   []Logds
	cluster cluster
}

func (mapd *Mapd) init() error {
	log.Println(`msg="initializing map deamon..."`)
	// Add Logger File to Logd
	file, err := os.OpenFile(
		fmt.Sprintf("%smapd.db", os.Getenv("MAPD_INDEX_PATH")),
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
	mapd.cluster = make(cluster, len(mapds))
	for index, name := range logds {
		mapd.logds[index] = Logds{
			logd: logd.New(name),
			name: name,
			size: rand.Int63n(3)}
	}
	for index, name := range mapds {
		mapd.cluster[index] = mapdClient.New(name)
		mapd.cluster[index].Connect()
	}
	return err
}

func (mapd *Mapd) close() {
	log.Println(`msg="shutting down map deamon..."`)
	close(mapd.index.sync.quit)
	mapd.index.disk.Close()
}

func (mapd *Mapd) sheduleDiskSync() {
	mapd.index.sync.ticker = time.NewTicker(defaultSyncSchedule * time.Second)
	mapd.index.sync.quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-mapd.index.sync.ticker.C:
				log.Println(`msg="Syncing map to disk"`)
				mapd.index.disk.Sync()
			case <-mapd.index.sync.quit:
				log.Println(`msg="stopping disk sync"`)
				mapd.index.sync.ticker.Stop()
				return
			}
		}
	}()
}

func (mapd *Mapd) updatePersistentIndex() error {
	encoder := gob.NewEncoder(mapd.index.disk)
	err := encoder.Encode(mapd.index.memory)
	if err != nil {
		panic(err)
	}
	return err
}

func (mapd *Mapd) retrivePersistentIndex() error {
	filestat, err := mapd.index.disk.Stat(); if err != nil {
		panic(err)
	}
	memoryIndex := make(map[string][]Replica, filestat.Size())
	log.Println(`msg="retrieving map from disk..."`)
	decoder := gob.NewDecoder(mapd.index.disk)
	err = decoder.Decode(&memoryIndex)
	mapd.index.memory = memoryIndex
	return err
}

func (mapd *Mapd) set(key string, value []byte) {
	sort.Slice(mapd.logds, func(i, j int) bool {
		return mapd.logds[i].size > mapd.logds[j].size
	})
	for logdIndex := range []int{0, 1} {
		logdInstance := mapd.logds[logdIndex].logd
		logdInstance.Connect()
		valueOffset := logdInstance.Append(value)
		mapd.logds[logdIndex].size = valueOffset
		replica := Replica{
			Offset:   valueOffset,
			LogStore: logdIndex}
		if replicas, ok := mapd.index.memory[key]; ok {
			mapd.index.memory[key] = append(replicas, replica)
		} else {
			mapd.index.memory[key] = []Replica{replica}
		}
		logdInstance.Close()
		go func(replica Replica) {
			for _, member := range mapd.cluster {
				member.Replicate(&pb.Entry{Key: key, LogStore: int32(replica.LogStore), Offset: replica.Offset})
			}
		}(replica)
	}
	go mapd.updatePersistentIndex()
	log.Println(`msg="set new key"`, key, len(value))
}

func (mapd *Mapd) setSafe(key string, value []byte) {
	mapd.set(key, value)
	mapd.index.disk.Sync()
}

func (mapd *Mapd) get(key string) []byte {
	log.Println(`msg="get key"`, key)
	if entrys, exist := mapd.index.memory[key]; exist {
		replic := len(entrys) - 1
		logd := mapd.logds[entrys[replic].LogStore].logd
		offset := entrys[replic].Offset
		logd.Connect()
		defer logd.Close()
		return logd.Get(offset)
	} else {
		log.Println(`msg="key not found"`)
		return []byte{}
	}
}

func (mapd *Mapd) setReplica(key string, replica Replica) {
	if replicas, ok := mapd.index.memory[key]; ok {
		mapd.index.memory[key] = append(replicas, replica)
	} else {
		mapd.index.memory[key] = []Replica{replica}
	}
}
