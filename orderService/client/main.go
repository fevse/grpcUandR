package main

import (
	"context"
	"log"
	pb "orderService/client/orderService"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	order1 := pb.Order{Id: "11", Items: []string{"Mooer Micro Looper"}, Destination: "Huul", Price: 45.0}
	res, _ := client.AddOrder(ctx, &order1)
	if res != nil {
		log.Print("AddOrder Response -> : ", res.Value)
	}

	retrievedOrder, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: "15"})
	log.Print("GetOrder Response -> : ", retrievedOrder)

}
