package main

import (
	"context"
	"fmt"
	"log"
	"net"

	logd "github.com/nk-designz/metroDB/services/logd/logd"
	pb "github.com/nk-designz/metroDB/services/logd/pb"
	"google.golang.org/grpc"
)

type LogdServer struct {
	Log *logd.LogStore
}

func newLogdServer() pb.LogdServer {
	logdServer := new(LogdServer)
	logdServer.Log = new(logd.LogStore)
	logdServer.Log.open()
	return logdServer
}

func (server *LogdServer) Append(ctx context.Context, request *pb.LogdRequest) (*pb.LogdReply, error) {
	reply := new(pb.LogdReply)
	reply.Data = server.Log.append(request.Data)
	return reply, nil
}

func (server *LogdServer) ReadEntryAt(ctx context.Context, request *pb.LogdRequest) (*pb.LogdReply, error) {
	reply := new(pb.LogdReply)
	reply.Data = server.Log.get(request.Data)
	return reply, nil
}

func (server *LogdServer) ReadEntryLast(ctx context.Context, request *pb.LogdRequest) (*pb.LogdReply, error) {
	reply := new(pb.LogdReply)
	reply.Data = server.Log.get(s.LogStore.LastOffset)
	return reply, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 70558))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLogdServer(grpcServer, &LogdServer{})
	grpcServer.Serve(lis)
}
