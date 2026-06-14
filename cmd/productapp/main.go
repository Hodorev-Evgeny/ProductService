package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	core_logger "github.com/Hodorev-Evgeny/ProductService/internal/core/logger"
	core_pgx_pool "github.com/Hodorev-Evgeny/ProductService/internal/core/repository/postgres/pgx"
	"github.com/Hodorev-Evgeny/ProductService/internal/core/transport/server"
	feature_product_repository "github.com/Hodorev-Evgeny/ProductService/internal/feature/product/repository"
	feature_product_service "github.com/Hodorev-Evgeny/ProductService/internal/feature/product/service"
	feature_product_transport "github.com/Hodorev-Evgeny/ProductService/internal/feature/product/transport"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	logger_config := core_logger.MustNewConfig()
	logger, err := core_logger.NewLogger(logger_config)
	if err != nil {
		fmt.Println("Error initializing logger")
	}
	ctx = core_logger.ToContext(ctx, logger)

	configPool := core_pgx_pool.MustPostgresConfig()
	pool := core_pgx_pool.CreatePoolMust(ctx, configPool)

	productRepository := feature_product_repository.NewProductRepository(pool)
	productService := feature_product_service.NewProductService(productRepository)
	productCase := feature_product_transport.NewProductFeatureCase(productService)

	server_config := server.MustGetServerConfig()
	server := server.NewServer(
		server_config,
		productCase,
	)

	logger.Info("Starting server")
	if err := server.Start(ctx); err != nil {
		logger.Error("Error starting server", zap.Error(err))
	}
}
