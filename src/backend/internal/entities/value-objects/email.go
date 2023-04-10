package valueobjects

import (
	"errors"
	"fmt"
	"strings"
)

type EmailAddress struct {
	value string
}

func (emailAddress *EmailAddress) Value() string {
	return emailAddress.value
}

func (emailAddress *EmailAddress) String() string {
	return emailAddress.value
}

func NewEmailAddress(email string) (*EmailAddress, error) {
	emailTrimmed := strings.TrimSpace(email)
	if err := validateEmail(emailTrimmed); err != nil {
		return nil, err
	}

	return &EmailAddress{
		value: email,
	}, nil
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New(getInvalidEmail(email))
	}

	splitedEmail := strings.Split(email, "@")
	emailName := splitedEmail[0]
	emailDomain := splitedEmail[1]

	if len(emailName) == 0 {
		return errors.New(getInvalidEmail(email))
	}

	if len(emailDomain) == 0 {
		return errors.New(getInvalidEmail(email))
	}

	if !strings.Contains(emailDomain, ".") {
		return errors.New(getInvalidEmail(email))
	}

	return nil
}

func getInvalidEmail(email string) string {
	return fmt.Sprintf("'Email' '%v' is invalid", email)
}
