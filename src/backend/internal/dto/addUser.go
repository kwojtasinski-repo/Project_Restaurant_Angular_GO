package dto

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type AddUserDto struct {
	Email    string
	Password string
}

func (user *AddUserDto) Validate() error {
	var validationErrors strings.Builder

	if !strings.Contains(user.Email, "@") {
		validationErrors.WriteString(fmt.Sprintf("'Email' '%v' is invalid", user.Email))
	}

	if len(user.Password) < 12 {
		validationErrors.WriteString("'Password' should contain at least 12 characters")
	}

	if len(user.Password) > 64 {
		validationErrors.WriteString("'Password' cannot contain more than 64 characters")
	}

	if match, _ := regexp.MatchString("^(.{12,64}|[^a-z]{1,}|[^A-Z]{1,}|[^\\d]{1,}|[^\\W]{1,})$|[\\s]", user.Password); !match {
		validationErrors.WriteString("'Password' should contain upper case letters, lower case letters and numbers")
	}

	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	}

	return nil
}
