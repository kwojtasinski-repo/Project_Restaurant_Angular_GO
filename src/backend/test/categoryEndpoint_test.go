package test

import (
	"encoding/json"
	"net/http"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
)

func (suite *IntegrationTestSuite) Test_Add_CategoryEndpoint_ShouldReturnCreated() {
	category := struct {
		Name string
	}{
		Name: "Category#1",
	}
	req := suite.CreateAuthorizedRequest("POST", "/api/categories", createPayload(category))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_Add_CategoryEndpoint_ShouldAddToDatabase() {
	category := dto.CategoryDto{
		Name: "Category#2",
	}
	req := suite.CreateAuthorizedRequest("POST", "/api/categories", createPayload(category))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
	var categoryAdded dto.CategoryDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &categoryAdded))
	suite.Require().Equal(category.Name, categoryAdded.Name)
	repository := repositories.CreateCategoryRepository(suite.database)
	categoryFromDb, err := repository.Get(categoryAdded.Id.ValueInt)
	suite.Require().Nil(err)
	suite.Require().NotNil(categoryFromDb)
	suite.Require().Equal(category.Name, categoryFromDb.Name.Value())
}
