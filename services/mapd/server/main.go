package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	logd "github.com/nk-designz/metroDB/services/logd/client"
	pb "github.com/nk-designz/metroDB/services/mapd/pb"
	"google.golang.org/grpc"
)

const (
	defaultSyncSchedule = 10
	PORT                = 7550
)

type Replica struct {
	logStore int
	offset   int64
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
	logds []*logd.Logd
}

func (mapd *Mapd) init() error {
	log.Println(`msg="initializing map deamon..."`)
	// Add Logger File to Logd
	mapd.logds = make([]*logd.Logd, len(os.Args))
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
	for i, v := range os.Args {
		if i > 0 {
			mapd.logds[i-1] = logd.New(v)
		}
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
	log.Println(`msg="set new key"`, key, len(value))
}

func (mapd *Mapd) setSafe(key string, value []byte) {
	mapd.set(key, value)
	mapd.index.disk.Sync()
}

func (mapd *Mapd) get(key string) []byte {
	entry := mapd.index.memory[key]
	logd := mapd.logds[entry[0].logStore]
	offset := entry[0].offset
	logd.Connect()
	defer logd.Close()
	log.Println(`msg="get key"`, key)
	return logd.Get(offset)

}

type MapdServer struct {
	mapd *Mapd
}

func (server *MapdServer) Set(ctx context.Context, request *pb.SetRequest) (*pb.SetReply, error) {
	reply := new(pb.SetReply)
	// TODO: Error handling on Panic
	server.mapd.set(request.Key, request.Value)
	reply.Err = false
	return reply, nil
}

func (server *MapdServer) SetSafe(ctx context.Context, request *pb.SetRequest) (*pb.SetReply, error) {
	reply := new(pb.SetReply)
	reply.Err = func() bool {
		server.mapd.setSafe(request.Key, request.Value)
		errChan := make(chan bool)
		defer func(errChan chan bool) {
			r := recover()
			if r != nil {
				errChan <- false
			} else {
				errChan <- true
			}
		}(errChan)
		rbool := <-errChan
		close(errChan)
		return rbool
	}()
	return reply, nil
}

func (server *MapdServer) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetReply, error) {
	reply := new(pb.GetReply)
	reply.Value = server.mapd.get(request.Key)
	return reply, nil
}

func newMapdServer() *MapdServer {
	mapdServerInstance := new(MapdServer)
	mapdServerInstance.mapd = new(Mapd)
	err := mapdServerInstance.mapd.init()
	if err != nil {
		log.Fatalln(err)
	}
	return mapdServerInstance
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatalf(fmt.Sprintf(`msg="Failed running server" port="%d" error="%v"`, PORT, err))
	}
	log.Println(fmt.Sprintf(`msg="Running server" port="%d"`, PORT))
	grpcServer := grpc.NewServer()
	pb.RegisterMapdServer(grpcServer, newMapdServer())
	grpcServer.Serve(lis)
}
