package entities

import valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"

type Category struct {
	Id      valueobjects.Id
	Name    valueobjects.Name
	Deleted bool // soft delete
}
