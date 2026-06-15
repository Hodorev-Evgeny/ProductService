package feature_product_transport

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

func (c *ProductFeatureCase) Deactivate(
	ctx context.Context,
	id *pb.ProductID,
) (*pb.ProductResponse, error) {
	domain, err := c.service.Deactivate(ctx, id.Id)
	if err != nil {
		return nil, fmt.Errorf("could not deactivate item: %w", err)
	}

	resp := core_domain.DomainFromResponse(domain)
	return resp, nil
}
