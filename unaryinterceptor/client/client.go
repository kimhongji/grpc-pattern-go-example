package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"grpc-pattern-go-example/proto/ecommerce"
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

	client := ecommerce.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)

	defer cancel()

	client.AddOrder(ctx, &ecommerce.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00})
	client.AddOrder(ctx, &ecommerce.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00})
	client.AddOrder(ctx, &ecommerce.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00})
	client.AddOrder(ctx, &ecommerce.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00})
	client.AddOrder(ctx, &ecommerce.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00})

	// Search Order : Server streaming scenario
	//searchStream, _ := client.SearchOrders(ctx, &wrappers.StringValue{Value: "Google"})
	//for {
	//	searchOrder, err := searchStream.Recv()
	//	if err == io.EOF {
	//		log.Print("EOF")
	//		break
	//	}
	//
	//	if err == nil {
	//		log.Print("Search Result : ", searchOrder)
	//	}
	//}
}