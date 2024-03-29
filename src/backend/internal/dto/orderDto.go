package dto

import (
	"time"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
)

type OrderDto struct {
	Id          IdObject   `json:"id"`
	OrderNumber string     `json:"orderNumber"`
	Price       string     `json:"price"`
	Created     time.Time  `json:"created"`
	Modified    *time.Time `json:"modified"`
}

type OrderDetailsDto struct {
	Id            IdObject          `json:"id"`
	OrderNumber   string            `json:"orderNumber"`
	Price         string            `json:"price"`
	Created       time.Time         `json:"created"`
	Modified      *time.Time        `json:"modified"`
	OrderProducts []OrderProductDto `json:"orderProducts"`
}

func MapToOrderDto(order entities.Order) *OrderDto {
	idObj, err := NewIntIdObject(order.Id.Value())
	if err != nil {
		panic(err)
	}

	return &OrderDto{
		Id:          *idObj,
		OrderNumber: order.OrderNumber.Value(),
		Created:     order.Created,
		Modified:    order.Modified,
		Price:       order.Price.Value().StringFixedBank(2),
	}
}

func MapToOrderDetailsDto(order entities.Order) *OrderDetailsDto {
	idObj, err := NewIntIdObject(order.Id.Value())
	if err != nil {
		panic(err)
	}

	return &OrderDetailsDto{
		Id:            *idObj,
		OrderNumber:   order.OrderNumber.Value(),
		Created:       order.Created,
		Modified:      order.Modified,
		Price:         order.Price.Value().StringFixedBank(2),
		OrderProducts: mapToOrderProductsDto(order.OrderProducts),
	}
}

func mapToOrderProductsDto(orderProducts []entities.OrderProduct) []OrderProductDto {
	dtos := make([]OrderProductDto, 0)
	for i := 0; i < len(orderProducts); i++ {
		dtos = append(dtos, *mapToOrderProductDto(orderProducts[i]))
	}

	return dtos
}
