package feature_product_repository

import (
	core_repository_pool "github.com/Hodorev-Evgeny/ProductService/internal/core/repository/postgres"
)

type ProductRepository struct {
	pool core_repository_pool.Pool
}

func NewProductRepository(
	pool core_repository_pool.Pool,
) *ProductRepository {
	return &ProductRepository{
		pool: pool,
	}
}
