package entities

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/shopspring/decimal"
)

type Product struct {
	Id          int64
	Name        string
	Description string
	Price       decimal.Decimal
	Category    Category
	Deleted     bool // soft delete
}

func (product *Product) MapToDto() *dto.ProductDto {
	return &dto.ProductDto{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}
}
