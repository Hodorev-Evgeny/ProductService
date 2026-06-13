package server

import (
	core_logger "ProductService/internal/core/logger"
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type OrderCase interface {
}

type Server struct {
	pb.UnimplementedProductServiceServer
	orderCase OrderCase
	config    ServerConfig
}

func NewServer(config ServerConfig, ord OrderCase) *Server {
	return &Server{
		config:    config,
		orderCase: ord,
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
