package entities

import (
	"strings"

	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type Category struct {
	Id      valueobjects.Id
	Name    valueobjects.Name
	Deleted bool // soft delete
}

func NewCategory(id int64, name string) (*Category, error) {
	var validationErrors strings.Builder
	var err error
	var newId *valueobjects.Id
	var newName *valueobjects.Name

	newId, err = valueobjects.NewId(id)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}
	newName, err = valueobjects.NewName(name)
	if err != nil {
		validationErrors.WriteString(err.Error())
	}

	return &Category{
		Id:      *newId,
		Name:    *newName,
		Deleted: false,
	}, nil
}
