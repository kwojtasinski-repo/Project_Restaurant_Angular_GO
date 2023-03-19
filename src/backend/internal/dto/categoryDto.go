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
	var validationErrors strings.Builder
	if len(product.Name) < 3 || len(strings.TrimSpace(product.Name)) < 3 {
		validationErrors.WriteString("'Name' should have at least 3 characters")
	}

	if len(strings.TrimSpace(product.Name)) > 200 {
		validationErrors.WriteString("'Name' cannot have more than 200 characters. ")
	}

	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	} else {
		return nil
	}
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
