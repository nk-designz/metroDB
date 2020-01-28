package main

import (
	"context"
	"log"

	pb "github.com/nk-designz/metroDB/services/mapd/pb"
)

type MapdServer struct {
	mapd *Mapd
}

func (server *MapdServer) Set(ctx context.Context, request *pb.SetRequest) (*pb.SetReply, error) {
	reply := new(pb.SetReply)
	server.mapd.set(request.Key, request.Value)
	reply.Err = false
	return reply, nil
}

func (server *MapdServer) SetSafe(ctx context.Context, request *pb.SetRequest) (*pb.SetReply, error) {
	reply := new(pb.SetReply)
	reply.Err = false
	return reply, nil
}

func (server *MapdServer) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetReply, error) {
	reply := new(pb.GetReply)
	reply.Value = server.mapd.get(request.Key)
	return reply, nil
}

func (server *MapdServer) Replicate(ctx context.Context, request *pb.Entry) (*pb.Void, error) {
	server.mapd.setReplica(request.Key, Replica{
		LogStore: int(request.LogStore),
		Offset:   request.Offset,
		Sum:      request.SumOffset})
	return new(pb.Void), nil
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
