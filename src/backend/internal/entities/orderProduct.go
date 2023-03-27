package entities

import "github.com/shopspring/decimal"

type OrderProduct struct {
	Id        int64
	Name      string
	Price     decimal.Decimal
	ProductId int64
	Product   Product
}
