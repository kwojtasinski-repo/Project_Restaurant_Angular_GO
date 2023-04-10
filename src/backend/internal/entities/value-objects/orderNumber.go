package valueobjects

import (
	"strings"

	"github.com/google/uuid"
)

type OrderNumber struct {
	value string
}

func NewOrderNumberWithGiven(orderNumber string) OrderNumber {
	trimmedOrderNumber := strings.TrimSpace(orderNumber)
	return OrderNumber{
		value: trimmedOrderNumber,
	}
}

func NewOrderNumber() OrderNumber {
	return OrderNumber{
		value: uuid.New().String(),
	}
}

func (orderNumber *OrderNumber) Value() string {
	return orderNumber.value
}

func (orderNumber *OrderNumber) String() string {
	return orderNumber.value
}
