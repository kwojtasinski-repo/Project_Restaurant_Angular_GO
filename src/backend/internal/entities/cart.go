package entities

import (
	"strings"

	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type Cart struct {
	Id        valueobjects.Id
	UserId    valueobjects.Id
	User      User
	ProductId valueobjects.Id
	Product   Product
}

func NewCart(id int64, user User, product Product) (*Cart, error) {
	var validationErrors strings.Builder
	var err error
	var newId *valueobjects.Id

	newId, err = valueobjects.NewId(id)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}

	return &Cart{
		Id:        *newId,
		User:      user,
		UserId:    user.Id,
		Product:   product,
		ProductId: product.Id,
	}, nil
}
