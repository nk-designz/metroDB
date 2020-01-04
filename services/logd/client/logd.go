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

func (newInstance *Logd) Append(entry []byte) []byte {
	offset, err := newInstance.client.Append(context.Background(), &pb.LogdRequest{Data: entry})
	if err != nil {
		log.Fatalln(err)
	}
	return offset.Data
}

func (newInstance *Logd) Get(offset []byte) []byte {
	var entry *pb.LogdReply
	var err error
	if offset == nil {
		entry, err = newInstance.client.ReadEntryLast(context.Background(), &pb.LogdRequest{Data: nil})
	} else {
		entry, err = newInstance.client.ReadEntryAt(context.Background(), &pb.LogdRequest{Data: offset})
	}
	if err != nil {
		log.Fatalln(err)
	}
	return entry.Data
}

func New(address string) *Logd {
	instance := new(Logd)
	instance.address = fmt.Sprintf("%s:7558", address)
	instance.Connect()
	return instance
}
