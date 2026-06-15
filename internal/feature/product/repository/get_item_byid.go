package feature_product_repository

import (
	"context"
	"fmt"
	"time"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
)

func (r *ProductRepository) GetItemById(
	ctx context.Context,
	productID int64,
) (core_domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		SELECT id, name, description, price, status
		FROM service_product.product
		WHERE id = $1;`

	row := r.pool.QueryRow(ctx, query, productID)

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
