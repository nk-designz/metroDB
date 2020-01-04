package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/nk-designz/metroDB/services/logd/pb"
	"google.golang.org/grpc"
)

const (
	PORT = 7558
)

type LogdServer struct {
	Log *Log
}

func newLogdServer() pb.LogdServer {
	logdServer := new(LogdServer)
	logdServer.Log = new(Log)
	logdServer.Log.open()
	return logdServer
}

func (server *LogdServer) Append(ctx context.Context, request *pb.SetRequest) (*pb.SetReply, error) {
	reply := new(pb.SetReply)
	reply.Offset = server.Log.append(request.Entry)
	return reply, nil
}

func (server *LogdServer) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetReply, error) {
	reply := new(pb.GetReply)
	reply.Entry = server.Log.get(request.Offset)
	return reply, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatalf(fmt.Sprintf(`msg="Failed running server" port="%d" error="%v"`, PORT, err))
	}
	log.Println(fmt.Sprintf(`msg="Running server" port="%d"`, PORT))
	grpcServer := grpc.NewServer()
	pb.RegisterLogdServer(grpcServer, newLogdServer())
	grpcServer.Serve(lis)
}
