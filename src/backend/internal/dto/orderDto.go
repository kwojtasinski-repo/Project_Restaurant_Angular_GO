package dto

import (
	"time"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

type OrderDetailsDto struct {
	Id            int64
	OrderNumber   string
	Price         string
	Created       time.Time
	Modified      *time.Time
	OrderProducts []OrderProductDto
}

func MapToOrderDetailsDto(order entities.Order) *OrderDetailsDto {
	return &OrderDetailsDto{
		Id:            order.Id.Value(),
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
