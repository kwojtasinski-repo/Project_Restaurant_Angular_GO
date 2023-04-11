package dto

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/shopspring/decimal"
)

type OrderProductDto struct {
	Id        int64
	Name      string
	Price     decimal.Decimal
	ProductId int64
}

func mapToOrderProductDto(orderProduct entities.OrderProduct) *OrderProductDto {
	return &OrderProductDto{
		Id:        orderProduct.Id.Value(),
		Name:      orderProduct.Name.Value(),
		Price:     orderProduct.Price.Value(),
		ProductId: orderProduct.ProductId.Value(),
	}
}
