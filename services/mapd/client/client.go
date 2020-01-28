package mapd

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nk-designz/metroDB/services/mapd/pb"
	"google.golang.org/grpc"
)

type Mapd struct {
	address    string
	client     pb.MapdClient
	connection *grpc.ClientConn
}

func New(address string) *Mapd {
	mapdInstance := new(Mapd)
	mapdInstance.address = fmt.Sprintf("%s:7550", address)
	mapdInstance.Connect()
	return mapdInstance
}

func (mapd *Mapd) Connect() {
	var err error
	mapd.connection, err = grpc.Dial(mapd.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	mapd.client = pb.NewMapdClient(mapd.connection)
}

func (mapd *Mapd) Set(key string, value []byte) (bool, error) {
	t, err := mapd.client.Set(context.Background(), &pb.SetRequest{Key: key, Value: value})
	return t.Err, err
}

func (mapd *Mapd) SetSafe(key string, value []byte) (bool, error) {
	t, err := mapd.client.SetSafe(context.Background(), &pb.SetRequest{Key: key, Value: value})
	return t.Err, err
}

func (mapd *Mapd) Replicate(entry *pb.Entry) error {
	if _, err := mapd.client.Replicate(context.Background(), entry); err != nil {
		return err
	}
}

func (mapd *Mapd) GetSum(key string) ([]byte, error) {
	sum, err := mapd.client.GetSum(context.Background(), &pb.GetRequest{Key: key})
	return sum.Value, err
}

func (mapd *Mapd) Get(key string) ([]byte, error) {
	value, err := mapd.client.Get(context.Background(), &pb.GetRequest{Key: key})
	if err != nil {
		log.Println(
			fmt.Sprintf(`msg="could not get key" key="%s" error="%v"`, key, err))
	}
	return value.GetValue(), err
}
