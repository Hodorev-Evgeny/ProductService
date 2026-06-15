package feature_product_transport

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

func (c *ProductFeatureCase) GetItemById(
	ctx context.Context,
	productID *pb.ProductID,
) (*pb.ProductResponse, error) {
	product_domain, err := c.service.GetItemById(ctx, productID.Id)
	if err != nil {
		return nil, fmt.Errorf("could not get item by id: %w", err)
	}

	resp := core_domain.DomainFromResponse(product_domain)
	return resp, nil
}
