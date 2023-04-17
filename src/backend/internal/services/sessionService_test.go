package services

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/settings"
	"github.com/stretchr/testify/suite"
)

type SessionServiceTestSuite struct {
	suite.Suite
	service        SessionService
	passwordHasher PasswordHasherService
	testUsers      []entities.User
	testSessions   []entities.Session
}

func (suite *SessionServiceTestSuite) SetupTest() {
	log.Println("---- Setup SessionServiceTestSuite Before Each Test ----")
	suite.testUsers = make([]entities.User, 0)
	suite.testSessions = make([]entities.Session, 0)
	suite.passwordHasher = CreatePassworHasherService()
	userRepository := suite.createTestInMemoryUserRepository()
	sessinoRepository := suite.createTestInMemorySessionRepository()
	suite.service = CreateSessionService(sessinoRepository, userRepository)
}

func (suite *SessionServiceTestSuite) createTestInMemorySessionRepository() repositories.SessionRepository {
	repo := repositories.NewInMemorySessionRepository()
	for _, user := range suite.testUsers {
		session1, _ := repo.AddSession(suite.addTestSession(&user))
		session2, _ := repo.AddSession(suite.addTestSession(&user))
		session3, _ := repo.AddSession(entities.CreateSession(user, time.Now().Add(time.Duration(time.Hour)*3*-1)))
		session4, _ := repo.AddSession(entities.CreateSession(user, time.Now().Add(time.Duration(settings.CookieLifeTime)*300*-1)))
		session5, _ := repo.AddSession(entities.CreateSession(user, time.Now().Add(time.Duration(settings.CookieLifeTime)*300*-1)))
		suite.testSessions = append(suite.testSessions, session1)
		suite.testSessions = append(suite.testSessions, session2)
		suite.testSessions = append(suite.testSessions, session3)
		suite.testSessions = append(suite.testSessions, session4)
		suite.testSessions = append(suite.testSessions, session5)
	}
	return repo
}

func (suite *SessionServiceTestSuite) createTestInMemoryUserRepository() repositories.UserRepository {
	repo := repositories.NewInMemoryUserRepository()
	user1 := suite.addTestUser("admin@admin.com", "admin")
	user2 := suite.addTestUser("email123@admin.com", "user")
	user3 := suite.addTestUser("test123@admin.com", "user")
	repo.Add(user1)
	repo.Add(user2)
	repo.Add(user3)
	suite.testUsers = append(suite.testUsers, *user1)
	suite.testUsers = append(suite.testUsers, *user2)
	suite.testUsers = append(suite.testUsers, *user3)
	return repo
}

func (suite *SessionServiceTestSuite) addTestUser(email, role string) *entities.User {
	hashedPassword, _ := suite.passwordHasher.HashPassword("Password1234!")
	user, _ := entities.NewUser(int64(rand.Intn(1000000-1)+1), email, hashedPassword, role)
	return user
}

func (suite *SessionServiceTestSuite) addTestSession(user *entities.User) entities.Session {
	expiry := time.Now().UTC()
	var session entities.Session
	if user != nil {
		session = entities.CreateSession(*user, expiry)
	} else {
		session = entities.CreateSession(*suite.addTestUser("email@email"+string(rune((rand.Intn(1000000-1)+1)))+".com", "user"), expiry)
	}
	return session
}

func (suite *SessionServiceTestSuite) findAdminUser() *entities.User {
	for _, user := range suite.testUsers {
		if user.Role == "admin" {
			return &user
		}
	}
	return nil
}

func (suite *SessionServiceTestSuite) containsSession(sessions []dto.SessionDto, session dto.SessionDto) bool {
	for _, sessionInCollection := range sessions {
		if sessionInCollection.SessionId == session.SessionId {
			return true
		}
	}
	return false
}

func (suite *SessionServiceTestSuite) getPermanentlyExpiredSessions() []entities.Session {
	sessions := make([]entities.Session, 0)
	for _, session := range suite.testSessions {
		if session.Expiry().Before(time.Now().UTC().Add(time.Duration(settings.CookieLifeTime * -1))) {
			sessions = append(sessions, session)
		}
	}

	return sessions
}

func (suite *SessionServiceTestSuite) getExpiredSessions() []entities.Session {
	sessions := make([]entities.Session, 0)
	for _, session := range suite.testSessions {
		if session.Expiry().Before(time.Now().UTC().Add(time.Duration(time.Hour*2*-1))) && session.Expiry().After(time.Now().UTC().Add(time.Duration(settings.CookieLifeTime*-1))) {
			sessions = append(sessions, session)
		}
	}

	return sessions
}

func TestSessionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SessionServiceTestSuite))
}

