package dto

import (
	"errors"
	"strings"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

type CategoryDto struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (product *CategoryDto) Validate() error {
	if len(product.Name) < 3 || len(strings.TrimSpace(product.Name)) < 3 {
		return errors.New("'Name' should have at least 3 characters")
	}

	return nil
}

func (category *CategoryDto) Normalize() {
	category.Name = strings.TrimSpace(category.Name)
}

func MapToCategoryDto(category entities.Category) *CategoryDto {
	return &CategoryDto{
		Id:   category.Id,
		Name: category.Name,
	}
}
