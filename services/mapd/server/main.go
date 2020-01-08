package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/nk-designz/metroDB/services/mapd/pb"
	"google.golang.org/grpc"
)

const (
	defaultSyncSchedule = 10
	PORT                = 7550
)

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
