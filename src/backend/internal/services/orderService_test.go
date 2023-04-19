package services

import (
	"log"
	"net/http"
	"testing"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
)

type OrderServiceTestSuite struct {
	suite.Suite
	orderRepository   repositories.OrderRepository
	cartRepository    repositories.CartRepository
	productRepository repositories.ProductRepository
	service           OrderService
}

func (suite *OrderServiceTestSuite) SetupTest() {
	log.Println("---- Setup OrderServiceTestSuite Before Each Test ----")
	suite.orderRepository = repositories.NewInMemoryOrderRepository()
	suite.productRepository = suite.createProductRepository()
	suite.cartRepository = suite.createCartRepository()
	suite.service = CreateOrderService(suite.orderRepository, suite.cartRepository, suite.productRepository)
}

func (suite *OrderServiceTestSuite) createProductRepository() repositories.ProductRepository {
	repository := repositories.NewInMemoryProductRepository()
	category, _ := entities.NewCategory(1, "Category#1")
	product1, _ := entities.NewProduct(1, "Name#1", decimal.New(10050, -2), "Description1", category)
	product2, _ := entities.NewProduct(2, "Name#2", decimal.New(55550, -2), "Description2", category)
	product3, _ := entities.NewProduct(3, "Name#3", decimal.New(12550, -2), "Description3", category)
	repository.Add(product1)
	repository.Add(product2)
	repository.Add(product3)
	return repository
}

