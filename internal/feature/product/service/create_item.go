package feature_product_service

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

func (s *ProductService) CreateItem(
	ctx context.Context,
	request *pb.ProductRequest,
) (core_domain.Product, error) {
	domain := core_domain.CreateProduct(
		request.Name,
		request.Description,
		request.Price,
	)
	// добавить проверку на валидность

	product, err := s.repository.CreateItem(ctx, domain)
	if err != nil {
		return core_domain.Product{}, fmt.Errorf("repository.CreateItem: %w", err)
	}

	return product, nil
}
