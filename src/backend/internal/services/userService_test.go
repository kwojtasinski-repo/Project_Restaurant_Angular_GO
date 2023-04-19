package services

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	userRepository repositories.UserRepository
	sessionService SessionService
	passwordHasher PasswordHasherService
	service        UserService
	commonPassword string
}

func (suite *UserServiceTestSuite) SetupTest() {
	log.Println("---- Setup UserServiceTestSuite Before Each Test ----")
	suite.commonPassword = "PasW0Rd9871!2hth"
	suite.passwordHasher = CreatePassworHasherService()
	suite.userRepository = suite.createUserRepository()
	suite.sessionService = CreateSessionService(repositories.NewInMemorySessionRepository(), suite.userRepository)
	suite.service = CreateUserService(suite.userRepository, suite.passwordHasher, suite.sessionService)
}

func (suite *UserServiceTestSuite) createUserRepository() repositories.UserRepository {
	repository := repositories.NewInMemoryUserRepository()
	password, _ := suite.passwordHasher.HashPassword(suite.commonPassword)
	user1, _ := entities.NewUser(1, "email@email1.com", password, "user")
	user2, _ := entities.NewUser(2, "email@email2.com", password, "admin")
	user3, _ := entities.NewUser(3, "email@email3.com", password, "user")
	repository.Add(user1)
	repository.Add(user2)
	repository.Add(user3)
	return repository
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) Test_Get_ValidUserId_ShouldReturnUser() {
	userId := 1

	userDto, err := suite.service.Get(int64(userId))

	suite.Assertions.NotNil(userDto)
	suite.Assertions.Nil(userDto.Deleted)
	suite.Assertions.Nil(err)
}

func (suite *UserServiceTestSuite) Test_Get_NotFoundUser_ShouldReturnNotFound() {
	userId := 1000

	userDto, err := suite.service.Get(int64(userId))

	suite.Assertions.Nil(userDto)
	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusNotFound)
}

func (suite *UserServiceTestSuite) Test_Get_AnErrorOccuredInUserRepository_ShouldReturnInternalServerError() {
	service := CreateUserService(repositories.NewErrorUserRepository(), suite.passwordHasher, suite.sessionService)
	userId := 1

	userDto, err := service.Get(int64(userId))

	suite.Assertions.Nil(userDto)
	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *UserServiceTestSuite) Test_Delete_ValidUserId_ShouldDeleteUser() {
	user, _ := suite.userRepository.Get(1)

	err := suite.service.Delete(user.Id.Value())

	suite.Assertions.Nil(err)
	user2, err := suite.service.Get(1)
	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(user2)
	suite.Assertions.NotNil(user2.Deleted)
	suite.Assertions.Equal(*user2.Deleted, true)
}
func (suite *UserServiceTestSuite) Test_Delete_ValidUserId_ShouldDeleteUserAndRevokeAllActiveSessions() {
	user, _ := suite.userRepository.Get(1)
	suite.sessionService.CreateSession(*user)
	suite.sessionService.CreateSession(*user)
	sessions, _ := suite.sessionService.GetUserSessions(user.Id.Value())

	err := suite.service.Delete(user.Id.Value())

	sessionsAfterDelete, _ := suite.sessionService.GetUserSessions(user.Id.Value())
	suite.Assertions.NotEmpty(sessions)
	suite.Assertions.Empty(sessionsAfterDelete)
	suite.Assertions.Nil(err)
	user2, err := suite.service.Get(1)
	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(user2)
	suite.Assertions.NotNil(user2.Deleted)
	suite.Assertions.Equal(*user2.Deleted, true)
}

func (suite *UserServiceTestSuite) Test_Delete_AnErrorOccuredInUserRepository_ShouldReturnInternalServerError() {
	service := CreateUserService(repositories.NewErrorUserRepository(), suite.passwordHasher, suite.sessionService)

	err := service.Delete(1)

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *UserServiceTestSuite) Test_Delete_UserNotFound_ShouldReturnNotFound() {
	err := suite.service.Delete(10000)

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusNotFound)
}

func (suite *UserServiceTestSuite) Test_GetAll_ShouldReturnAllUsers() {
	users, err := suite.service.GetAll()

	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(users)
	suite.Assertions.NotEmpty(users)
}

