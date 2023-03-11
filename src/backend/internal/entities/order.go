package entities

import "github.com/shopspring/decimal"

type Order struct {
	Id            int64
	Number        string
	Price         decimal.Decimal
	User          User
	UserId        int64
	OrderProducts []OrderProduct
}
