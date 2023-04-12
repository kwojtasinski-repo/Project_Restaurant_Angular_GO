package dto

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

type ProductDto struct {
	Id          int64
	Name        string
	Description string
	Price       string
}

type ProductDetailsDto struct {
	Id          int64
	Name        string
	Description string
	Price       string
	Category    CategoryDto
	Deleted     bool
}

func MapToProductDto(product entities.Product) *ProductDto {
	return &ProductDto{
		Id:          product.Id.Value(),
		Name:        product.Name.Value(),
		Description: product.Description.Value(),
		Price:       product.Price.Value().StringFixedBank(2),
	}
}

func MapToProductDetailsDto(product entities.Product) *ProductDetailsDto {
	return &ProductDetailsDto{
		Id:          product.Id.Value(),
		Name:        product.Name.Value(),
		Description: product.Description.Value(),
		Price:       product.Price.Value().StringFixedBank(2),
		Deleted:     product.Deleted,
		Category: CategoryDto{
			Id:   product.Category.Id.Value(),
			Name: product.Category.Name.Value(),
		},
	}
}
