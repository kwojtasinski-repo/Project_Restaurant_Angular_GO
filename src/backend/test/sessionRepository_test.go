package test

import (
	"time"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func (suite *IntegrationTestSuite) Test_AddSession_ShouldReturnSession() {
	user, _ := entities.NewUser(1, "email@email.com", "1234", "user")
	expiry := time.Now().UTC().Add(time.Hour * 2)
	session := entities.CreateSession(*user, expiry)
	sessionRepository := repositories.CreateSessionRepository(suite.database)

	sessionAdded, err := sessionRepository.AddSession(session)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), sessionAdded)
	assert.Equal(suite.T(), session.Expiry(), sessionAdded.Expiry())
	assert.Equal(suite.T(), session.Role(), sessionAdded.Role())
	assert.Equal(suite.T(), session.UserId(), sessionAdded.UserId())
}

func (suite *IntegrationTestSuite) Test_AddSession_ShouldGetSessionFromDatabase() {
	user, _ := entities.NewUser(1, "email@email.com", "1234", "user")
	expiry := time.Now().UTC().Add(time.Hour * 2)
	session := entities.CreateSession(*user, expiry)
	sessionRepository := repositories.CreateSessionRepository(suite.database)

	sessionAdded, err := sessionRepository.AddSession(session)

	sessionFromDatabase, errRepo := sessionRepository.GetSession(sessionAdded.SessionId())
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), errRepo)
	assert.NotNil(suite.T(), sessionAdded)
	assert.NotNil(suite.T(), sessionFromDatabase)
	assert.Equal(suite.T(), session.Expiry(), sessionAdded.Expiry())
	assert.Equal(suite.T(), session.Role(), sessionAdded.Role())
	assert.Equal(suite.T(), session.UserId(), sessionAdded.UserId())
	assert.Equal(suite.T(), session.Expiry().Format(time.RFC3339), sessionFromDatabase.Expiry().Format(time.RFC3339))
	assert.Equal(suite.T(), session.Role(), sessionFromDatabase.Role())
	assert.Equal(suite.T(), session.UserId(), sessionFromDatabase.UserId())
}

func (suite *IntegrationTestSuite) Test_AddMultipleUserSessions_DeleteAllUserSessions() {
	user, _ := entities.NewUser(1, "email@email.com", "1234", "user")
	expiry := time.Now().UTC().Add(time.Hour * 2)
	session := entities.CreateSession(*user, expiry)
	sessionRepository := repositories.CreateSessionRepository(suite.database)
	sessionRepository.AddSession(session)
	sessionRepository.AddSession(session)
	sessionRepository.AddSession(session)
	sessionRepository.AddSession(session)
	sessionRepository.AddSession(session)

	err := sessionRepository.DeleteAllUsersSessions(user.Id)

	sessions, errRepo := sessionRepository.GetSessionsByUserId(user.Id)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), errRepo)
	assert.NotNil(suite.T(), sessions)
	assert.Equal(suite.T(), 0, len(sessions))
}

func (suite *IntegrationTestSuite) Test_AddMultipleUsersSessions_CleanAllSessions() {
	user, _ := entities.NewUser(1, "email@email.com", "1234", "user")
	user2, _ := entities.NewUser(2, "email@email.com", "1234", "user")
	expiry := time.Now().UTC().Add(time.Hour * 2 * -1)
	session := entities.CreateSession(*user, expiry)
	session2 := entities.CreateSession(*user2, expiry)
	sessionRepository := repositories.CreateSessionRepository(suite.database)
	sessionRepository.AddSession(session)
	sessionRepository.AddSession(session)
	sessionRepository.AddSession(session)
	sessionRepository.AddSession(session)
	sessionRepository.AddSession(session)
	sessionRepository.AddSession(session2)
	sessionRepository.AddSession(session2)
	sessionRepository.AddSession(session2)
	sessionRepository.AddSession(session2)
	sessionRepository.AddSession(session2)
	sessionRepository.AddSession(entities.CreateSession(*user, time.Now().UTC().Add(time.Hour*2)))
	sessionRepository.AddSession(entities.CreateSession(*user2, time.Now().UTC().Add(time.Hour*2)))

	sessionRepository.DeleteSessionsExpiredAfter(time.Hour)

	sessionsUser, errUser := sessionRepository.GetSessionsByUserId(user.Id)
	sessionsUser2, errUser2 := sessionRepository.GetSessionsByUserId(user2.Id)
	assert.Nil(suite.T(), errUser)
	assert.NotNil(suite.T(), sessionsUser)
	assert.Nil(suite.T(), errUser2)
	assert.NotNil(suite.T(), sessionsUser2)
	assert.Equal(suite.T(), 1, len(sessionsUser))
	assert.Equal(suite.T(), 1, len(sessionsUser2))
}
