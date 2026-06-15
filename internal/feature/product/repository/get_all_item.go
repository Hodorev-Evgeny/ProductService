package feature_product_repository

import (
	"context"
	"fmt"
	"time"

	core_domain "github.com/Hodorev-Evgeny/ProductService/internal/core/domain"
)

func (r *ProductRepository) GetAllItems(
	ctx context.Context,
) ([]core_domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := `
		SELECT id, name, description, price, status
		FROM service_product.product;`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error read rows in data base: %w", err)
	}

	var list []core_domain.Product
	for rows.Next() {
		var item core_domain.Product
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Status,
		); err != nil {
			return nil, fmt.Errorf("error scan row in data base: %w", err)
		}

		list = append(list, item)
	}

	return list, nil
}
