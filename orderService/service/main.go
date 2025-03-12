//go:generate protoc --proto_path=proto --go_out=. --go-grpc_out=. proto/*.proto
package main

import (
	"context"
	"log"
	"net"
	pb "orderService/service/orderService"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const port = ":50051"

var orderMap = make(map[string]*pb.Order, 0)

type server struct {
	pb.UnimplementedOrderManagementServer
	orderMap map[string]*pb.Order
}

func (s *server) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	ord := orderMap[orderId.Value]
	return ord, nil
}

func (s *server) AddOrder(ctx context.Context, order *pb.Order) (*wrapperspb.StringValue, error) {
	orderMap[order.Id] = order
	log.Printf("Order %v added", order.Id)
	return &wrapperspb.StringValue{Value: "Order " + order.Id + " added"}, nil
}

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	orderMap["12"] = &pb.Order{Id: "12", Items: []string{"Fender Telecaster", "Fender Blues Junior"}, Destination: "Balmora", Price: 2500.00}
	orderMap["13"] = &pb.Order{Id: "13", Items: []string{"Boss BD-2"}, Destination: "Sadrit Mora", Price: 140.00}
	orderMap["14"] = &pb.Order{Id: "14", Items: []string{"Gibson Les Paul", "Roland Space Echo RE-201"}, Destination: "Vivek", Price: 3400.00}
	orderMap["15"] = &pb.Order{Id: "15", Items: []string{"Behringer Model-D"}, Destination: "Seida Nin", Price: 330.00}
	orderMap["16"] = &pb.Order{Id: "16", Items: []string{"Squier Jazzmaster", "Eventide Space", "EHX Big Muff"}, Destination: "Aldrun", Price: 920.00}
}
