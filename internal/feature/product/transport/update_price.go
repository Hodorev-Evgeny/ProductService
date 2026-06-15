package feature_product_transport

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

func (c *ProductFeatureCase) UpdatePrice(
	ctx context.Context,
	updateProduct *pb.UpdatePriceRequest,
) (*pb.ProductResponse, error) {
	domain, err := c.service.UpdatePrice(ctx, updateProduct.ProductId.Id, updateProduct.NewPrice)
	if err != nil {
		return nil, fmt.Errorf("product update price %w", err)
	}

	resp := core_domain.DomainFromResponse(domain)
	return resp, nil
}
