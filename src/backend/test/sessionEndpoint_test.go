package test

import (
	"encoding/json"
	"net/http"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
)

func (suite *IntegrationTestSuite) Test_GetAllUserSession_SessionEndpoint_ShouldReturnOkWithAllSessions() {
	login(suite, nil)
	user := login(suite, nil)
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/sessions/"+user.Id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var sessions []dto.SessionDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &sessions))
	suite.Require().NotEmpty(sessions)
}

func (suite *IntegrationTestSuite) Test_GetAllUserSession_WithNonAdminUser_SessionEndpoint_ShouldReturnForbidden() {
	userCredendtials := suite.users["user"]
	login(suite, &userCredendtials)
	user := login(suite, &userCredendtials)
	req := suite.CreateAuthorizedRequestForUser(http.MethodGet, "/api/sessions/"+user.Id.Value, http.NoBody, userCredendtials)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusForbidden, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_RevokeAllUserSessions_SessionEndpoint_ShouldReturnOkAndClearSessions() {
	userCredentials := suite.users["revokeSessionUser"]
	login(suite, &userCredentials)
	user := login(suite, &userCredentials)
	sessions := getAllUserSessions(suite, &user)
	suite.Require().NotEmpty(sessions)
	req := suite.CreateAuthorizedRequestForUser(http.MethodDelete, "/api/sessions/"+user.Id.Value, http.NoBody, userCredentials)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusNoContent, rec.Result().StatusCode)
	sessionsAfterDelete := getAllUserSessions(suite, &user)
	suite.Require().Greater(len(sessions), len(sessionsAfterDelete))
	suite.Require().Empty(sessionsAfterDelete)
}

func (suite *IntegrationTestSuite) Test_RevokeAllUserSessions_WithNonAdminUser_SessionEndpoint_ShouldReturnForbidden() {
	userCredentials := suite.users["user"]
	login(suite, &userCredentials)
	user := login(suite, &userCredentials)
	req := suite.CreateAuthorizedRequestForUser(http.MethodDelete, "/api/sessions/"+user.Id.Value, http.NoBody, userCredentials)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusForbidden, rec.Result().StatusCode)
}

func login(suite *IntegrationTestSuite, user *dto.AddUserDto) dto.UserDto {
	var req *http.Request
	var userCredendtials *dto.AddUserDto
	if user != nil {
		userCredendtials = &dto.AddUserDto{
			Email:    user.Email,
			Password: user.Password,
		}
	} else {
		u := suite.users["admin"]
		userCredendtials = &u
	}
	req = suite.CreateRequest(http.MethodPost, "/api/sign-in", createPayload(*userCredendtials))
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	req = suite.CreateAuthorizedRequestForUser(http.MethodGet, "/api/users/me", http.NoBody, *userCredendtials)
	rec = suite.SendRequest(req)
	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var userDto dto.UserDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &userDto))
	return userDto
}

func getAllUserSessions(suite *IntegrationTestSuite, user *dto.UserDto) []dto.SessionDto {
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/sessions/"+user.Id.Value, http.NoBody)
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var sessions []dto.SessionDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &sessions))
	return sessions
}
