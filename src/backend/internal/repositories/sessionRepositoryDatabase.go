package repositories

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/settings"
	"github.com/patrickmn/go-cache"
)

type sessionRepository struct {
	database *sql.DB
}

var sessionRepositoryCached = &cachedSessionRepository{
	cacheStore: cache.New(settings.TimeStoreInCache, settings.TimeStoreInCache),
}

func CreateSessionRepository(database *sql.DB) SessionRepository {
	sessionRepository := &sessionRepository{
		database: database,
	}
	sessionRepositoryCached.innerRepo = sessionRepository
	return sessionRepositoryCached
}

func (repo *sessionRepository) AddSession(session entities.Session) (entities.Session, error) {
	queryId := "SELECT UUID()"
	result := repo.database.QueryRow(queryId)
	var sessionId uuid.UUID
	if err := result.Scan(&sessionId); err != nil {
		return entities.Session{}, err
	}

	query := `INSERT INTO sessions (session_id, user_id, email, role, expiry) VALUES(UNHEX(REPLACE(?,'-','')), ?, ?, ?, ?);`
	userId := session.UserId()
	email := session.Email()
	_, err := repo.database.Exec(query, sessionId, userId.Value(), email.Value(), session.Role(), session.Expiry())
	if err != nil {
		return entities.Session{}, err
	}

	session.SetSessionId(sessionId)
	return session, nil
}

func (repo *sessionRepository) DeleteSession(session entities.Session) error {
	query := "DELETE FROM sessions WHERE session_id = UNHEX(REPLACE(?,'-',''))"
	if _, err := repo.database.Exec(query, session.SessionId()); err != nil {
		return nil
	}

	return nil
}

func (repo *sessionRepository) UpdateSession(session entities.Session) error {
	query := "UPDATE sessions SET role = ?, expiry = ?;"
	if _, err := repo.database.Exec(query, session.Role(), session.Expiry()); err != nil {
		return err
	}

	return nil
}

func (repo *sessionRepository) GetSession(sessionId uuid.UUID) (*entities.Session, error) {
	query := `SELECT LOWER(CONCAT(
		SUBSTR(HEX(session_id), 1, 8), '-',
		SUBSTR(HEX(session_id), 9, 4), '-',
		SUBSTR(HEX(session_id), 13, 4), '-',
		SUBSTR(HEX(session_id), 17, 4), '-',
		SUBSTR(HEX(session_id), 21)
		)), user_id, email, role, expiry FROM sessions WHERE session_id = UNHEX(REPLACE(?,'-',''))`
	row := repo.database.QueryRow(query, sessionId)

	var newSessionId uuid.UUID
	var expiry time.Time
	var userId int64
	var email string
	var role string
	if err := row.Scan(&newSessionId, &userId, &email, &role, &expiry); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	user, _ := entities.NewUser(userId, email, "", role)
	session := &entities.Session{}
	session.SetSessionId(newSessionId)
	session.SetUser(*user)
	session.SetExpiry(expiry)
	return session, nil
}

func (repo *sessionRepository) GetSessionsByUserId(userId valueobjects.Id) ([]entities.Session, error) {
	sessions := make([]entities.Session, 0)
	query := "SELECT session_id, user_id, email, role, expiry FROM sessions WHERE user_id = ?"
	rows, err := repo.database.Query(query, userId.Value())
	if err != nil {
		return sessions, err
	}
	defer rows.Close()

	for rows.Next() {
		var newSessionId uuid.UUID
		var expiry time.Time
		var userId int64
		var email string
		var role string
		if err := rows.Scan(&newSessionId, &userId, &email, &role, &expiry); err != nil {
			return nil, err
		}

		user, _ := entities.NewUser(userId, email, "", role)
		session := &entities.Session{}
		session.SetSessionId(newSessionId)
		session.SetUser(*user)
		session.SetExpiry(expiry)
		sessions = append(sessions, *session)
	}

	return sessions, nil
}

func (repo *sessionRepository) DeleteAllUsersSessions(userId valueobjects.Id) error {
	query := `DELETE FROM sessions WHERE user_id = ?;`
	_, err := repo.database.Exec(query, userId.Value())
	if err != nil {
		return err
	}

	return nil
}

func (repo *sessionRepository) DeleteSessionsExpiredAfter(timeDuration time.Duration) error {
	query := `DELETE FROM sessions WHERE expiry < ?;`
	expiryTime := time.Now().UTC().Add(timeDuration * -1)
	_, err := repo.database.Exec(query, expiryTime)
	if err != nil {
		return err
	}

	return nil
}

type cachedSessionRepository struct {
	cacheStore *cache.Cache
	innerRepo  SessionRepository
}

func (repo *cachedSessionRepository) AddSession(session entities.Session) (entities.Session, error) {
	sessionAdded, err := repo.innerRepo.AddSession(session)
	if err != nil {
		return sessionAdded, err
	}

	repo.cacheStore.Set(sessionAdded.SessionId().String(), &sessionAdded, settings.TimeStoreInCache)
	return sessionAdded, nil
}

func (repo *cachedSessionRepository) DeleteSession(session entities.Session) error {
	err := repo.innerRepo.DeleteSession(session)
	if err != nil {
		return err
	}

	repo.cacheStore.Delete(session.SessionId().String())
	return nil
}

func (repo *cachedSessionRepository) UpdateSession(session entities.Session) error {
	err := repo.innerRepo.UpdateSession(session)
	if err != nil {
		return err
	}

	repo.cacheStore.Set(session.SessionId().String(), &session, settings.TimeStoreInCache)
	return nil
}

func (repo *cachedSessionRepository) GetSession(sessionId uuid.UUID) (*entities.Session, error) {
	sessionInCache, exists := repo.cacheStore.Get(sessionId.String())
	if exists {
		return sessionInCache.(*entities.Session), nil
	}

	session, err := repo.innerRepo.GetSession(sessionId)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, nil
	}

	repo.cacheStore.Set(sessionId.String(), session, settings.TimeStoreInCache)
	return session, nil
}

func (repo *cachedSessionRepository) GetSessionsByUserId(userId valueobjects.Id) ([]entities.Session, error) {
	sessions, err := repo.innerRepo.GetSessionsByUserId(userId)
	if err != nil {
		return sessions, err
	}

	return sessions, nil
}

func (repo *cachedSessionRepository) DeleteAllUsersSessions(userId valueobjects.Id) error {
	if err := repo.innerRepo.DeleteAllUsersSessions(userId); err != nil {
		return err
	}

	sessions := repo.cacheStore.Items()
	for key, value := range sessions {
		session := value.Object.(*entities.Session)
		userIdCached := session.UserId()
		if userIdCached.Value() == userId.Value() {
			repo.cacheStore.Delete(key)
		}
	}

	return nil
}

func (repo *cachedSessionRepository) DeleteSessionsExpiredAfter(timeDuration time.Duration) error {
	if err := repo.innerRepo.DeleteSessionsExpiredAfter(timeDuration); err != nil {
		return err
	}

	return nil
}