func (suite *OrderServiceTestSuite) createCartRepository() repositories.CartRepository {
	repository := repositories.NewInMemoryCartRepository()
	user, _ := entities.NewUser(1, "test@test.com", "", "user")
	product, _ := suite.productRepository.Get(1)
	cart1, _ := entities.NewCart(1, *user, *product)
	cart2, _ := entities.NewCart(2, *user, *product)
	cart3, _ := entities.NewCart(3, *user, *product)
	repository.Add(cart1)
	repository.Add(cart2)
	repository.Add(cart3)
	return repository
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

func (suite *OrderServiceTestSuite) Test_Add_WithProducts_ShouldAddOrder() {
	addOrder := dto.AddOrderDto{
		ProductIds: []int64{1, 2, 3, 1},
		UserId:     1,
	}
	expectedPrice := decimal.New(10050, -2).Add(decimal.New(55550, -2)).Add(decimal.New(12550, -2)).Add(decimal.New(10050, -2))

	orderDto, err := suite.service.Add(addOrder)

	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(orderDto)
	suite.Assertions.NotNil(orderDto.OrderNumber)
	suite.Assertions.Equal(expectedPrice.StringFixedBank(2), orderDto.Price)
	suite.Assertions.Equal(len(addOrder.ProductIds), len(orderDto.OrderProducts))
}

func (suite *OrderServiceTestSuite) Test_Add_WithoutProducts_ShouldAddOrderWithEmptyPosistions() {
	addOrder := dto.AddOrderDto{
		ProductIds: make([]int64, 0),
		UserId:     1,
	}

	orderDto, err := suite.service.Add(addOrder)

	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(orderDto)
	suite.Assertions.NotNil(orderDto.OrderNumber)
	suite.Assertions.Equal(decimal.Zero.StringFixedBank(2), orderDto.Price)
	suite.Assertions.Empty(orderDto.OrderProducts)
}

func (suite *OrderServiceTestSuite) Test_Add_WithInvalidUserId_ShouldReturnBadRequest() {
	addOrder := dto.AddOrderDto{
		ProductIds: make([]int64, 0),
		UserId:     -1,
	}

	orderDto, err := suite.service.Add(addOrder)

	suite.Assertions.NotNil(err)
	suite.Assertions.Nil(orderDto)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
}

func (suite *OrderServiceTestSuite) Test_Add_ProductNotExists_ShouldReturnBadRequest() {
	addOrder := dto.AddOrderDto{
		ProductIds: []int64{20},
		UserId:     -1,
	}

	orderDto, err := suite.service.Add(addOrder)

	suite.Assertions.NotNil(err)
	suite.Assertions.Nil(orderDto)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
}

func (suite *OrderServiceTestSuite) Test_Add_AnErrorOccuredInProductRepository_ShouldReturnInternalErrorServer() {
	service := CreateOrderService(suite.orderRepository, suite.cartRepository, repositories.NewErrorProductRepository())
	addOrder := dto.AddOrderDto{
		ProductIds: []int64{20},
		UserId:     1,
	}

	orderDto, err := service.Add(addOrder)

	suite.Assertions.NotNil(err)
	suite.Assertions.Nil(orderDto)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *OrderServiceTestSuite) Test_Add_WithNoOrderProductsAndAnErrorOccuredInOrderRepository_ShouldReturnInternalErrorServer() {
	service := CreateOrderService(repositories.NewErrorOrderRepository(), suite.cartRepository, suite.productRepository)
	addOrder := dto.AddOrderDto{
		ProductIds: make([]int64, 0),
		UserId:     1,
	}

	orderDto, err := service.Add(addOrder)

	suite.Assertions.NotNil(err)
	suite.Assertions.Nil(orderDto)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *OrderServiceTestSuite) Test_Add_AnErrorOccuredInOrderRepository_ShouldReturnInternalErrorServer() {
	service := CreateOrderService(repositories.NewErrorOrderRepository(), suite.cartRepository, suite.productRepository)
	addOrder := dto.AddOrderDto{
		ProductIds: []int64{1},
		UserId:     1,
	}

	orderDto, err := service.Add(addOrder)

	suite.Assertions.NotNil(err)
	suite.Assertions.Nil(orderDto)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *OrderServiceTestSuite) Test_AddFromCart_ShouldAddOrderWithProductsFromCart() {
	order, err := suite.service.AddFromCart(1)

	carts, errRepo := suite.cartRepository.GetAllByUser(1)
	suite.Assertions.NotNil(order)
	suite.Assertions.Nil(err)
	suite.Assertions.Nil(errRepo)
	suite.Assertions.NotNil(carts)
	suite.Assertions.Empty(carts)
	suite.Assertions.NotNil(order.OrderNumber)
	suite.Assertions.NotEmpty(order.OrderNumber)
	suite.Assertions.NotEmpty(order.OrderProducts)
}

func (suite *OrderServiceTestSuite) Test_AddFromCart_InvalidId_ShouldReturnBadRequest() {
	order, err := suite.service.AddFromCart(-1)

	suite.Assertions.Nil(order)
	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
}

func (suite *OrderServiceTestSuite) Test_AddFromCart_EmptyCart_ShouldReturnBadRequest() {
	order, err := suite.service.AddFromCart(10)

	suite.Assertions.Nil(order)
	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
}

func (suite *OrderServiceTestSuite) Test_AddFromCart_ProductNotExists_ShouldReturnBadRequest() {
	carts, _ := suite.cartRepository.GetAllByUser(1)
	newProductId, _ := valueobjects.NewId(100)
	cart := carts[0]
	cart.Id = *newProductId
	cart.ProductId = *newProductId
	suite.cartRepository.Add(&cart)

	order, err := suite.service.AddFromCart(1)

	suite.Assertions.Nil(order)
	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
}

func (suite *OrderServiceTestSuite) Test_AddFromCart_AnErrorOccuredInProductRepository_ShouldReturnInternalServerError() {
	service := CreateOrderService(suite.orderRepository, suite.cartRepository, repositories.NewErrorProductRepository())

	order, err := service.AddFromCart(1)

	suite.Assertions.Nil(order)
	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *OrderServiceTestSuite) Test_AddFromCart_AnErrorOccuredInCartRepository_ShouldReturnInternalServerError() {
	service := CreateOrderService(suite.orderRepository, repositories.NewErrorCartRepository(), suite.productRepository)

	order, err := service.AddFromCart(1)

	suite.Assertions.Nil(order)
	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *OrderServiceTestSuite) Test_AddFromCart_AnErrorOccuredInOrderRepository_ShouldReturnInternalServerError() {
	service := CreateOrderService(repositories.NewErrorOrderRepository(), suite.cartRepository, suite.productRepository)

	order, err := service.AddFromCart(1)

	suite.Assertions.Nil(order)
	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}
