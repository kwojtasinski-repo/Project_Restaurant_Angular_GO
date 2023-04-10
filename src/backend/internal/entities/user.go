package entities

import valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"

type User struct {
	Id       valueobjects.Id
	Email    valueobjects.EmailAddress
	Password string
	Role     string
}
