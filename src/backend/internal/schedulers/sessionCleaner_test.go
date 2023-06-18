package schedulers

import (
	"database/sql"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/settings"
	"github.com/stretchr/testify/suite"
)

type SessionCleanerTestSuite struct {
	suite.Suite
}

func (suite *SessionCleanerTestSuite) SetupTest() {
	log.Println("---- Setup SessionCleanerTestSuite Before Each Test ----")
	testSessionRepositoryStub = nil
	sessionsBeforeClear = make([]entities.Session, 0)
}

func createDatabaseTest() (*sql.DB, error) {
	return nil, nil
}

func createDatabaseTestWithError() (*sql.DB, error) {
	return nil, errors.New("some error")
}

func createTestSessionRepository(database *sql.DB) repositories.SessionRepository {
	repository := repositories.NewInMemorySessionRepository()
	user1, _ := entities.NewUser(1, "email@email.com", "", "user")
	session1 := entities.CreateSession(*user1, time.Now().UTC().Add(time.Duration(settings.CookieLifeTime+2)*-1))
	session2 := entities.CreateSession(*user1, time.Now().UTC().Add(time.Duration(settings.CookieLifeTime+2)*-1))
	session3 := entities.CreateSession(*user1, time.Now().UTC().Add(time.Duration(settings.CookieLifeTime+2)*-1))
	sessionAdded1, _ := repository.AddSession(session1)
	sessionAdded2, _ := repository.AddSession(session2)
	sessionAdded3, _ := repository.AddSession(session3)
	testSessionRepositoryStub = repository
	sessionsBeforeClear = append(sessionsBeforeClear, sessionAdded1)
	sessionsBeforeClear = append(sessionsBeforeClear, sessionAdded2)
	sessionsBeforeClear = append(sessionsBeforeClear, sessionAdded3)
	return repository
}

var testSessionRepositoryStub repositories.SessionRepository
var sessionsBeforeClear []entities.Session = make([]entities.Session, 0)

func TestSessionCleanerTestSuite(t *testing.T) {
	suite.Run(t, new(SessionCleanerTestSuite))
}

func (suite *SessionCleanerTestSuite) Test_CleanPermanentlyExpiredSessions_ShouldDeletePermanentlyExpiredSessions() {
	createDatabase = createDatabaseTest
	createSessionRepository = createTestSessionRepository

	cleanPermanentlyExpiredSessions()

	suite.Assertions.NotNil(sessionsBeforeClear)
	suite.Assertions.NotNil(testSessionRepositoryStub)
	session := sessionsBeforeClear[0]
	sessionAferDelete, _ := testSessionRepositoryStub.GetSessionsByUserId(session.UserId())
	suite.Assertions.NotEqual(len(sessionAferDelete), len(sessionsBeforeClear))
}

func (suite *SessionCleanerTestSuite) Test_CleanPermanentlyExpiredSessions_WithErrorCreatingConnectionWithDatabase_ShouldntCreateSessionRepository() {
	createDatabase = createDatabaseTestWithError

	cleanPermanentlyExpiredSessions()

	suite.Assertions.Empty(sessionsBeforeClear)
	suite.Assertions.Nil(testSessionRepositoryStub)
}
