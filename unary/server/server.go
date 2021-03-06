package main

import (
	"grpc-pattern-go-example/proto/ecommerce"
	"grpc-pattern-go-example/proto/handler"

	"google.golang.org/grpc"

	"log"
	"net"
)

const (
	port = "8081"
)

func main() {
	lis, err := net.Listen("tcp", "localhost"+":"+port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	ecommerce.RegisterOrderManagementServer(s, &handler.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
