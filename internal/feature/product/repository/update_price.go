package feature_product_repository

import (
	"context"
	"fmt"
	"time"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
)

func (r *ProductRepository) UpdatePrice(
	ctx context.Context,
	id int64,
	price int64,
) (core_domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		UPDATE service_product.product
		SET price = $1
		WHERE id = $2;`

	row := r.pool.QueryRow(ctx, query, price, id)

	var domain_product core_domain.Product
	err := row.Scan(
		&domain_product.ID,
		&domain_product.Name,
		&domain_product.Description,
		&domain_product.Price,
		&domain_product.Status,
	)
	if err != nil {
		return core_domain.Product{}, fmt.Errorf("error scan order: %w", err)
	}

	return domain_product, nil
}
