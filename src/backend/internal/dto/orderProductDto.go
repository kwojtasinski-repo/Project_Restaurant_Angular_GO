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
	idObj, err := NewIntIdObject(orderProduct.Id.Value())
	if err != nil {
		panic(err)
	}

	productIdObj, errProductId := NewIntIdObject(orderProduct.ProductId.Value())
	if errProductId != nil {
		panic(errProductId)
	}

	return &OrderProductDto{
		Id:        *idObj,
		Name:      orderProduct.Name.Value(),
		Price:     orderProduct.Price.Value().StringFixedBank(2),
		ProductId: *productIdObj,
	}
}
