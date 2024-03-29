package dto

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
)

type SessionDto struct {
	SessionId uuid.UUID `json:"sessionId"`
	Expiry    int64     `json:"expiry"`
	UserId    IdObject  `json:"userId"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}

func MapToSessionDto(session entities.Session) SessionDto {
	userId := session.UserId()
	email := session.Email()
	userIdObj, err := NewIntIdObject(userId.Value())
	if err != nil {
		panic(err)
	}
	return SessionDto{
		SessionId: session.SessionId(),
		UserId:    *userIdObj,
		Email:     email.Value(),
		Expiry:    session.Expiry().UnixMilli(),
		Role:      session.Role(),
	}
}

func (session *SessionDto) AsMap() map[string]interface{} {
	sessionMap := make(map[string]interface{})
	sessionMap["sessionId"] = session.SessionId
	sessionMap["email"] = session.Email
	sessionMap["userId"] = session.UserId
	sessionMap["expiry"] = session.Expiry
	sessionMap["role"] = session.Role
	return sessionMap
}

func (session *SessionDto) Validate() error {
	var validationErrors strings.Builder
	if len(session.Email) == 0 {
		validationErrors.WriteString("Empty Email. ")
	}

	if !strings.Contains(session.Email, "@") {
		validationErrors.WriteString("Invalid Email. ")
	}

	if len(session.Role) == 0 {
		validationErrors.WriteString("Invalid Role. ")
	}

	if session.Expiry == 0 {
		validationErrors.WriteString("Invalid Expiry time. ")
	}

	if validationErrors.Len() > 0 {
		return errors.New(validationErrors.String())
	}
	return nil
}
