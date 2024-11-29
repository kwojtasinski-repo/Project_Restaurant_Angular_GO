package dto

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"
)

type AddUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var passwordPattern = `^(?=.*\p{Ll})(?=.*\p{Lu})(?=.*\d)(?=.*[\p{P}\p{S}]).+$`
var passwordPatternCompiled = regexp2.MustCompile(passwordPattern, regexp2.RE2)

func (user *AddUserDto) Validate() error {
	var validationErrors strings.Builder

	if !strings.Contains(user.Email, "@") {
		validationErrors.WriteString(fmt.Sprintf("'Email' '%v' is invalid", user.Email))
	}

	if len(user.Password) < 12 {
		validationErrors.WriteString("'Password' should contain at least 12 characters. ")
	}

	if len(user.Password) > 64 {
		validationErrors.WriteString("'Password' cannot contain more than 64 characters. ")
	}

	if match, _ := passwordPatternCompiled.MatchString(user.Password); !match {
		validationErrors.WriteString("'Password' should contain upper case letters, lower case letters, number and special character. ")
	}

	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	}

	return nil
}
