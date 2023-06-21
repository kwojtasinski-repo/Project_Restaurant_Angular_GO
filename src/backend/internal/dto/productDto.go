package dto

import (
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
)

type ProductDto struct {
	Id          IdObject `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       string   `json:"price"`
}

type ProductDetailsDto struct {
	Id          IdObject    `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       string      `json:"price"`
	Category    CategoryDto `json:"category"`
	Deleted     bool        `json:"deleted"`
}

func MapToProductDto(product entities.Product) *ProductDto {
	return &ProductDto{
		Id:          IdObject{ValueInt: product.Id.Value()},
		Name:        product.Name.Value(),
		Description: product.Description.Value(),
		Price:       product.Price.Value().StringFixedBank(2),
	}
}

func MapToProductDetailsDto(product entities.Product) *ProductDetailsDto {
	return &ProductDetailsDto{
		Id:          IdObject{ValueInt: product.Id.Value()},
		Name:        product.Name.Value(),
		Description: product.Description.Value(),
		Price:       product.Price.Value().StringFixedBank(2),
		Deleted:     product.Deleted,
		Category: CategoryDto{
			Id:   IdObject{ValueInt: product.Category.Id.Value()},
			Name: product.Category.Name.Value(),
		},
	}
}
