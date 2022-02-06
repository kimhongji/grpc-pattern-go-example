package main

import (
	"google.golang.org/grpc"
	pb "grpc-pattern-go-example/proto/ecommerce"
	"grpc-pattern-go-example/proto/handler"
	"log"
	"net"
)

const (
	port = "8081"
)

func main() {
	lis, err := net.Listen("tcp", "localhost"+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &handler.Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
