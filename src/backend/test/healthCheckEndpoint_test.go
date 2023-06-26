package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (suite *IntegrationTestSuite) Test_SendRequest_HealthCheckEndpoint_ShouldReturnStatusOkAndContainsWelcomeText() {
	req, err := http.NewRequest(http.MethodGet, "/api", http.NoBody)
	suite.Require().NoError(err)
	rec := httptest.NewRecorder()
	welcomeText := "Welcome to Restaurant API"

	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var welcomeTextFromApi string
	suite.Require().NoError(json.NewDecoder(rec.Result().Body).Decode(&welcomeTextFromApi))
	suite.Require().Contains(welcomeTextFromApi, welcomeText)
}
