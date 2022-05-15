package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"homeWorkAdvancedPart2_gRPC_serv/internal/pb"
	"homeWorkAdvancedPart2_gRPC_serv/internal/server"
	"homeWorkAdvancedPart2_gRPC_serv/internal/store/inmemorystore"
	"log"
	"net"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	store := inmemorystore.NewStore()
	pb.RegisterDeliveryServiceServer(s, server.NewServer(store))
	log.Printf("grpc-server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
