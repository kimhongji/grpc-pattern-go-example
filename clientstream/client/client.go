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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	client.AddOrder(ctx, &pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00})
	client.AddOrder(ctx, &pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00})
	client.AddOrder(ctx, &pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00})
	client.AddOrder(ctx, &pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00})
	client.AddOrder(ctx, &pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00})

	retrievedOrder , err := client.GetOrder(ctx, &wrappers.StringValue{Value: "102"})
	log.Print("GetOrder Response -> : ", retrievedOrder)
	retrievedOrder , err = client.GetOrder(ctx, &wrappers.StringValue{Value: "103"})
	log.Print("GetOrder Response -> : ", retrievedOrder)
	retrievedOrder , err = client.GetOrder(ctx, &wrappers.StringValue{Value: "104"})
	log.Print("GetOrder Response -> : ", retrievedOrder)

	updOrder1 := pb.Order{Id: "102", Items:[]string{"Google Pixel 3A", "Google Pixel Book"}, Destination:"Mountain View, CA", Price:1100.00}
	updOrder2 := pb.Order{Id: "103", Items:[]string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination:"San Jose, CA", Price:2800.00}
	updOrder3 := pb.Order{Id: "104", Items:[]string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination:"Mountain View, CA", Price:2200.00}

	updateStream, err := client.UpdateOrders(ctx)

	if err != nil {
		log.Fatalf("%v.updateOrders(_) = _, %v", client, err)
	}

	if err := updateStream.Send(&updOrder1); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder1, err)
	}
	if err := updateStream.Send(&updOrder2); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder2, err)
	}
	if err := updateStream.Send(&updOrder3); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder3, err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	}

	log.Printf("Update Orders Res : %s", updateRes)

	retrievedOrder , err = client.GetOrder(ctx, &wrappers.StringValue{Value: "102"})
	log.Print("GetOrder Response -> : ", retrievedOrder)
	retrievedOrder , err = client.GetOrder(ctx, &wrappers.StringValue{Value: "103"})
	log.Print("GetOrder Response -> : ", retrievedOrder)
	retrievedOrder , err = client.GetOrder(ctx, &wrappers.StringValue{Value: "104"})
	log.Print("GetOrder Response -> : ", retrievedOrder)

}
