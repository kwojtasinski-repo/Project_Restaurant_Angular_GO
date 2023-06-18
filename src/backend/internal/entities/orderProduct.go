package entities

import (
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type OrderProduct struct {
	Id        valueobjects.Id
	Name      valueobjects.Name
	Price     valueobjects.Price
	ProductId valueobjects.Id
	Product   Product
}
