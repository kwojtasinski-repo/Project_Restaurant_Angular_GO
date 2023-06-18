package entities

import (
	"errors"
	"strings"

	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type User struct {
	Id       valueobjects.Id
	Email    valueobjects.EmailAddress
	Password string
	Role     string
	Deleted  bool // soft delete
}

func NewUser(id int64, email, password, role string) (*User, error) {
	var validationErrors strings.Builder
	var newId *valueobjects.Id
	var newEmail *valueobjects.EmailAddress
	var err error

	if newId, err = valueobjects.NewId(id); err != nil {
		validationErrors.WriteString(err.Error())
	}

	if newEmail, err = valueobjects.NewEmailAddress(email); err != nil {
		validationErrors.WriteString(err.Error())
	}

	if validationErrors.Len() > 0 {
		return nil, errors.New(validationErrors.String())
	}

	user := &User{
		Id:       *newId,
		Email:    *newEmail,
		Password: password,
		Role:     role,
	}

	return user, nil
}
