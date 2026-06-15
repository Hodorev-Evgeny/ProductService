package feature_product_service

import (
	"context"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
)

type ProductRepository interface {
	CreateItem(
		ctx context.Context,
		request core_domain.Product,
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

type ProductService struct {
	repository ProductRepository
}

func NewProductService(
	repository ProductRepository,
) *ProductService {
	return &ProductService{
		repository: repository,
	}
}
