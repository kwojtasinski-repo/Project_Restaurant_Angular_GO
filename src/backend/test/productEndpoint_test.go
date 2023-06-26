package test

import (
	"encoding/json"
	"net/http"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/shopspring/decimal"
)

func (suite *IntegrationTestSuite) Test_Add_ProductEndpoint_ShouldReturnCreated() {
	id, _ := dto.NewIntIdObject(1)
	product := struct {
		Name        string
		Description string
		CategoryId  string
		Price       decimal.Decimal
	}{
		Name:        "Product#1",
		Description: "Description#123456789",
		CategoryId:  id.Value,
		Price:       decimal.New(100, 1),
	}
	req := suite.CreateAuthorizedRequest(http.MethodPost, "/api/products", createPayload(product))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_Add_ProductEndpoint_ShouldAddToDatabase() {
	id, _ := dto.NewIntIdObject(1)
	product := dto.AddProductDto{
		Name:        "Product#1",
		Description: "Description#123456789",
		CategoryId:  *id,
		Price:       decimal.New(100, 1),
	}
	req := suite.CreateAuthorizedRequest(http.MethodPost, "/api/products", createPayload(product))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
	var productAdded dto.ProductDetailsDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &productAdded))
	suite.Require().Equal(product.Name, productAdded.Name)
	suite.Require().Equal(product.Description, productAdded.Description)
	suite.Require().Equal(product.CategoryId.ValueInt, productAdded.Category.Id.ValueInt)
	suite.Require().Equal(product.Price.StringFixedBank(2), productAdded.Price)
	repository := repositories.CreateProductRepository(suite.database)
	productFromDb, err := repository.Get(productAdded.Id.ValueInt)
	suite.Require().Nil(err)
	suite.Require().NotNil(productFromDb)
	suite.Require().Equal(product.Name, productFromDb.Name.Value())
	suite.Require().Equal(product.Description, productFromDb.Description.Value())
	suite.Require().Equal(product.CategoryId.ValueInt, productFromDb.Category.Id.Value())
	suite.Require().Equal(product.Price.StringFixedBank(2), productFromDb.Price.Value().StringFixedBank(2))
}

func (suite *IntegrationTestSuite) Test_Update_ProductEndpoint_ShouldReturnOk() {
	product := suite.AddProduct()
	product.Name = "Name#1#Product"
	product.Description = "Afterwards_NextStep_1234"
	price, err := decimal.NewFromString(product.Price)
	suite.Require().Nil(err)
	req := suite.CreateAuthorizedRequest(http.MethodPut, "/api/products/"+product.Id.Value, createPayload(dto.UpdateProductDto{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		CategoryId:  product.Category.Id,
		Price:       price,
	}))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_Update_ProductEndpoint_ShouldAddToDatabase() {
	product := suite.AddProduct()
	product.Name = "Name#1#Product"
	product.Description = "Afterwards_NextStep_1234"
	price := decimal.New(150, 1)
	req := suite.CreateAuthorizedRequest(http.MethodPut, "/api/products/"+product.Id.Value, createPayload(dto.UpdateProductDto{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		CategoryId:  product.Category.Id,
		Price:       price,
	}))

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var productDetails dto.ProductDetailsDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &productDetails))
	suite.Require().Equal(product.Name, productDetails.Name)
	suite.Require().Equal(product.Description, productDetails.Description)
	suite.Require().Equal(price.StringFixedBank(2), productDetails.Price)
	repository := repositories.CreateProductRepository(suite.database)
	productFromDb, err := repository.Get(productDetails.Id.ValueInt)
	suite.Require().Nil(err)
	suite.Require().NotNil(productFromDb)
	suite.Require().Equal(product.Name, productFromDb.Name.Value())
	suite.Require().Equal(product.Description, productFromDb.Description.Value())
	suite.Require().Equal(price.StringFixedBank(2), productFromDb.Price.Value().StringFixedBank(2))
}

func (suite *IntegrationTestSuite) Test_Delete_ProductEndpoint_ShouldNoContent() {
	product := suite.AddProduct()
	req := suite.CreateAuthorizedRequest(http.MethodDelete, "/api/products/"+product.Id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusNoContent, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_Delete_ProductEndpoint_ShouldDeleteFromDatabase() {
	product := suite.AddProduct()
	req := suite.CreateAuthorizedRequest(http.MethodDelete, "/api/products/"+product.Id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusNoContent, rec.Result().StatusCode)
	repository := repositories.CreateProductRepository(suite.database)
	productFromDb, err := repository.Get(product.Id.ValueInt)
	suite.Require().Nil(err)
	suite.Require().NotNil(productFromDb)
	suite.Require().Equal(true, productFromDb.Deleted)
}

func (suite *IntegrationTestSuite) Test_Get_ProductEndpoint_ShouldReturnOkWithProduct() {
	product := suite.AddProduct()
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/products/"+product.Id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var productFromRequest dto.ProductDetailsDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &productFromRequest))
	suite.Require().NotNil(productFromRequest)
	suite.Require().Equal(product.Name, productFromRequest.Name)
	suite.Require().Equal(product.Name, productFromRequest.Name)
	suite.Require().Equal(product.Description, productFromRequest.Description)
	suite.Require().Equal(product.Category.Id.ValueInt, productFromRequest.Category.Id.ValueInt)
	suite.Require().Equal(product.Price, productFromRequest.Price)
}

func (suite *IntegrationTestSuite) Test_Get_NotExistProduct_ProductEndpoint_ShouldReturnNotFound() {
	id, err := dto.NewIntIdObject(1000)
	suite.Require().Nil(err)
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/products/"+id.Value, http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusNotFound, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_GetAll_ProductEndpoint_ShouldReturnOkWithCategories() {
	suite.AddProduct()
	suite.AddProduct()
	req := suite.CreateAuthorizedRequest(http.MethodGet, "/api/products", http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var products []dto.ProductDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &products))
	suite.Require().NotNil(products)
	suite.Require().Greater(len(products), 1)
}
