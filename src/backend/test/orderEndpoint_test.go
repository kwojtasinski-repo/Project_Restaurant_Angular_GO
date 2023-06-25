package test

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
)

func (suite *IntegrationTestSuite) Test_AddFromCart_OrderEndpoint_ShouldReturnCreated() {
	user := suite.users["test"]
	suite.AddProductToCartForUser(user)
	suite.AddProductToCartForUser(user)
	req := suite.CreateAuthorizedRequestForUser("POST", "/api/orders/from-cart", http.NoBody, user)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_AddFromCart_OrderEndpoint_ShouldAddToDatabase() {
	user := suite.users["test"]
	product1 := suite.AddProduct()
	suite.AddProductWithIdToCartForUser(product1.Id, user)
	product2 := suite.AddProduct()
	suite.AddProductWithIdToCartForUser(product2.Id, user)
	req := suite.CreateAuthorizedRequestForUser("POST", "/api/orders/from-cart", http.NoBody, user)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
	var orderCreated dto.OrderDetailsDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &orderCreated))
	suite.Require().True(containsProductInOrder(orderCreated, product1))
	suite.Require().True(containsProductInOrder(orderCreated, product2))
	suite.Require().NotNil(orderCreated.OrderNumber)
	req = suite.CreateAuthorizedRequestForUser("GET", "/api/orders/"+orderCreated.Id.Value, http.NoBody, suite.users["test"])
	rec = suite.SendRequest(req)
	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var orderAdded dto.OrderDetailsDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &orderAdded))
	// returns time in utc but by default it is converted to local time and truncate seconds nanoseconds
	suite.Require().Equal(orderCreated.Created.UTC().Truncate(time.Second), orderAdded.Created)
	suite.Require().Equal(orderCreated.OrderNumber, orderAdded.OrderNumber)
	suite.Require().Equal(len(orderCreated.OrderProducts), len(orderAdded.OrderProducts))
	suite.Require().True(containsProductInOrder(orderAdded, product1))
	suite.Require().True(containsProductInOrder(orderAdded, product2))
}

func (suite *IntegrationTestSuite) Test_AddFromCart_ForDifferentUser_OrderEndpoint_ShouldReturnBadRequest() {
	user := suite.users["test"]
	product1 := suite.AddProduct()
	suite.AddProductWithIdToCart(product1.Id)
	product2 := suite.AddProduct()
	suite.AddProductWithIdToCart(product2.Id)
	req := suite.CreateAuthorizedRequestForUser("POST", "/api/orders/from-cart", http.NoBody, user)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusBadRequest, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_AddFromCart_ForDifferentUser_OrderEndpoint_ShouldErrorResponseWithMessage() {
	user := suite.users["test"]
	product1 := suite.AddProduct()
	suite.AddProductWithIdToCart(product1.Id)
	product2 := suite.AddProduct()
	suite.AddProductWithIdToCart(product2.Id)
	req := suite.CreateAuthorizedRequestForUser("POST", "/api/orders/from-cart", http.NoBody, user)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusBadRequest, rec.Result().StatusCode)
	errorResponse := suite.getErrorResponse(rec)
	suite.Require().NotEmpty(errorResponse.Errors)
	suite.Require().Contains(errorResponse.Errors, "'Cart' is empty, add something before create an 'Order'")
}

func (suite *IntegrationTestSuite) Test_Add_OrderEndpoint_ShouldReturnCreated() {
	user := suite.users["user"]
	product1 := suite.AddProduct()
	product2 := suite.AddProduct()
	addOrder := dto.AddOrderDto{
		ProductIds: []dto.IdObject{product1.Id, product2.Id},
	}
	req := suite.CreateAuthorizedRequestForUser("POST", "/api/orders", createPayload(addOrder), user)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
	var orderCreated dto.OrderDetailsDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &orderCreated))
	suite.Require().True(containsProductInOrder(orderCreated, product1))
	suite.Require().True(containsProductInOrder(orderCreated, product2))
	suite.Require().NotNil(orderCreated.OrderNumber)
	orderRepository := repositories.CreateOrderRepository(suite.database)
	order, err := orderRepository.Get(orderCreated.Id.ValueInt)
	suite.Require().Nil(err)
	suite.Require().NotNil(order)
	suite.Require().Equal(order.OrderNumber.String(), orderCreated.OrderNumber)
}

func (suite *IntegrationTestSuite) Test_Add_OrderEndpoint_ShouldAddToDatabase() {
	user := suite.users["user"]
	product1 := suite.AddProduct()
	product2 := suite.AddProduct()
	addOrder := dto.AddOrderDto{
		ProductIds: []dto.IdObject{product1.Id, product2.Id},
	}
	req := suite.CreateAuthorizedRequestForUser("POST", "/api/orders", createPayload(addOrder), user)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_GetMyOrders_OrderEndpoint_ShouldReturnOkWithOrders() {
	user := suite.users["test"]
	addOrder(suite, &user)
	req := suite.CreateAuthorizedRequestForUser("GET", "/api/orders/my", http.NoBody, user)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var orders []dto.OrderDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &orders))
	suite.Require().NotEmpty(orders)
}

func (suite *IntegrationTestSuite) Test_GetAllOrders_WithNonAdminUser_OrderEndpoint_ShouldReturnForbidden() {
	user := suite.users["test"]
	addOrder(suite, &user)
	req := suite.CreateAuthorizedRequestForUser("GET", "/api/orders", http.NoBody, user)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusForbidden, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) Test_GetAllOrders_OrderEndpoint_ShouldReturnOkWithOrders() {
	addOrder(suite, nil)
	req := suite.CreateAuthorizedRequest("GET", "/api/orders", http.NoBody)

	rec := suite.SendRequest(req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var orders []dto.OrderDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &orders))
	suite.Require().NotEmpty(orders)
}

func containsProductInOrder(order dto.OrderDetailsDto, product dto.ProductDetailsDto) bool {
	for _, orderProduct := range order.OrderProducts {
		if orderProduct.ProductId.ValueInt == product.Id.ValueInt {
			return true
		}
	}

	return false
}

func addOrder(suite *IntegrationTestSuite, addUser *dto.AddUserDto) dto.OrderDetailsDto {
	product1 := suite.AddProduct()
	product2 := suite.AddProduct()
	addOrder := dto.AddOrderDto{
		ProductIds: []dto.IdObject{product1.Id, product2.Id},
	}
	var req *http.Request
	if addUser != nil {
		req = suite.CreateAuthorizedRequestForUser("POST", "/api/orders", createPayload(addOrder), *addUser)
	} else {
		req = suite.CreateAuthorizedRequest("POST", "/api/orders", createPayload(addOrder))
	}
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
	var orderCreated dto.OrderDetailsDto
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &orderCreated))
	return orderCreated
}
