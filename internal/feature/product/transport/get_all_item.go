package feature_product_transport

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

func (c *ProductFeatureCase) GetAllItems(
	ctx context.Context,
) (*pb.ListProduct, error) {
	list, err := c.service.GetAllItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get all items: %w", err)
	}

	resp := ListDomainToProto(list)
	return resp, nil
}

func ListDomainToProto(arr []core_domain.Product) *pb.ListProduct {
	arrProto := make([]*pb.ProductResponse, len(arr))
	for i, v := range arr {
		arrProto[i] = core_domain.DomainFromResponse(v)
	}
	return &pb.ListProduct{
		ListProduct: arrProto,
	}
}
