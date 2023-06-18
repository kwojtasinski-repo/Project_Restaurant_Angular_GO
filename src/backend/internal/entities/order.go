package entities

import (
	"errors"
	"fmt"
	"strings"
	"time"

	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/shopspring/decimal"
)

type Order struct {
	Id            valueobjects.Id
	OrderNumber   valueobjects.OrderNumber
	Price         valueobjects.Price
	Created       time.Time
	Modified      *time.Time
	User          User
	UserId        valueobjects.Id
	OrderProducts []OrderProduct
}

func NewOrder(id int64, user User, orderProducts []OrderProduct) (*Order, error) {
	var validationErrors strings.Builder
	var err error
	var newId *valueobjects.Id
	var newUserId *valueobjects.Id

	if newId, err = valueobjects.NewId(id); err != nil {
		validationErrors.WriteString(err.Error())
	}
	if newUserId, err = valueobjects.NewId(user.Id.Value()); err != nil {
		validationErrors.WriteString(err.Error())
	}
	if validationErrors.Len() > 0 {
		return nil, errors.New(validationErrors.String())
	}

	price, _ := valueobjects.NewPrice(decimal.Zero)

	order := &Order{
		Id:            *newId,
		OrderNumber:   valueobjects.NewOrderNumber(),
		Price:         *price,
		User:          user,
		UserId:        *newUserId,
		Created:       time.Now(),
		OrderProducts: orderProducts,
	}

	for _, orderProduct := range orderProducts {
		err = order.AddProduct(orderProduct)
		if err != nil {
			validationErrors.WriteString(err.Error())
		}
	}

	if validationErrors.Len() > 0 {
		return nil, errors.New(validationErrors.String())
	}

	return order, nil
}

func NewOrderWithNumber(id int64, orderNumber string, user User, orderProducts []OrderProduct) (*Order, error) {
	order, err := NewOrder(id, user, orderProducts)

	if err != nil {
		return nil, err
	}

	order.OrderNumber = valueobjects.NewOrderNumberWithGiven(orderNumber)
	return order, nil
}

func (order *Order) AddProduct(orderProduct OrderProduct) error {
	order.OrderProducts = append(order.OrderProducts, orderProduct)
	order.Price.Add(orderProduct.Price)
	return nil
}

func (order *Order) RemoveProduct(orderProduct OrderProduct) error {
	var exists bool = false
	for index, orderProductFromCollection := range order.OrderProducts {
		if orderProductFromCollection.Id.Value() == orderProduct.Id.Value() {
			exists = true
			order.OrderProducts = append(order.OrderProducts[:index], order.OrderProducts[index+1:]...)
			order.Price.Sub(orderProduct.Price)
		}
	}

	if !exists {
		return fmt.Errorf("'Product' with id %v is not exists", orderProduct.Id.Value())
	}

	return nil
}
