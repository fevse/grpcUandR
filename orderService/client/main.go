package main

import (
	"context"
	"io"
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

	order1 := pb.Order{Id: "11", Items: []string{"Mooer Micro Looper"}, Destination: "Balmora", Price: 45.0}
	res, err := client.AddOrder(ctx, &order1)
	if err != nil {
		log.Fatalf("cannot add order: %v", err)
	}
	if res != nil {
		log.Print("AddOrder Response -> : ", res.Value)
	}

	retrievedOrder, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: "15"})
	if err != nil {
		log.Fatalf("cannot get order: %v", err)
	}
	log.Print("GetOrder Response -> : ", retrievedOrder)

	searchStream, err := client.SearchOrders(ctx, &wrapperspb.StringValue{Value: "Boss"})
	if err != nil {
		log.Fatalf("cannot search orders: %v", err)
	}

	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("cannot search order: %v", err)
		}
		log.Printf("Search Result : %v", searchOrder)
	}

	updOrder1 := pb.Order{Id: "12", Items: []string{"Coca-Cola Zero", "Big Mac"}, Destination: "Batumi"}
	updOrder2 := pb.Order{Id: "14", Items: []string{"Sofa", "Table", "Chair"}, Destination: "Balmora"}
	updOrder3 := pb.Order{Id: "16", Items: []string{"Holy Grail"}, Destination: "Erathia"}

	updateStream, err := client.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("cannot update orders: %v", err)
	}

	if err := updateStream.Send(&updOrder1); err != nil {
		log.Fatalf("cannot send order: %v", err)
	}
	if err := updateStream.Send(&updOrder2); err != nil {
		log.Fatalf("cannot send order: %v", err)
	}
	if err := updateStream.Send(&updOrder3); err != nil {
		log.Fatalf("cannot send order: %v", err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("cannot close and recive: %v", err)
	}
	log.Printf("Update Orders Res : %s", updateRes)

	streamProcOrder, err := client.ProcessOrders(ctx)
	if err != nil {
		log.Fatalf("cannot process orders: %v", err)
	}
	if err := streamProcOrder.Send(&wrapperspb.StringValue{Value: "12"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", client, "12", err)
	}
	if err := streamProcOrder.Send(&wrapperspb.StringValue{Value: "13"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", client, "13", err)
	}
	if err := streamProcOrder.Send(&wrapperspb.StringValue{Value: "14"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", client, "14", err)
	}

	ch := make(chan struct{})

	go asncClientBidirectionalRPC(streamProcOrder, ch)
	time.Sleep(1 * time.Second)

	if err := streamProcOrder.Send(&wrapperspb.StringValue{Value: "11"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", client, "14", err)
	}

	if err := streamProcOrder.CloseSend(); err != nil {
		log.Fatal(err)
	}

	ch <- struct{}{}
}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan struct{}) {
	for {
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder == io.EOF {
			break
		}
		log.Printf("Combined shipment : %v", combinedShipment.OrderList)
	}
	<-c
}
