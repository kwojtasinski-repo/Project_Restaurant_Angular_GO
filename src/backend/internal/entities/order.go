package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	Id            int64
	OrderNumber   string
	Price         decimal.Decimal
	Created       time.Time
	Modified      time.Time
	User          User
	UserId        int64
	OrderProducts []OrderProduct
}