func (suite *SessionServiceTestSuite) Test_CreateSession_ValidUser_ShouldCreateAndReturnDto() {
	user := suite.testUsers[0]

	session, err := suite.service.CreateSession(user)

	suite.Assertions.NotNil(session)
	suite.Assertions.Nil(err)
}

func (suite *SessionServiceTestSuite) Test_GetUserSessions_WithCreatedSessions_ShouldReturnDtos() {
	user := suite.findAdminUser()
	suite.service.CreateSession(*user)
	suite.service.CreateSession(*user)
	sessionDto, _ := suite.service.CreateSession(*user)

	sessions, err := suite.service.GetUserSessions(user.Id.Value())

	suite.Assertions.NotNil(sessions)
	suite.Assertions.Nil(err)
	suite.Assertions.NotEmpty(sessions)
	suite.Assertions.Equal(8, len(sessions))
	suite.Assertions.Equal(suite.containsSession(sessions, *sessionDto), true)
}

func (suite *SessionServiceTestSuite) Test_RevokeSession_ValidSessionId_ShouldRevokeSession() {
	session := suite.testSessions[0]
	userId := session.UserId()
	user, _ := entities.NewUser(userId.Value(), "email@asfasf.cam", "hashedPassword", "role")
	sessionCreated, _ := suite.service.CreateSession(*user)

	err := suite.service.RevokeSession(session.SessionId())

	suite.Assertions.Nil(err)
	sessions, err := suite.service.GetUserSessions(userId.Value())
	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(sessions)
	suite.Assertions.NotEmpty(sessions)
	suite.Assertions.Equal(suite.containsSession(sessions, *sessionCreated), true)
	suite.Assertions.Equal(suite.containsSession(sessions, dto.MapToSessionDto(session)), false)
}

func (suite *SessionServiceTestSuite) Test_RevokeAllUserSessions_ValidSessionId_ShouldRevokeSessions() {
	session := suite.testSessions[0]
	userId := session.UserId()
	user, _ := entities.NewUser(userId.Value(), "email@asfasf.cam", "hashedPassword", "role")
	sessionCreated, _ := suite.service.CreateSession(*user)
	sessionsBefore, _ := suite.service.GetUserSessions(userId.Value())

	err := suite.service.RevokeAllUsersSessions(userId.Value())

	suite.Assertions.Nil(err)
	sessions, err := suite.service.GetUserSessions(userId.Value())
	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(sessions)
	suite.Assertions.NotEmpty(sessionsBefore)
	suite.Assertions.Empty(sessions)
	suite.Assertions.Equal(suite.containsSession(sessions, *sessionCreated), false)
	suite.Assertions.NotEqual(len(sessions) == len(sessionsBefore), true)
}

func (suite *SessionServiceTestSuite) Test_ClearExpiredSessions_ShouldRevokeExpiredSessions() {
	sessionsExpiredBefore := suite.getPermanentlyExpiredSessions()
	suite.Assertions.NotNil(sessionsExpiredBefore)
	suite.Assertions.NotEmpty(sessionsExpiredBefore)
	userId := sessionsExpiredBefore[0].UserId()
	userSessionsBefore, err := suite.service.GetUserSessions(userId.Value())
	suite.Assertions.Nil(err)
	suite.Assertions.NotEmpty(userSessionsBefore)

	suite.service.ClearExpiredSessions()

	userSessionsAfter, err := suite.service.GetUserSessions(userId.Value())
	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(userSessionsAfter)
	suite.Assertions.Equal(len(userSessionsBefore) == len(userSessionsAfter), false)
}

func (suite *SessionServiceTestSuite) Test_RefreshSession_ValidSessionId_ShouldRevokeExpiredSessions() {
	session := suite.testSessions[0]

	sessionRefreshed, err := suite.service.RefreshSession(session.SessionId())

	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(sessionRefreshed)
	suite.Assertions.Equal(sessionRefreshed.Expiry > session.Expiry().UnixMilli(), true)
}

func (suite *SessionServiceTestSuite) Test_ManageSession_ExpiredSession_ShouldRefreshSession() {
	sessions := suite.getExpiredSessions()
	suite.Assertions.NotEmpty(sessions)
	sessionDto := dto.MapToSessionDto(sessions[0])

	suite.service.ManageSession(sessionDto)

	userSessions, _ := suite.service.GetUserSessions(sessionDto.UserId)
	var sessionRefreshed dto.SessionDto
	for _, session := range userSessions {
		if session.SessionId == sessionDto.SessionId {
			sessionRefreshed = session
			break
		}
	}
	suite.Assertions.Equal(sessionRefreshed.Expiry > sessionDto.Expiry, true)
}
