package valueobjects

import (
	"errors"
	"strings"
)

type Description struct {
	value string
}

func NewDescription(description string) (*Description, error) {
	trimmedDescription := strings.TrimSpace(description)
	if err := validateDescription(trimmedDescription); err != nil {
		return nil, err
	}

	return &Description{
		value: strings.TrimSpace(trimmedDescription),
	}, nil
}

func (description *Description) Value() string {
	return description.value
}

func (description *Description) String() string {
	return description.value
}

func validateDescription(description string) error {
	var validationErrors strings.Builder
	if len(description) > 5000 {
		validationErrors.WriteString("'Description' cannot have more than 5000 characters. ")
	}
	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	} else {
		return nil
	}
}
