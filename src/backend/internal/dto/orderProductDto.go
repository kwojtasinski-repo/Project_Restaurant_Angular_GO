package dto

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

type OrderProductDto struct {
	Id        int64
	Name      string
	Price     string
	ProductId int64
}

func mapToOrderProductDto(orderProduct entities.OrderProduct) *OrderProductDto {
	return &OrderProductDto{
		Id:        orderProduct.Id.Value(),
		Name:      orderProduct.Name.Value(),
		Price:     orderProduct.Price.Value().StringFixedBank(2),
		ProductId: orderProduct.ProductId.Value(),
	}
}