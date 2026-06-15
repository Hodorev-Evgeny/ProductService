package server

import (
	"context"
	"fmt"
	"net"
	"time"

	core_logger "github.com/Hodorev-Evgeny/ProductService/internal/core/logger"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ProductCase interface {
	CreateItem(
		ctx context.Context,
		product *pb.ProductRequest,
	) (*pb.ProductResponse, error)

	GetItemById(
		ctx context.Context,
		productID *pb.ProductID,
	) (*pb.ProductResponse, error)

	UpdatePrice(
		ctx context.Context,
		updateProduct *pb.UpdatePriceRequest,
	) (*pb.ProductResponse, error)

	GetAllItems(
		ctx context.Context,
	) (*pb.ListProduct, error)

	Deactivate(
		ctx context.Context,
		id *pb.ProductID,
	) (*pb.ProductResponse, error)
}

type Server struct {
	pb.UnimplementedProductServiceServer
	productCase ProductCase
	config      ServerConfig
}

func NewServer(config ServerConfig, ord ProductCase) *Server {
	return &Server{
		config:      config,
		productCase: ord,
	}
}

func (s *Server) Start(ctx context.Context) error {
	log := core_logger.FromContext(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.config.ADDR))
	if err != nil {
		log.Error("Error starting server", zap.Error(err))
		return err
	}
	defer lis.Close()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProductServiceServer(grpcServer, s)

	// написать функцию дл реализации
	reflection.Register(grpcServer)

	ch := make(chan error)
	go func() {
		defer close(ch)

		err := grpcServer.Serve(lis)
		if err != nil {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		return err

	case <-ctx.Done():
		stop := make(chan struct{})
		go func() {
			grpcServer.GracefulStop()
			close(stop)
		}()
		select {
		case <-stop:
			log.Info("Server stopped")
			return nil
		case <-time.After(s.config.TIMEOUT):
			log.Error("Server timeout")
			grpcServer.GracefulStop()
		}
	}
	return nil
}
