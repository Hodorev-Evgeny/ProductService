package feature_product_service

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
)

func (s *ProductService) GetItemById(
	ctx context.Context,
	productID int64,
) (core_domain.Product, error) {
	if productID == core_domain.UnanalyzedID {
		return core_domain.Product{}, fmt.Errorf("product id cannot be un-analyzed")
	}

	product, err := s.repository.GetItemById(ctx, productID)
	if err != nil {
		return core_domain.Product{}, fmt.Errorf("repository.GetItemById: %w", err)
	}

	return product, nil
}
