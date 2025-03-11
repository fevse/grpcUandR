package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "service/sappgrpc"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	pb.UnimplementedProductInfoServer
	DB   *mongo.Client
	Coll *mongo.Collection
}

func NewProductService() *ProductService {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your MONGODB_URI environment variable")
	}
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	coll := client.Database("fevse").Collection("storage")
	return &ProductService{
		DB:   client,
		Coll: coll,
	}
}

func (c *ProductService) Disconnect(ctx context.Context) error {
	if err := c.DB.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}

func (c *ProductService) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	var result pb.Product
	err := c.Coll.FindOne(ctx, bson.D{{Key: "id", Value: in.Value}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("No document was found with id: %s\n", in.Value)
	}
	if err != nil {
		panic(err)
	}
	return &result, nil
}

func (c *ProductService) AddProduct(ctx context.Context, req *pb.Product) (*pb.ProductID, error) {
	prod := &pb.Product{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID: %v", err)
	}
	prod.Id = out.String()
	_, err = c.Coll.InsertOne(ctx, prod)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add product: %v", err)
	}
	return &pb.ProductID{Value: prod.Id}, nil
}
