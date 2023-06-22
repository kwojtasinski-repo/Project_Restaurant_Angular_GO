package test

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/settings"
)

func (suite *IntegrationTestSuite) Test_SignIn_IdentityEndpoint_ShouldReturnStatusOK() {
	adminUser := suite.users["admin"]
	req := suite.CreateRequest("POST", "/api/sign-in", createPayload(adminUser))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_SignIn_IdentityEndpoint_ShouldIssueCookie() {
	adminUser := suite.users["admin"]
	req := suite.CreateRequest("POST", "/api/sign-in", createPayload(adminUser))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	cookies := rec.Result().Cookies()
	suite.Require().NotEmpty(cookies)
	sessionCookie := suite.FindSessionCookie(cookies)
	suite.Require().NotNil(sessionCookie)
	suite.Require().NotEmpty(sessionCookie.Value)
	suite.Require().NotEmpty(sessionCookie.Path)
	suite.Require().Equal(settings.Location, sessionCookie.Path)
	suite.Require().Greater(sessionCookie.MaxAge, 0)
	suite.Require().Equal(settings.CookieLifeTime, sessionCookie.MaxAge)
}

func (suite *IntegrationTestSuite) Test_SignIn_IdentityEndpoint_ShouldCreateSession() {
	adminUser := suite.users["admin"]
	req := suite.CreateRequest("POST", "/api/sign-in", createPayload(adminUser))
	sessionRepository := repositories.CreateSessionRepository(suite.database)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	cookie := suite.FindSessionCookie(rec.Result().Cookies())
	req = suite.CreateRequest("GET", "/api/users/me", http.NoBody)
	req.AddCookie(cookie)
	rec = suite.SendRequest(req)
	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var userDto dto.UserDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &userDto))
	id, err := valueobjects.NewId(userDto.Id.ValueInt)
	suite.Require().Nil(err)
	sessions, err := sessionRepository.GetSessionsByUserId(*id)
	suite.Require().Nil(err)
	suite.Require().NotEmpty(sessions)
}

func (suite *IntegrationTestSuite) Test_SignUp_IdentityEndpoint_ShouldReturnCreated() {
	addUser := dto.AddUserDto{
		Email:    "kocica@test.com",
		Password: "Kocica1234!@Abc",
	}
	req := suite.CreateRequest("POST", "/api/sign-up", createPayload(addUser))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_SignUp_IdentityEndpoint_ShouldSignInAfterCreateAccount() {
	addUser := dto.AddUserDto{
		Email:    "kocica-the-best@test.com",
		Password: "Kocica1234!@Abc",
	}
	req := suite.CreateRequest("POST", "/api/sign-up", createPayload(addUser))
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
	req = suite.CreateRequest("POST", "/api/sign-in", createPayload(addUser))

	rec = suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	cookies := rec.Result().Cookies()
	suite.Require().NotEmpty(cookies)
	sessionCookie := suite.FindSessionCookie(cookies)
	suite.Require().NotNil(sessionCookie)
	suite.Require().NotEmpty(sessionCookie.Value)
}

func (suite *IntegrationTestSuite) Test_SignOut_ShouldReturnOK() {
	user := suite.users["user"]
	req := suite.CreateRequest("POST", "/api/sign-in", createPayload(user))
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	sessionCookie := suite.FindSessionCookie(rec.Result().Cookies())
	req = suite.CreateRequest("POST", "/api/sign-out", http.NoBody)
	req.AddCookie(sessionCookie)

	rec = suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_SignOut_ShouldClearSession() {
	user := suite.users["user"]
	req := suite.CreateRequest("POST", "/api/sign-in", createPayload(user))
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	sessionCookie := suite.FindSessionCookie(rec.Result().Cookies())
	req = suite.CreateRequest("POST", "/api/sign-out", http.NoBody)
	req.AddCookie(sessionCookie)

	rec = suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	req = suite.CreateRequest("GET", "/api/users/me", http.NoBody)
	req.AddCookie(sessionCookie)
	rec = suite.SendRequest(req)
	suite.Require().Equal(http.StatusUnauthorized, rec.Result().StatusCode)
	suite.Require().Empty(rec.Result().Cookies())
}
