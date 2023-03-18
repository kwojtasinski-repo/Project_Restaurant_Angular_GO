package dto

import (
	"errors"
	"strings"
)

type CategoryDto struct {
	Id   int64
	Name string
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
