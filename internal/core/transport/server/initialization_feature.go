package server

import (
	"context"

	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateItem(ctx context.Context, item *pb.ProductRequest) (*pb.ProductResponse, error) {
	return s.productCase.CreateItem(ctx, item)
}

func (s *Server) GetItemByID(ctx context.Context, id *pb.ProductID) (*pb.ProductResponse, error) {
	return s.productCase.GetItemById(ctx, id)
}

func (s *Server) GetAllItem(ctx context.Context, empty *emptypb.Empty) (*pb.ListProduct, error) {
	return &pb.ListProduct{}, nil
}

func (s *Server) UpdatePrice(ctx context.Context, item *pb.UpdatePriceRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{}, nil
}

func (s *Server) DeactivateItem(ctx context.Context, id *pb.ProductID) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{}, nil
}
