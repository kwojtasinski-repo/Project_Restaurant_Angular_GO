package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type SessionRepository interface {
	AddSession(entities.Session) (entities.Session, error)
	DeleteSession(session entities.Session) error
	DeleteAllUsersSessions(userId valueobjects.Id) error
	UpdateSession(session entities.Session) error
	GetSession(sessionId uuid.UUID) (*entities.Session, error)
	GetSessionsByUserId(userId valueobjects.Id) ([]entities.Session, error)
	DeleteSessionsExpiredAfter(timeDuration time.Duration) error
}

type inMemorySessionRepository struct {
	sessions []entities.Session
}

func NewInMemorySessionRepository() SessionRepository {
	return &inMemorySessionRepository{
		sessions: make([]entities.Session, 0),
	}
}

func (repo *inMemorySessionRepository) AddSession(session entities.Session) (entities.Session, error) {
	var length int = len(repo.sessions)
	if length == 0 {
		repo.sessions = append(repo.sessions, session)
		return session, nil
	}

	repo.sessions = append(repo.sessions, session)
	return session, nil
}

func (repo *inMemorySessionRepository) DeleteSession(session entities.Session) error {
	for index, sessionRepo := range repo.sessions {
		if sessionRepo.SessionId() == session.SessionId() {
			repo.sessions = append(repo.sessions[:index], repo.sessions[index+1:]...)
			return nil
		}
	}
	return nil
}

func (repo *inMemorySessionRepository) UpdateSession(session entities.Session) error {
	for index, sessionRepo := range repo.sessions {
		if sessionRepo.SessionId() == session.SessionId() {
			repo.sessions[index] = session
			return nil
		}
	}

	return nil
}

func (repo *inMemorySessionRepository) GetSession(sessionId uuid.UUID) (*entities.Session, error) {
	for _, session := range repo.sessions {
		if session.SessionId() == sessionId {
			return &session, nil
		}
	}

	return nil, nil
}

func (repo *inMemorySessionRepository) GetSessionsByUserId(userId valueobjects.Id) ([]entities.Session, error) {
	sessions := make([]entities.Session, 0)
	for _, session := range repo.sessions {
		repoUserId := session.UserId()
		if repoUserId.Value() == userId.Value() {
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}

func (repo *inMemorySessionRepository) DeleteAllUsersSessions(userId valueobjects.Id) error {
	sessionsToDelete := make([]entities.Session, len(repo.sessions))
	copy(sessionsToDelete, repo.sessions)
	for _, session := range sessionsToDelete {
		sessionUserId := session.UserId()
		if sessionUserId.Value() == userId.Value() {
			repo.DeleteSession(session)
		}
	}
	return nil
}

func (repo *inMemorySessionRepository) DeleteSessionsExpiredAfter(timeDuration time.Duration) error {
	currentTime := time.Now().UTC()
	sessionsToDelete := make([]entities.Session, len(repo.sessions))
	copy(sessionsToDelete, repo.sessions)
	for _, session := range sessionsToDelete {
		if session.Expiry().Before(currentTime.Add(timeDuration * -1)) {
			repo.DeleteSession(session)
		}
	}
	return nil
}
