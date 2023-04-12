package valueobjects

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

type Price struct {
	value decimal.Decimal
}

func NewPrice(price decimal.Decimal) (*Price, error) {
	if err := validatePrice(price); err != nil {
		return nil, err
	}
	return &Price{
		value: price,
	}, nil
}

func (price *Price) Value() decimal.Decimal {
	return price.value
}

func (price *Price) String() string {
	return price.value.String()
}

func (price *Price) Add(priceSecond Price) {
	price.value = price.value.Add(priceSecond.value)
}

func (price *Price) Sub(priceSecond Price) {
	price.value = price.value.Sub(priceSecond.value)
}

func (price *Price) Div(priceSecond Price) {
	price.value = price.value.Div(priceSecond.value)
}

func (price *Price) Mul(priceSecond Price) {
	price.value = price.value.Mul(priceSecond.value)
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
