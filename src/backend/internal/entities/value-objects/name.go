package valueobjects

import (
	"errors"
	"strings"
)

type Name struct {
	value string
}

func NewName(name string) (*Name, error) {
	trimmedName := strings.TrimSpace(name)
	if err := validateName(trimmedName); err != nil {
		return nil, err
	}
	return &Name{
		value: trimmedName,
	}, nil
}

func (name *Name) Value() string {
	return name.value
}

func validateName(name string) error {
	var validationErrors strings.Builder
	if len(name) < 3 {
		validationErrors.WriteString("'Name' should have at least 3 characters. ")
	}
	if len(name) > 200 {
		validationErrors.WriteString("'Name' cannot have more than 200 characters. ")
	}

	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	} else {
		return nil
	}
}
