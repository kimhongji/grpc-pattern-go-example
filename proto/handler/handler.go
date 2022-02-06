package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-pattern-go-example/proto/ecommerce"
	"io"
	"log"
	"strings"
	"time"
)

type Server struct {
	orderMap map[string]*pb.Order
}

func (s *Server) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*pb.Order, error){
	ord, exists := s.orderMap[orderId.Value]
	if exists {
		return ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist. : ", orderId)
}

func (s *Server) AddOrder(ctx context.Context, order *pb.Order) (*wrappers.StringValue, error) {
	log.Printf("Order Added. ID : %v", order.Id)

	if s.orderMap == nil {
		s.orderMap = make(map[string]*pb.Order)
	}

	s.orderMap[order.Id] = order

	return &wrappers.StringValue{Value: "Order Added: " + order.Id}, nil
}

func (s *Server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range s.orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf("error sending massage to stream : %v", err)
				}
				log.Print("Matching Order Found: " + key)
				time.Sleep(500 * time.Millisecond)
				break
			}
		}
	}
	return nil
}

func (s *Server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	updatedIds := "Updated Order Ids: "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrappers.StringValue{Value: updatedIds})
		}
		if err != nil {
			return err
		}
		s.orderMap[order.Id] = order

		log.Printf("Order ID : %s - %s", order.Id, "Updated")
		updatedIds += order.Id + ", "
	}
}