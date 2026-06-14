package feature_product_transport

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

func (c *ProductFeatureCase) CreateItem(
	ctx context.Context,
	product *pb.ProductRequest,
) (*pb.ProductResponse, error) {
	product_domain, err := c.service.CreateItem(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("create item: %w", err)
	}

	resp := core_domain.DomainFromResponse(product_domain)
	return resp, nil
}
