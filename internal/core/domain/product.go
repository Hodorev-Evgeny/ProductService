package core_domain

import (
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

type ProductStatus int

var UnanalyzedID int64 = -1

const (
	ACTIVE ProductStatus = iota
	INACTIVE
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       int64
	Status      ProductStatus
}

func CreateProduct(
	name string,
	description string,
	price int64,
) Product {
	return Product{
		ID:          UnanalyzedID,
		Name:        name,
		Description: description,
		Price:       price,
		Status:      ACTIVE,
	}
}

func (o *Product) StatusToResponse() pb.StatusProduct {
	var status pb.StatusProduct
	switch o.Status {
	case ACTIVE:
		status = pb.StatusProduct_ACTIVE
	case INACTIVE:
		status = pb.StatusProduct_INACTIVE
	}

	return status
}

func DomainFromResponse(product Product) *pb.ProductResponse {
	return &pb.ProductResponse{
		Id:          &pb.ProductID{Id: product.ID},
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		IsActive:    product.StatusToResponse(),
	}
}
