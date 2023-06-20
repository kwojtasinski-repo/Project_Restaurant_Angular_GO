package dto

import (
	"errors"
	"strings"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
)

type CategoryDto struct {
	Id   IdObject `json:"id"`
	Name string   `json:"name"`
}

func (product *CategoryDto) Validate() error {
	var validationErrors strings.Builder
	if len(product.Name) < 3 || len(strings.TrimSpace(product.Name)) < 3 {
		validationErrors.WriteString("'Name' should have at least 3 characters. ")
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

type CategoryDetailsDto struct {
	Id      IdObject `json:"id"`
	Name    string   `json:"name"`
	Deleted bool     `json:"deleted"`
}

func MapToCategoryDto(category entities.Category) *CategoryDto {
	return &CategoryDto{
		Id:   IdObject{ValueInt: category.Id.Value()},
		Name: category.Name.Value(),
	}
}

func MapToCategoryDetailsDto(category entities.Category) *CategoryDetailsDto {
	return &CategoryDetailsDto{
		Id:      IdObject{ValueInt: category.Id.Value()},
		Name:    category.Name.Value(),
		Deleted: category.Deleted,
	}
}
