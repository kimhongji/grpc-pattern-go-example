package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-pattern-go-example/proto/ecommerce"
	"grpc-pattern-go-example/proto/handler"
	"log"
	"net"
)

const (
	port = "8081"
)

func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Pre-processing logic
	// Gets info about the current RPC call by examining the args passed in
	log.Println("==== [Server Interceptor] ", info.FullMethod)
	log.Printf(" Pre Proc Message : %s", req)


	// Invoking the handler to complete the normal execution of a unary RPC.
	m, err := handler(ctx, req)

	// Post processing logic
	log.Printf(" Post Proc Message : %s", m)
	return m, err
}
func main() {
	lis, err := net.Listen("tcp", "localhost"+":"+port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(orderUnaryServerInterceptor))
	pb.RegisterOrderManagementServer(s, &handler.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



