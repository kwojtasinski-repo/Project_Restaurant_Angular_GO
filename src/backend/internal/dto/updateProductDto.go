package dto

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

type UpdateProductDto struct {
	Id          IdObject        `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CategoryId  IdObject        `json:"categoryId"`
	Price       decimal.Decimal `json:"price"`
}

func (product *UpdateProductDto) Validate() error {
	var validationErrors strings.Builder
	if len(product.Name) < 3 || len(strings.TrimSpace(product.Name)) < 3 {
		validationErrors.WriteString("'Name' should have at least 3 characters. ")
	}

	if len(strings.TrimSpace(product.Name)) > 200 {
		validationErrors.WriteString("'Name' cannot have more than 200 characters. ")
	}

	if len(strings.TrimSpace(product.Description)) > 5000 {
		validationErrors.WriteString("'Description' cannot have more than 5000 characters. ")
	}

	if product.Price.LessThan(decimal.Zero) {
		validationErrors.WriteString("'Price' cannot be negative. ")
	}

	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	} else {
		return nil
	}
}
