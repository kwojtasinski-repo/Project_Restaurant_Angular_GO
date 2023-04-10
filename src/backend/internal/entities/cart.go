package entities

import valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"

type Cart struct {
	Id        valueobjects.Id
	UserId    valueobjects.Id
	User      User
	ProductId valueobjects.Id
	Product   Product
}
