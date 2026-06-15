package feature_product_service

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
)

func (s *ProductService) Deactivate(
	ctx context.Context,
	id int64,
) (core_domain.Product, error) {
	if id == core_domain.UnanalyzedID {
		return core_domain.Product{}, fmt.Errorf("unanalyzed product id")
	}

	domain, err := s.repository.Deactivate(ctx, id)
	if err != nil {
		return core_domain.Product{}, fmt.Errorf("error product deactivate: %w", err)
	}

	return domain, nil
}
