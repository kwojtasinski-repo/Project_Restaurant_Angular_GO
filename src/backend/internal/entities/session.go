package entities

import (
	"time"

	"github.com/google/uuid"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type Session struct {
	sessionId uuid.UUID
	userId    valueobjects.Id
	email     valueobjects.EmailAddress
	role      string
	expiry    time.Time
}

func CreateSession(user User, expiry time.Time) Session {
	return Session{
		sessionId: uuid.New(),
		userId:    user.Id,
		email:     user.Email,
		role:      user.Role,
		expiry:    expiry,
	}
}

func (s *Session) SessionId() uuid.UUID {
	return s.sessionId
}

func (s *Session) UserId() valueobjects.Id {
	return s.userId
}

func (s *Session) Email() valueobjects.EmailAddress {
	return s.email
}

func (s *Session) Role() string {
	return s.role
}

func (s *Session) Expiry() time.Time {
	return s.expiry
}

func (s *Session) IsExpired() bool {
	return s.expiry.Before(time.Now().In(time.UTC))
}

func (s *Session) SetSessionId(sessionId uuid.UUID) {
	s.sessionId = sessionId
}

func (s *Session) SetUser(user User) {
	s.userId = user.Id
	s.email = user.Email
	s.role = user.Role
}

func (s *Session) SetExpiry(expiry time.Time) {
	s.expiry = expiry
}
