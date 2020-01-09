package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"

	logd "github.com/nk-designz/metroDB/services/logd/client"
)

type Replica struct {
	logStore int
	offset   int64
}

type Logds struct {
	logd *logd.Logd
	name string
	size int64
}

type Mapd struct {
	index struct {
		memory map[string][]Replica
		disk   *os.File
		sync   struct {
			ticker *time.Ticker
			quit   chan struct{}
		}
	}
	logds []Logds
}

func (mapd *Mapd) init() error {
	log.Println(`msg="initializing map deamon..."`)
	// Add Logger File to Logd
	mapd.logds = make([]Logds, len(os.Args))
	file, err := os.OpenFile(
		fmt.Sprintf("%smapd.db", os.Getenv("MAPD_INDEX_PATH")),
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0600)
	mapd.index.disk = file
	log.Println(fmt.Sprintf(`msg="Persisting to File: %s"`, mapd.index.disk.Name()))
	err = mapd.retrivePersistentIndex()
	if err != nil {
		mapd.index.memory = map[string][]Replica{}
		err = nil
	}
	mapd.sheduleDiskSync()
	for k, v := range mapd.index.memory {
		log.Println(
			fmt.Sprintf(`msg="found key-value-pair" key="%s" offset="%x" logstore="%d"`, k, v[0].offset, v[0].logStore))
	}
	for index, name := range os.Args {
		if index != 0 {
			mapd.logds[index-1] = Logds{
				logd: logd.New(name),
				name: name,
				size: rand.Int63()}
		}
	}
	//TODO: Find out where the last entry nil comes from
	mapd.logds = mapd.logds[:len(mapd.logds)-1]
	log.Println(mapd.logds, os.Args)
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
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(mapd.index.memory)
	_, err = mapd.index.disk.Write(buffer.Bytes())
	return err
}

func (mapd *Mapd) retrivePersistentIndex() error {
	var memoryIndex map[string][]Replica
	var buffer []byte

	log.Println(`msg="retrieving map from disk..."`)
	bufferReader := new(bytes.Buffer)
	mapd.index.disk.Read(buffer)
	bufferReader.Read(buffer)
	decoder := gob.NewDecoder(bufferReader)
	err := decoder.Decode(&memoryIndex)
	mapd.index.memory = memoryIndex
	return err
}

func (mapd *Mapd) set(key string, value []byte) {
	sort.Slice(mapd.logds, func(i, j int) bool {
		return mapd.logds[i].size > mapd.logds[j].size
	})
	logdIndex := len(mapd.logds) - 1
	logd := mapd.logds[logdIndex].logd
	logd.Connect()
	defer logd.Close()
	offset := logd.Append(value)
	mapd.index.memory[key] = []Replica{
		Replica{
			offset:   offset,
			logStore: logdIndex}}
	go mapd.updatePersistentIndex()
	log.Println(`msg="set new key"`, key, len(value))
}

func (mapd *Mapd) setSafe(key string, value []byte) {
	mapd.set(key, value)
	mapd.index.disk.Sync()
}

func (mapd *Mapd) get(key string) []byte {
	entry := mapd.index.memory[key]
	logd := mapd.logds[entry[0].logStore].logd
	offset := entry[0].offset
	logd.Connect()
	defer logd.Close()
	log.Println(`msg="get key"`, key)
	return logd.Get(offset)

}
