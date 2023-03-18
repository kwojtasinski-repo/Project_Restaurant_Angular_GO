package entities

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	Id          int64
	Name        string
	Description string
	Price       decimal.Decimal
	Category    Category
	Deleted     bool // soft delete
}
