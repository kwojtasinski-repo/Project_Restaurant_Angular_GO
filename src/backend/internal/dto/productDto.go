package dto

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/shopspring/decimal"
)

type ProductDto struct {
	Id          int64
	Name        string
	Description string
	Price       decimal.Decimal
}

type ProductDetailsDto struct {
	Id          int64
	Name        string
	Description string
	Price       decimal.Decimal
	Category    CategoryDto
}

func MapToProductDto(product entities.Product) *ProductDto {
	return &ProductDto{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}

func MapToProductDetailsDto(product entities.Product) *ProductDetailsDto {
	return &ProductDetailsDto{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category: CategoryDto{
			Id:   product.Category.Id,
			Name: product.Category.Name,
		},
	}
}
