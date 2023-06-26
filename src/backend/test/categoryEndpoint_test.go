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
	req := suite.CreateAuthorizedRequest(http.MethodPost, "/api/categories", createPayload(category))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_Add_CategoryEndpoint_ShouldAddToDatabase() {
	category := dto.CategoryDto{
		Name: "Category#2",
	}
	req := suite.CreateAuthorizedRequest(http.MethodPost, "/api/categories", createPayload(category))

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

func (suite *IntegrationTestSuite) Test_Update_CategoryEndpoint_ShouldReturnOk() {
	category := suite.AddCategory()
	category.Name = "Name#1#Category"
	req := suite.CreateAuthorizedRequest(http.MethodPut, "/api/categories/"+category.Id.Value, createPayload(category))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_Update_CategoryEndpoint_ShouldAddToDatabase() {
	category := suite.AddCategory()
	category.Name = "Name#2#Category"
	req := suite.CreateAuthorizedRequest(http.MethodPut, "/api/categories/"+category.Id.Value, createPayload(category))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var categoryUpdated dto.CategoryDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &categoryUpdated))
	suite.Require().Equal(category.Name, categoryUpdated.Name)
	repository := repositories.CreateCategoryRepository(suite.database)
	categoryFromDb, err := repository.Get(categoryUpdated.Id.ValueInt)
	suite.Require().Nil(err)
	suite.Require().NotNil(categoryFromDb)
	suite.Require().Equal(category.Name, categoryFromDb.Name.Value())
}

func (suite *IntegrationTestSuite) Test_Delete_CategoryEndpoint_ShouldNoContent() {
	category := suite.AddCategory()
	req := suite.CreateAuthorizedRequest(http.MethodDelete, "/api/categories/"+category.Id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusNoContent, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_Delete_CategoryEndpoint_ShouldDeleteFromDatabase() {
	category := suite.AddCategory()
	req := suite.CreateAuthorizedRequest(http.MethodDelete, "/api/categories/"+category.Id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusNoContent, rec.Result().StatusCode)
	repository := repositories.CreateCategoryRepository(suite.database)
	categoryFromDb, err := repository.Get(category.Id.ValueInt)
	suite.Require().Nil(err)
	suite.Require().NotNil(categoryFromDb)
	suite.Require().Equal(true, categoryFromDb.Deleted)
}

func (suite *IntegrationTestSuite) Test_Get_CategoryEndpoint_ShouldReturnOkWithCategory() {
	category := suite.AddCategory()
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/categories/"+category.Id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var categoryFromRequest dto.CategoryDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &categoryFromRequest))
	suite.Require().NotNil(categoryFromRequest)
	suite.Require().Equal(category.Name, categoryFromRequest.Name)
}

func (suite *IntegrationTestSuite) Test_Get_NotExistCategory_CategoryEndpoint_ShouldReturnNotFound() {
	id, err := dto.NewIntIdObject(1000)
	suite.Require().Nil(err)
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/categories/"+id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusNotFound, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_GetAll_CategoryEndpoint_ShouldReturnOkWithCategories() {
	suite.AddCategory()
	suite.AddCategory()
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/categories", http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var categories []dto.CategoryDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &categories))
	suite.Require().NotNil(categories)
	suite.Require().Greater(len(categories), 1)
}
