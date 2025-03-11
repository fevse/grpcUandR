//go:generate protoc --proto_path=proto --go_out=. --go-grpc_out=. proto/*.proto
package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "service/sappgrpc"
	"service/storage"

	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stor := storage.NewProductService()
	defer func() {
		if err := stor.Disconnect(ctx); err != nil {
			log.Fatalf("Failded to disconnect form MongoDB: %v", err)
		}
	}()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %w", err)
	}
	s := grpc.NewServer()
	server := storage.NewProductService()
	pb.RegisterProductInfoServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %w", err)
	}
}
