package feature_product_service

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
)

func (s *ProductService) GetAllItems(
	ctx context.Context,
) ([]core_domain.Product, error) {
	list, err := s.repository.GetAllItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("repository GetAllItems: %w", err)
	}

	return list, nil
}
