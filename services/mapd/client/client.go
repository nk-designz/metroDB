package mapd

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	pb "github.com/nk-designz/metroDB/services/mapd/pb"
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

func (mapd *Mapd) Connect() error {
	var err error
	mapd.connection, err = grpc.Dial(mapd.address, grpc.WithInsecure())
	mapd.client = pb.NewMapdClient(mapd.connection)
	return err
}

func (mapd *Mapd) Set(key string, value []byte) (bool, error) {
	t, err := mapd.client.Set(context.Background(), &pb.SetRequest{Key: key, Value: value})
	return t.GetErr(), err
}

func (mapd *Mapd) SetSafe(key string, value []byte) (bool, error) {
	t, err := mapd.client.SetSafe(context.Background(), &pb.SetRequest{Key: key, Value: value})
	return t.GetErr(), err
}

func (mapd *Mapd) Replicate(entry *pb.Entry) error {
	_, err := mapd.client.Replicate(context.Background(), entry)
	return err
}

func (mapd *Mapd) Get(key string) ([]byte, error) {
	value, err := mapd.client.Get(context.Background(), &pb.GetRequest{Key: key})
	return value.GetValue(), err
}

func (mapd *Mapd) GetRandProbe() (string, int64, []byte, []byte, error) {
	reply, err := mapd.client.GetRandProbe(context.Background(), &pb.Void{})
	return reply.GetKey(), reply.GetOffset(), reply.GetHash(), reply.GetValue(), err
}

func (mapd *Mapd) GetDefiProbe(key string, deph uint64) (string, int64, []byte, []byte, error) {
	reply, err := mapd.client.GetDefiProbe(context.Background(), &pb.ProbeRequest{Key: key, Deph: deph})
	return reply.GetKey(), reply.GetOffset(), reply.GetHash(), reply.GetValue(), err
}