func (suite *UserServiceTestSuite) Test_GetAll_AnErrorOccuredInUserRepository_ShouldReturnInternalServerError() {
	service := CreateUserService(repositories.NewErrorUserRepository(), suite.passwordHasher, suite.sessionService)

	users, err := service.GetAll()

	suite.Assertions.NotNil(err)
	suite.Assertions.Nil(users)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *UserServiceTestSuite) Test_Login_ShouldLoginAndCreateSession() {
	signInDto := dto.SignInDto{
		Email:    "email@email1.com",
		Password: suite.commonPassword,
	}
	timeSessionExpireBefore := time.Now().UTC().Add((time.Hour * 2) - time.Second)

	sessionDto, err := suite.service.Login(signInDto)

	timeSessionExpireAfter := time.Now().UTC().Add((time.Hour * 2) + time.Second)
	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(sessionDto)
	suite.Assertions.Greater(timeSessionExpireAfter.UnixMilli(), sessionDto.Expiry)
	suite.Assertions.Greater(sessionDto.Expiry, timeSessionExpireBefore.UnixMilli())
}

func (suite *UserServiceTestSuite) Test_Login_EmailNotExists_ShouldReturnBadRequest() {
	signInDto := dto.SignInDto{
		Email:    "email@email1231.com",
		Password: suite.commonPassword,
	}

	sessionDto, err := suite.service.Login(signInDto)

	suite.Assertions.NotNil(err)
	suite.Assertions.Nil(sessionDto)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
	suite.Assertions.Contains(err.Message, "Invalid Credentials")
}

func (suite *UserServiceTestSuite) Test_Login_InvalidPassword_ShouldReturnBadRequest() {
	signInDto := dto.SignInDto{
		Email:    "email@email1.com",
		Password: "suite.commonPassword",
	}

	sessionDto, err := suite.service.Login(signInDto)

	suite.Assertions.NotNil(err)
	suite.Assertions.Nil(sessionDto)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
	suite.Assertions.Contains(err.Message, "Invalid Credentials")
}

func (suite *UserServiceTestSuite) Test_Login_AnErrorOccuredInUserRepository_ShouldReturnInternalServerError() {
	service := CreateUserService(repositories.NewErrorUserRepository(), suite.passwordHasher, suite.sessionService)
	signInDto := dto.SignInDto{
		Email:    "email@email3.com",
		Password: suite.commonPassword,
	}

	sessionDto, err := service.Login(signInDto)

	suite.Assertions.NotNil(err)
	suite.Assertions.Nil(sessionDto)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *UserServiceTestSuite) Test_Register_ShouldAddNewUser() {
	addUser := dto.AddUserDto{
		Email:    "newemail@email.com",
		Password: "Testowe1231232!",
	}

	userDto, err := suite.service.Register(&addUser)

	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(userDto)
	suite.Assertions.Equal(addUser.Email, userDto.Email)
	suite.Assertions.Nil(userDto.Deleted)
}

func (suite *UserServiceTestSuite) Test_Register_ShouldLoginAsNewUser() {
	addUser := dto.AddUserDto{
		Email:    "newemail@email.com",
		Password: "Testowe1231232!",
	}

	user, errRegister := suite.service.Register(&addUser)

	sessionDto, errService := suite.service.Login(dto.SignInDto(addUser))
	suite.Assertions.Nil(errRegister)
	suite.Assertions.Nil(errService)
	suite.Assertions.NotNil(user)
	suite.Assertions.NotNil(sessionDto)
	suite.Assertions.Equal(addUser.Email, sessionDto.Email)
	suite.Assertions.Equal(user.Id, sessionDto.UserId)
}

func (suite *UserServiceTestSuite) Test_Register_InvalidEmailAndPassword_ShouldReturnBadRequest() {
	addUser := dto.AddUserDto{
		Email:    "newemail",
		Password: "testowe",
	}

	user, errRegister := suite.service.Register(&addUser)

	suite.Assertions.NotNil(errRegister)
	suite.Assertions.Nil(user)
	suite.Assertions.Equal(http.StatusBadRequest, errRegister.Status)
}

func (suite *UserServiceTestSuite) Test_Register_ExistsEmail_ShouldReturnBadRequest() {
	addUser := dto.AddUserDto{
		Email:    "email@email2.com",
		Password: "Testowe123!aewe",
	}

	user, errRegister := suite.service.Register(&addUser)

	suite.Assertions.NotNil(errRegister)
	suite.Assertions.Nil(user)
	suite.Assertions.Equal(http.StatusBadRequest, errRegister.Status)
}

func (suite *UserServiceTestSuite) Test_Register_AnErrorOccuredInUserRepository_ShouldReturnInternalServerError() {
	service := CreateUserService(repositories.NewErrorUserRepository(), suite.passwordHasher, suite.sessionService)
	addUser := dto.AddUserDto{
		Email:    "email@email2.com",
		Password: "Testowe123!aewe",
	}

	user, errRegister := service.Register(&addUser)

	suite.Assertions.NotNil(errRegister)
	suite.Assertions.Nil(user)
	suite.Assertions.Equal(http.StatusInternalServerError, errRegister.Status)
}
