package test

import (
	"encoding/json"
	"net/http"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
)

func (suite *IntegrationTestSuite) Test_AddToCart_CartEndpoint_ShouldCreated() {
	product := suite.AddProduct()
	addCart := dto.AddCart{
		ProductId: product.Id,
	}
	req := suite.CreateAuthorizedRequest(http.MethodPost, "/api/carts", createPayload(addCart))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_AddToCart_CartEndpoint_ShouldAddToDatabase() {
	product := suite.AddProduct()
	addCart := dto.AddCart{
		ProductId: product.Id,
	}
	req := suite.CreateAuthorizedRequest(http.MethodPost, "/api/carts", createPayload(addCart))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
	myCart := suite.getMyCart()
	suite.Require().NotEmpty(myCart)
}

func (suite *IntegrationTestSuite) Test_DeleteFromCart_CartEndpoint_ShouldNoContent() {
	suite.AddProductToCart()
	suite.AddProductToCart()
	myCart := suite.getMyCart()
	lastElement := myCart[len(myCart)-1]
	req := suite.CreateAuthorizedRequest(http.MethodDelete, "/api/carts/"+lastElement.Id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusNoContent, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_DeleteFromCart_CartEndpoint_ShouldRemoveFromDatabase() {
	suite.AddProductToCart()
	myCart := suite.getMyCart()
	lastElement := myCart[len(myCart)-1]
	req := suite.CreateAuthorizedRequest(http.MethodDelete, "/api/carts/"+lastElement.Id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusNoContent, rec.Result().StatusCode)
	myCartAfterDelete := suite.getMyCart()
	suite.Require().Greater(len(myCart), len(myCartAfterDelete))
}

func (suite *IntegrationTestSuite) getMyCart() []dto.CartDto {
	myCartResponse := suite.SendRequest(suite.CreateAuthorizedRequest("GET", "/api/carts/my", http.NoBody))
	var myCart []dto.CartDto
	suite.Require().Nil(json.Unmarshal(myCartResponse.Body.Bytes(), &myCart))
	return myCart
}
