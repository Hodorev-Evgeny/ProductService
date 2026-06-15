package feature_product_service

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
)

func (s *ProductService) UpdatePrice(
	ctx context.Context,
	id int64,
	price int64,
) (core_domain.Product, error) {
	if id == core_domain.UnanalyzedID {
		return core_domain.Product{}, fmt.Errorf("Product ID cannot be un-analyzed")
	}
	if price <= 0 {
		return core_domain.Product{}, fmt.Errorf("Price cannot be negative")
	}

	domain, err := s.repository.UpdatePrice(ctx, id, price)
	if err != nil {
		return core_domain.Product{}, fmt.Errorf("Error update product by id: %d", id)
	}

	return domain, nil
}
