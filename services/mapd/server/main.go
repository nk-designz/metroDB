package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"os"

	logd "github.com/nk-designz/metroDB/services/logd/client"
)

type Replica struct {
	logStore int
	offset   int64
}

type Mapd struct {
	index struct {
		memory map[string][]Replica
		disk   *os.File
	}
	logds map[int]*logd.Logd
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
	mapd.logds[0] = logd.New("127.0.0.1")
	return err
}

func (mapd *Mapd) close() {
	log.Println(`msg="shutting down map deamon..."`)
	mapd.index.disk.Close()
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

	bufferReader := new(bytes.Buffer)
	mapd.index.disk.Read(buffer)
	bufferReader.Read(buffer)
	decoder := gob.NewDecoder(bufferReader)
	err := decoder.Decode(&memoryIndex)
	mapd.index.memory = memoryIndex
	return err
}

func (mapd *Mapd) set(key string, value []byte) {
	// TODO: Based decision on Size in Logd
	logdIndex := rand.Intn(len(mapd.logds) - 1)
	//
	logd := mapd.logds[logdIndex]
	logd.Connect()
	defer logd.Close()
	offset := logd.Append(value)
	mapd.index.memory[key] = []Replica{
		Replica{
			offset:   offset,
			logStore: logdIndex}}
	go mapd.updatePersistentIndex()
}

func (mapd *Mapd) get(key string) []byte {
	entry := mapd.index.memory[key]
	logd := mapd.logds[entry[0].logStore]
	offset := entry[0].offset
	logd.Connect()
	defer logd.Close()
	return logd.Get(offset)

}

func main() {
	mapd := new(Mapd)
	err := mapd.init()
	defer mapd.close()
	if err != nil {
		log.Fatalln(err)
	}
	mapd.set(
		"1",
		[]byte("test"))
	fmt.Println(
		mapd.get("1"))
}
