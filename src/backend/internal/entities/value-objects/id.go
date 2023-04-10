package valueobjects

import (
	"errors"
	"fmt"
)

type Id struct {
	value int64
}

func NewId(id int64) (*Id, error) {
	if err := validateId(id); err != nil {
		return nil, err
	}

	return &Id{
		value: id,
	}, nil
}

func (id *Id) Value() int64 {
	return id.value
}

func (id *Id) String() string {
	return fmt.Sprintf("%v", id.value)
}

func validateId(id int64) error {
	if id < 0 {
		return errors.New("'Id' cannot be negative")
	}

	return nil
}
