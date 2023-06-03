package dto

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

type ProductDto struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type ProductDetailsDto struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       string      `json:"price"`
	Category    CategoryDto `json:"category"`
	Deleted     bool        `json:"deleted"`
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
