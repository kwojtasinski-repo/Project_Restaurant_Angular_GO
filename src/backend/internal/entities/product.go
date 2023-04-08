package entities

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

type Product struct {
	id          int64
	name        string
	description string
	price       decimal.Decimal
	category    Category
	deleted     bool // soft delete
}

func NewProduct(id int64, name string, price decimal.Decimal, description string, category *Category) (*Product, error) {
	if err := validate(name, price, description, category); err != nil {
		return nil, err
	}

	product := &Product{
		id: id,
	}
	product.SetName(name)
	product.SetPrice(price)
	product.SetDescription(description)
	product.SetCategory(category)
	product.SetDeleted(false)
	return product, nil
}

func (product *Product) Id() int64 {
	return product.id
}

func (product *Product) Name() string {
	return product.name
}

func (product *Product) Description() string {
	return product.description
}

func (product *Product) Price() decimal.Decimal {
	return product.price
}

func (product *Product) Category() Category {
	return product.category
}

func (product *Product) Deleted() bool {
	return product.deleted
}

func (product *Product) SetName(name string) error {
	nameTrimmed := strings.TrimSpace(name)
	if err := validateName(nameTrimmed); err != nil {
		return err
	}

	product.name = nameTrimmed
	return nil
}

func (product *Product) SetPrice(price decimal.Decimal) error {
	if err := validatePrice(price); err != nil {
		return err
	}

	product.price = price
	return nil
}

func (product *Product) SetDescription(description string) error {
	descriptionTrimmed := strings.TrimSpace(description)
	if err := validateDescription(descriptionTrimmed); err != nil {
		return err
	}

	product.description = descriptionTrimmed
	return nil
}

func (product *Product) SetCategory(category *Category) error {
	if err := validateCategory(category); err != nil {
		return err
	}

	product.category = *category
	return nil
}

func (product *Product) SetDeleted(deleted bool) {
	product.deleted = deleted
}

func (product *Product) SetId(id int64) {
	product.id = id
}

func validate(name string, price decimal.Decimal, description string, category *Category) error {
	var validationErrors strings.Builder
	err := validateName(name)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}

	err = validateDescription(description)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}

	err = validatePrice(price)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}

	err = validateCategory(category)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}

	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	} else {
		return nil
	}
}

func validateName(name string) error {
	var validationErrors strings.Builder
	if len(name) < 3 || len(strings.TrimSpace(name)) < 3 {
		validationErrors.WriteString("'Name' should have at least 3 characters. ")
	}
	if len(strings.TrimSpace(name)) > 200 {
		validationErrors.WriteString("'Name' cannot have more than 200 characters. ")
	}

	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	} else {
		return nil
	}
}

func validateDescription(description string) error {
	var validationErrors strings.Builder
	if len(strings.TrimSpace(description)) > 5000 {
		validationErrors.WriteString("'Description' cannot have more than 5000 characters. ")
	}
	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	} else {
		return nil
	}
}

func validatePrice(price decimal.Decimal) error {
	var validationErrors strings.Builder
	if price.LessThan(decimal.Zero) {
		validationErrors.WriteString("'Price' cannot be negative. ")
	}
	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	} else {
		return nil
	}
}

func validateCategory(category *Category) error {
	if category == nil {
		return errors.New("'Category' cannot be empty")
	}
	return nil
}
