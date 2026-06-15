package feature_product_repository

import (
	"context"
	"fmt"
	"time"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
)

func (r *ProductRepository) CreateItem(
	ctx context.Context,
	request core_domain.Product,
) (core_domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO service_product.product (name, description, price, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, description, price, status;`

	row := r.pool.QueryRow(ctx, query,
		request.Name,
		request.Description,
		request.Price,
		request.Status)

	var domain_product core_domain.Product
	err := row.Scan(
		&domain_product.ID,
		&domain_product.Name,
		&domain_product.Description,
		&domain_product.Price,
		&domain_product.Status,
	)
	if err != nil {
		return core_domain.Product{}, fmt.Errorf("error creating order: %w", err)
	}

	return domain_product, nil
}
