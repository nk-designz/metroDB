package logd

import (
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

func New(address string) *Logd {
	instance := new(Logd)
	instance.address = fmt.Sprintf("%s:7558", address)
	instance.Connect()
	return instance
}
