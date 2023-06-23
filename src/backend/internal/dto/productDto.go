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
	idObj, err := NewIntIdObject(product.Id.Value())
	if err != nil {
		panic(err)
	}

	return &ProductDto{
		Id:          *idObj,
		Name:        product.Name.Value(),
		Description: product.Description.Value(),
		Price:       product.Price.Value().StringFixedBank(2),
	}
}

func MapToProductDetailsDto(product entities.Product) *ProductDetailsDto {
	idObj, err := NewIntIdObject(product.Id.Value())
	if err != nil {
		panic(err)
	}

	categoryId, errCategory := NewIntIdObject(product.Category.Id.Value())
	if errCategory != nil {
		panic(errCategory)
	}

	return &ProductDetailsDto{
		Id:          *idObj,
		Name:        product.Name.Value(),
		Description: product.Description.Value(),
		Price:       product.Price.Value().StringFixedBank(2),
		Deleted:     product.Deleted,
		Category: CategoryDto{
			Id:   *categoryId,
			Name: product.Category.Name.Value(),
		},
	}
}
