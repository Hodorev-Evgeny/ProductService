package core_domain

import (
	"time"

	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/order"
	product "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderStatus int

var UnanalyzedID int64 = -1

const (
	CREATED OrderStatus = iota
	CANCELLED
	PAID
)

type Order struct {
	ID          int64
	ProductID   int64
	Quantity    int64
	Total       int64
	Status      OrderStatus
	TimeCreated time.Time
}

func CreateOrder(
	productID int64,
	quantity int64,
	total int64,
) Order {
	return Order{
		ID:          UnanalyzedID,
		ProductID:   productID,
		Quantity:    quantity,
		Total:       total,
		Status:      CREATED,
		TimeCreated: time.Now(),
	}
}

func (o *Order) StatusToResponse() pb.StatusOrder {
	var status pb.StatusOrder
	switch o.Status {
	case CREATED:
		status = pb.StatusOrder_CREATED
	case CANCELLED:
		status = pb.StatusOrder_CANCELLED
	case PAID:
		status = pb.StatusOrder_PAID
	}

	return status
}

func DomainFromResponse(order Order) *pb.OrderResponse {
	return &pb.OrderResponse{
		Id:          &pb.OrderID{Id: order.ID},
		ProductId:   &product.ProductID{Id: order.ProductID},
		Quantity:    order.Quantity,
		TotalPrice:  order.Total,
		Status:      order.StatusToResponse(),
		TimeCreated: timestamppb.New(order.TimeCreated),
	}
}
