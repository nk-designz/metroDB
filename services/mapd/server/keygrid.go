package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"

	"lukechampine.com/blake3"

	logd "github.com/nk-designz/metroDB/services/logd/client"
	mapdClient "github.com/nk-designz/metroDB/services/mapd/client"
	pb "github.com/nk-designz/metroDB/services/mapd/pb"
)

type Replica struct {
	LogStore int   `json:"store"`
	Offset   int64 `json:"offset"`
	Sum      int64 `json:"sum"`
}

type Logds struct {
	logd *logd.Logd
	name string
	size int64
}

type Cluster []*mapdClient.Mapd

type Mapd struct {
	index struct {
		memory map[string][]Replica `json:"index"`
		disk   *os.File
		sync   struct {
			ticker *time.Ticker
			quit   chan struct{}
		}
	}
	logds   []Logds
	cluster Cluster
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
	err := json.NewEncoder(mapd.index.disk).Encode(mapd.index.memory)
	if err != nil {
		panic(err)
	}
	return err
}

func (mapd *Mapd) retrivePersistentIndex() error {
	filestat, err := mapd.index.disk.Stat()
	if err != nil {
		panic(err)
	}
	memoryIndex := make(map[string][]Replica, filestat.Size())
	log.Println(`msg="retrieving map from disk..."`)
	err = json.NewDecoder(mapd.index.disk).Decode(&memoryIndex)
	mapd.index.memory = memoryIndex
	return err
}

func (mapd *Mapd) set(key string, value []byte) {
	sort.Slice(mapd.logds, func(i, j int) bool {
		return mapd.logds[i].size < mapd.logds[j].size
	})
	hashValue := blake3.Sum512(value)
	for logdIndex := range []int{0, 1} {
		logdInstance := mapd.logds[logdIndex].logd
		logdInstance.Connect()
		valueOffset := logdInstance.Append(value)
		hashOffset := logdInstance.Append(hashValue[:])
		mapd.logds[logdIndex].size = valueOffset
		replica := Replica{
			Offset:   valueOffset,
			LogStore: logdIndex,
			Sum:      hashOffset}
		if replicas, ok := mapd.index.memory[key]; ok {
			mapd.index.memory[key] = append(replicas, replica)
		} else {
			mapd.index.memory[key] = []Replica{replica}
		}
		logdInstance.Close()
		go func(replica Replica) {
			for _, member := range mapd.cluster {
				member.Replicate(
					&pb.Entry{
						Key:      key,
						LogStore: int32(replica.LogStore),
						Offset:   replica.Offset})
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
	if entrys, exist := mapd.index.memory[key]; !exist {
		log.Println(`msg="key not found"`)
		return []byte{}
	}
	for replic := len(entrys) - 1; replic < 0; n-- {
		logd := mapd.logds[entrys[replic].LogStore].logd
		logd.Connect()
		value := logd.Get(entrys[replic].Offset)
		sum := logd.Get(entrys[replic].Sum)
		logd.Close()
		if(sum != blake3.Sum512(value)) {
			log.Println(`msg="broken replica! moving on..."`)
			continue
		}
		return value
	}
	log.Println(`msg="no valid replicas"`)
	return []byte
}

func (mapd *Mapd) setReplica(key string, replica Replica) {
	if replicas, ok := mapd.index.memory[key]; ok {
		mapd.index.memory[key] = append(replicas, replica)
	} else {
		mapd.index.memory[key] = []Replica{replica}
	}
}

func (mapd *Mapd) getDefiProbe(key string, deph uint64) (string, int64, []byte, []byte, error) {
	return key,
		mapd.index.memory[key][deph].Offset,
		mapd.logds[mapd.index.memory[key][deph].LogStore].logd.Get(mapd.index.memory[key][deph].Sum),
		mapd.get(key),
		nil
}

func (mapd *Mapd) getRandProbe() (string, int64, []byte, []byte, error) {
	interator := 0
	randomBreak := rand.Intn(len(mapd.index.memory))
	for key, replica := range mapd.index.memory {
		if interator >= randomBreak {
			return key,
				replica[0].Offset,
				mapd.logds[replica[0].LogStore].logd.Get(replica[0].Sum),
				mapd.get(key),
				nil
		}
		interator++
	}
	return "", 0, []byte{}, []byte{}, nil //TODO Error handling
}
