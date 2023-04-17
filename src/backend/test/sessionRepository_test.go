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
