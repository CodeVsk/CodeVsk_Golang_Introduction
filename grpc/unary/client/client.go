package client

import (
	"context"
	"fmt"

	"github.com/codevsk/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	productClient := pb.NewProductServiceClient(conn)

	res, err := productClient.GetProduct(context.Background(), &pb.ProductRequest{
		ProductId: "12345",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product ID: %s\n", res.ProductId)
}
