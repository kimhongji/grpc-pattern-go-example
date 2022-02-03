package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-pattern-go-example/proto/ecommerce"
	"log"
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