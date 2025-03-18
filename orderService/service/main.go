//go:generate protoc --proto_path=proto --go_out=. --go-grpc_out=. proto/*.proto
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	pb "orderService/service/orderService"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const port = ":50051"
const orderBatchSize = 3

var orderMap = make(map[string]*pb.Order, 0)

type wrappedStream struct {
	grpc.ServerStream
}

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

	log.Println("Sleeping ...")
	time.Sleep(time.Duration(5 * time.Second))

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("Deadline %s", ctx.Err())
		return nil, ctx.Err()
	}

	log.Printf("Order %v added", order.Id)
	return &wrapperspb.StringValue{Value: "Order " + order.Id + " added"}, nil
}

func (s *server) SearchOrders(searchQuery *wrapperspb.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range orderMap {
		// log.Printf("#%v : %v", key, order)
		for _, itemStr := range order.Items {
			if strings.Contains(itemStr, searchQuery.Value) {
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf("error sending message to stream: %v", err)
				}
				log.Printf("Matching Order Found : %v", key)
				break
			}
		}
	}
	return nil
}

func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	oStr := strings.Builder{}
	oStr.WriteString("Update Order IDs : ")
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrapperspb.StringValue{Value: "Orders processed " + oStr.String()})
		}
		orderMap[order.Id] = order

		log.Printf("Order #%v : Updated", order.Id)
		oStr.WriteString(order.Id + ", ")
	}
}

func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	batchMarker := 1
	var combinedShipmentMap = make(map[string]*pb.CombinedShipment)
	for {
		orderId, err := stream.Recv()
		if err == io.EOF {
			for _, ship := range combinedShipmentMap {
				if err := stream.Send(ship); err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			log.Print(err)
			return err
		}

		destination := orderMap[orderId.GetValue()].Destination
		shipment, ok := combinedShipmentMap[destination]

		if ok {
			ord := orderMap[orderId.GetValue()]
			shipment.OrderList = append(shipment.OrderList, ord)
			combinedShipmentMap[destination] = shipment
		} else {
			comShip := &pb.CombinedShipment{Id: orderMap[orderId.GetValue()].Destination, Status: "In progress"}
			ord := orderMap[orderId.GetValue()]
			comShip.OrderList = append(comShip.OrderList, ord)
			combinedShipmentMap[destination] = comShip
		}

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrderList))
				if err := stream.Send(comb); err != nil {
					return err
				}
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]*pb.CombinedShipment)
		} else {
			batchMarker++
		}
	}
}

func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("*** [Server Unary Interceptor] ", info.FullMethod)
	log.Printf("Before handling the request: %s", req)

	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("Error handling the request: %s", err)
	}
	log.Printf("After handling the request: %s", m)
	return m, err
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("*** [Server Stream Interceptor Wrapper] Received message: %T", m)
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("*** [Server Stream Interceptor Wrapper] Sending message: %T", m)
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) *wrappedStream {
	return &wrappedStream{s}
}

func orederStreamServerInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("*** [Server Stream Interceptor] ", info.FullMethod)
	err := handler(srv, newWrappedStream(stream))
	if err != nil {
		log.Printf("RPC error %v", err)
	}
	return err
}

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(orderUnaryServerInterceptor),
		grpc.StreamInterceptor(orederStreamServerInterceptor),
	)
	pb.RegisterOrderManagementServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	orderMap["12"] = &pb.Order{Id: "12", Items: []string{"Fender Telecaster", "Fender Blues Junior"}, Destination: "Balmora", Price: 2500.00}
	orderMap["13"] = &pb.Order{Id: "13", Items: []string{"Boss BD-2"}, Destination: "Balmora", Price: 140.00}
	orderMap["14"] = &pb.Order{Id: "14", Items: []string{"Gibson Les Paul", "Roland Space Echo RE-201"}, Destination: "Balmora", Price: 3400.00}
	orderMap["15"] = &pb.Order{Id: "15", Items: []string{"Behringer Model-D"}, Destination: "Seyda Neen", Price: 330.00}
	orderMap["16"] = &pb.Order{Id: "16", Items: []string{"Squier Jazzmaster", "Eventide Space", "Boss RV-5"}, Destination: "Balmora", Price: 920.00}
}
