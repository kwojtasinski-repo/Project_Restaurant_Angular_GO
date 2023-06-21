package dto

import (
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
)

type OrderProductDto struct {
	Id        IdObject `json:"id"`
	Name      string   `json:"name"`
	Price     string   `json:"price"`
	ProductId IdObject `json:"productId"`
}

func mapToOrderProductDto(orderProduct entities.OrderProduct) *OrderProductDto {
	return &OrderProductDto{
		Id:        IdObject{ValueInt: orderProduct.Id.Value()},
		Name:      orderProduct.Name.Value(),
		Price:     orderProduct.Price.Value().StringFixedBank(2),
		ProductId: IdObject{ValueInt: orderProduct.ProductId.Value()},
	}
}
