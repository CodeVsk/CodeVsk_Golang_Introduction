package server

import (
	"context"
	"errors"
	"net"

	"github.com/codevsk/grpc/pb"
	"google.golang.org/grpc"
)

type Product struct {
	ProductID     string
	Name          string
	Description   string
	Price         float64
	StockQuantity int32
}

type ProductService struct {
	pb.UnimplementedProductServiceServer

	products map[string]*Product
}

func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	s := grpc.NewServer()
	productService := NewProductService()
	pb.RegisterProductServiceServer(s, productService)
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

func NewProductService() *ProductService {
	return &ProductService{
		products: map[string]*Product{
			"12345": {
				ProductID:     "12345",
				Name:          "Sample Product",
				Description:   "This is a sample product description.",
				Price:         19.99,
				StockQuantity: 100,
			},
		},
	}
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	product, exists := s.products[req.ProductId]
	if !exists {
		return nil, errors.New("product not found")
	}

	res := &pb.ProductResponse{
		ProductId:     product.ProductID,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		StockQuantity: product.StockQuantity,
	}

	return res, nil
}
