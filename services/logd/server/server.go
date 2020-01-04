package logd

import (
	"fmt"
	"log"
	"net"

	pb "../pb"
	"google.golang.org/grpc"
)

func main() {
	store := new(LogStore)
	store.open()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 70558))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
	grpcServer.Serve(lis)
	store.close()
}
