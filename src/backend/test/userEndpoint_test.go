package test

import (
	"encoding/json"
	"net/http"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
)

func (suite *IntegrationTestSuite) Test_Me_UserEndpoint_ShouldReturnOk() {
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/users/me", http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_Me_UserEndpoint_ShouldReturnUserInfo() {
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/users/me", http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var userDto dto.UserDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &userDto))
	suite.Require().NotNil(userDto)
	suite.Require().NotEmpty(userDto.Email)
	suite.Require().NotEmpty(userDto.Role)
	suite.Require().Equal("admin", userDto.Role)
}
