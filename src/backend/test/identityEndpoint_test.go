package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/goccy/go-json"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/settings"
)

func (suite *IntegrationTestSuite) Test_SignIn_IdentityEndpoint_ShouldReturnStatusOK() {
	adminUser := suite.users["admin"]
	req, err := http.NewRequest("POST", "/api/sign-in", createPayload(adminUser))
	suite.Require().NoError(err)
	rec := httptest.NewRecorder()

	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_SignIn_IdentityEndpoint_ShouldIssueCookie() {
	adminUser := suite.users["admin"]
	req, err := http.NewRequest("POST", "/api/sign-in", createPayload(adminUser))
	suite.Require().NoError(err)
	rec := httptest.NewRecorder()

	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	cookies := rec.Result().Cookies()
	suite.Require().NotEmpty(cookies)
	sessionCookie := findSessionCookie(cookies)
	suite.Require().NotNil(sessionCookie)
	suite.Require().NotEmpty(sessionCookie.Value)
	suite.Require().NotEmpty(sessionCookie.Path)
	suite.Require().Equal(settings.Location, sessionCookie.Path)
	suite.Require().Greater(sessionCookie.MaxAge, 0)
	suite.Require().Equal(settings.CookieLifeTime, sessionCookie.MaxAge)
}

func (suite *IntegrationTestSuite) Test_SignIn_IdentityEndpoint_ShouldCreateSession() {
	adminUser := suite.users["admin"]
	req, err := http.NewRequest("POST", "/api/sign-in", createPayload(adminUser))
	suite.Require().NoError(err)
	rec := httptest.NewRecorder()
	sessionRepository := repositories.CreateSessionRepository(suite.database)

	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	cookies := rec.Result().Cookies()
	cookie := findSessionCookie(cookies)
	req, err = http.NewRequest("GET", "/api/users/me", http.NoBody)
	req.AddCookie(cookie)
	suite.Require().NoError(err)
	rec = httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)
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
	req, err := http.NewRequest("POST", "/api/sign-up", createPayload(addUser))
	suite.Require().NoError(err)
	rec := httptest.NewRecorder()

	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_SignUp_IdentityEndpoint_ShouldSignInAfterCreateAccount() {
	addUser := dto.AddUserDto{
		Email:    "kocica-the-best@test.com",
		Password: "Kocica1234!@Abc",
	}
	req, err := http.NewRequest("POST", "/api/sign-up", createPayload(addUser))
	suite.Require().NoError(err)
	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)
	req, err = http.NewRequest("POST", "/api/sign-in", createPayload(addUser))
	suite.Require().NoError(err)
	rec = httptest.NewRecorder()

	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	cookies := rec.Result().Cookies()
	suite.Require().NotEmpty(cookies)
	sessionCookie := findSessionCookie(cookies)
	suite.Require().NotNil(sessionCookie)
	suite.Require().NotEmpty(sessionCookie.Value)
}

func (suite *IntegrationTestSuite) Test_SignOut_ShouldReturnOK() {
	user := suite.users["user"]
	req, err := http.NewRequest("POST", "/api/sign-in", createPayload(user))
	suite.Require().NoError(err)
	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)
	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	cookies := rec.Result().Cookies()
	suite.Require().NotEmpty(cookies)
	sessionCookie := findSessionCookie(cookies)
	req, err = http.NewRequest("POST", "/api/sign-out", http.NoBody)
	suite.Require().NoError(err)
	rec = httptest.NewRecorder()
	req.AddCookie(sessionCookie)

	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_SignOut_ShouldClearSession() {
	user := suite.users["user"]
	req, err := http.NewRequest("POST", "/api/sign-in", createPayload(user))
	suite.Require().NoError(err)
	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)
	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	cookies := rec.Result().Cookies()
	suite.Require().NotEmpty(cookies)
	sessionCookie := findSessionCookie(cookies)
	req, err = http.NewRequest("POST", "/api/sign-out", http.NoBody)
	suite.Require().NoError(err)
	rec = httptest.NewRecorder()
	req.AddCookie(sessionCookie)

	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	req, err = http.NewRequest("GET", "/api/users/me", http.NoBody)
	req.AddCookie(sessionCookie)
	suite.Require().NoError(err)
	rec = httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)
	suite.Require().Equal(http.StatusUnauthorized, rec.Result().StatusCode)
}

func findSessionCookie(cookies []*http.Cookie) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == settings.CookieSessionName {
			return cookie
		}
	}

	return nil
}
