package dto

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

type UpdateProductDto struct {
	Id          int64
	Name        string
	Description string
	CategoryId  int64
	Price       decimal.Decimal
}

func (product *UpdateProductDto) Validate() error {
	if len(product.Name) < 3 || len(strings.TrimSpace(product.Name)) < 3 {
		return errors.New("'Name' should have at least 3 characters")
	}

	if len(strings.TrimSpace(product.Name)) > 200 {
		return errors.New("'Name' cannot have more than 200 characters")
	}

	if len(strings.TrimSpace(product.Description)) > 5000 {
		return errors.New("'Description' cannot have more than 5000 characters")
	}

	if product.Price.LessThan(decimal.Zero) {
		return errors.New("'Price' cannot be negative")
	}

	return nil
}

func (product *UpdateProductDto) Normalize() {
	product.Name = strings.TrimSpace(product.Name)
	product.Description = strings.TrimSpace(product.Description)
}
