package repositories

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type sessionRepository struct {
	database sql.DB
}

func CreateSessionRepository(database sql.DB) SessionRepository {
	return &sessionRepository{
		database: database,
	}
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
	query := "DELETE FROM sessions WHERE session_id = ?"
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
