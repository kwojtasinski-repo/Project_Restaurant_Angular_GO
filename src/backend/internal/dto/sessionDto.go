package dto

import (
	"github.com/google/uuid"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

type SessionDto struct {
	SessionId uuid.UUID
	Expiry    int64
	UserId    int64
	Email     string
	Role      string
}

func MapToSessionDto(session entities.Session) SessionDto {
	userId := session.UserId()
	email := session.Email()
	return SessionDto{
		SessionId: session.SessionId(),
		UserId:    userId.Value(),
		Email:     email.Value(),
		Expiry:    session.Expiry().UnixMilli(),
		Role:      session.Role(),
	}
}
