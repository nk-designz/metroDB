package logd

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nk-designz/metroDB/services/logd/pb"
	"google.golang.org/grpc"
)

type Logd struct {
	address    string
	client     pb.LogdClient
	connection *grpc.ClientConn
}

func (newInstance *Logd) Connect() {
	var err error
	newInstance.connection, err = grpc.Dial(newInstance.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	newInstance.client = pb.NewLogdClient(newInstance.connection)
}

func (newInstance *Logd) Close() {
	defer newInstance.connection.Close()
}

func (newInstance *Logd) Append(entry []byte) int64 {
	offset, err := newInstance.client.Append(context.Background(), &pb.SetRequest{Entry: entry})
	if err != nil {
		log.Fatalln(err)
	}
	return offset.Offset
}

func (newInstance *Logd) Get(offset ...int64) []byte {
	var entry *pb.GetReply
	var err error
	if len(offset) == 0 {
		entry, err = newInstance.client.Get(context.Background(), &pb.GetRequest{})
	} else {
		entry, err = newInstance.client.Get(context.Background(), &pb.GetRequest{Offset: offset[0]})
	}
	if err != nil {
		log.Fatalln(err)
	}
	return entry.Entry
}

func New(address string) *Logd {
	instance := new(Logd)
	instance.address = fmt.Sprintf("%s:7558", address)
	instance.Connect()
	return instance
}
