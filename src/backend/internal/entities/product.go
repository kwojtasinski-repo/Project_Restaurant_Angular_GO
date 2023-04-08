package entities

import (
	"errors"
	"strings"

	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/shopspring/decimal"
)

type Product struct {
	Id          valueobjects.Id
	Name        valueobjects.Name
	Description valueobjects.Description
	Price       valueobjects.Price
	Category    Category
	Deleted     bool // soft delete
}

func NewProduct(id int64, name string, price decimal.Decimal, description string, category *Category) (*Product, error) {
	var validationErrors strings.Builder
	var err error
	product := &Product{}
	var newId *valueobjects.Id
	var newName *valueobjects.Name
	var newDescription *valueobjects.Description
	var newPrice *valueobjects.Price
	newId, err = valueobjects.NewId(id)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}
	newName, err = valueobjects.NewName(name)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}
	newDescription, err = valueobjects.NewDescription(description)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}
	newPrice, err = valueobjects.NewPrice(price)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}
	if validationErrors.Len() > 0 {
		return nil, errors.New(validationErrors.String())
	}
	product.Id = *newId
	product.Name = *newName
	product.Description = *newDescription
	product.Price = *newPrice
	product.Category = *category
	product.Deleted = false
	return product, nil
}
