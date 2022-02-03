package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	pb "grpc-pattern-go-example/proto/ecommerce"
	"log"
	"time"
)

const (
	port = "8081"
)

func main() {
	conn, err := grpc.Dial("localhost"+":"+port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	retrievedOrder, rr := client.GetOrder(ctx, &wrappers.StringValue{Value: "102"})
	if rr != nil {
		log.Fatalf("Server can't find the order: %v", rr)
	}
	log.Print("GetOrder Response -> : ", retrievedOrder)

}
