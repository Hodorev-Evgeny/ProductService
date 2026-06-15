package feature_product_transport

import (
	"context"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

type ProductService interface {
	CreateItem(
		ctx context.Context,
		request *pb.ProductRequest,
	) (core_domain.Product, error)

	GetItemById(
		ctx context.Context,
		productID int64,
	) (core_domain.Product, error)

	GetAllItems(
		ctx context.Context,
	) ([]core_domain.Product, error)

	UpdatePrice(
		ctx context.Context,
		id int64,
		price int64,
	) (core_domain.Product, error)

	Deactivate(
		ctx context.Context,
		id int64,
	) (core_domain.Product, error)
}

type ProductFeatureCase struct {
	service ProductService
}

func NewProductFeatureCase(
	service ProductService,
) *ProductFeatureCase {
	return &ProductFeatureCase{
		service: service,
	}
}
