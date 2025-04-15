package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "orderService/client/orderService"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	hwpb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(orderUnaryClientInterceptor),
		grpc.WithStreamInterceptor(clientStreamInterceptor),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)

	// Metadata
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"kn", "vn",
	)
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)

	ctxA := metadata.AppendToOutgoingContext(mdCtx, "k1", "v1", "k1", "v2", "k2", "v3")

	var header, trailer metadata.MD

	// clientDeadline := time.Now().Add(time.Duration(2 * time.Second))
	// ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	// defer cancel()

	order1 := pb.Order{Id: "11", Items: []string{"Mooer Micro Looper"}, Destination: "Balmora", Price: 45.0}

	res, err := client.AddOrder(ctxA, &order1, grpc.Header(&header), grpc.Trailer(&trailer))

	if err != nil {
		got := status.Code(err)
		log.Fatalf("Error Occured -> addOrder : %v", got)
	}
	if res != nil {
		log.Print("AddOrder Response -> : ", res.Value)
	}

	if t, ok := header["timestamp"]; ok {
		log.Println("timestamp from header")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in header")
	}
	if l, ok := header["location"]; ok {
		log.Println("location from header")
		for i, e := range l {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("location expected but doesn't exist in header")
	}

	helloClient := hwpb.NewGreeterClient(conn)

	hwcCtx, hwcCancel := context.WithTimeout(context.Background(), time.Second)
	defer hwcCancel()

	helloResponse, err := helloClient.SayHello(hwcCtx, &hwpb.HelloRequest{Name: "Check!!! Check!!!"})
	if err != nil {
		log.Fatalf("orderManagementClient.SayHello(_) = _, %v", err)
	}
	fmt.Println("Greetinf: ", helloResponse.Message)

	// order2 := pb.Order{Id: "-1",
	// 	Items:       []string{"Bebida", "Arroz"},
	// 	Destination: "Moon"}
	// res, addOrderError := client.AddOrder(ctx, &order2)

	// if addOrderError != nil {
	// 	errorCode := status.Code(addOrderError)
	// 	if errorCode == codes.InvalidArgument {
	// 		log.Printf("Invalid Argument Error : %s", errorCode)
	// 		errorStatus := status.Convert(addOrderError)
	// 		for _, d := range errorStatus.Details() {
	// 			switch info := d.(type) {
	// 			case *epb.BadRequest_FieldViolation:
	// 				log.Printf("Request Field Invalid: %s", info)
	// 			default:
	// 				log.Printf("Unexpected error type: %s", info)
	// 			}
	// 		}
	// 	} else {
	// 		log.Printf("Unhandled error : %s", errorCode)
	// 	}
	// } else {
	// 	log.Print("AddOrder Response -> ", res.Value)
	// }

	// retrievedOrder, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: "15"})
	// if err != nil {
	// 	log.Fatalf("cannot get order: %v", err)
	// }
	// log.Print("GetOrder Response -> : ", retrievedOrder)

	searchStream, err := client.SearchOrders(ctxA, &wrapperspb.StringValue{Value: "Boss"})
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

	updateStream, err := client.UpdateOrders(mdCtx)
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

	// cancel

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

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

	ch := make(chan bool, 1)

	go asncClientBidirectionalRPC(streamProcOrder, ch)
	time.Sleep(1 * time.Second)

	// cancel()
	log.Printf("RPC status : %s", ctx.Err())

	// if err := streamProcOrder.Send(&wrapperspb.StringValue{Value: "11"}); err != nil {
	// 	log.Fatalf("%v.Send(%v) = %v", client, "11", err)
	// }

	if err := streamProcOrder.CloseSend(); err != nil {
		log.Fatal(err)
	}

	<-ch
}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan bool) {
	for {
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder == io.EOF {
			break
		} else if errProcOrder != nil {
			log.Printf("!Error receiving message %v", errProcOrder)
			break
		}
		log.Printf("Combined shipment : %v", combinedShipment.OrderList)
	}
	c <- true
}

func orderUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Println("Method : " + method)
	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Println(reply)

	return err
}

func clientStreamInterceptor(
	ctx context.Context, desc *grpc.StreamDesc,
	cc *grpc.ClientConn, method string,
	streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Println("*** [Clietn Stream Interceptor] ", method)
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}

type wrappedStream struct {
	grpc.ClientStream
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("*** [Client Stream Interceptor] Receive a message: %T", m)
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("*** [Client Stream Interceptor] Send a message: %T ", m)
	return w.ClientStream.SendMsg(m)
}
