package services

import (
	"log"
	"net/http"
	"testing"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
)

type CartServiceTestSuite struct {
	suite.Suite
	cartRepository    repositories.CartRepository
	productRepository repositories.ProductRepository
	userRepository    repositories.UserRepository
	service           CartService
}

func (suite *CartServiceTestSuite) SetupTest() {
	log.Println("---- Setup CartServiceTestSuite Before Each Test ----")
	suite.productRepository = suite.createProductRepository()
	suite.cartRepository = suite.createCartRepository()
	suite.service = CreateCartService(suite.cartRepository, suite.productRepository)
}

func (suite *CartServiceTestSuite) createProductRepository() repositories.ProductRepository {
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

func (suite *CartServiceTestSuite) createCartRepository() repositories.CartRepository {
	repository := repositories.NewInMemoryCartRepository()
	product, _ := suite.productRepository.Get(1)
	user, _ := suite.userRepository.Get(1)
	cart1, _ := entities.NewCart(1, *user, *product)
	cart2, _ := entities.NewCart(2, *user, *product)
	cart3, _ := entities.NewCart(3, *user, *product)
	repository.Add(cart1)
	repository.Add(cart2)
	repository.Add(cart3)
	return repository
}

func TestCartServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

func (suite *CartServiceTestSuite) Test_AddToCart_ValidProductAndUserId_ShoudAddProductToCart() {
	addCart := dto.AddCart{
		ProductId: 1,
		UserId:    2,
	}

	err := suite.service.AddToCart(addCart)

	carts, errRepo := suite.service.GetMyCart(addCart.UserId)
	suite.Assertions.Nil(err)
	suite.Assertions.Nil(errRepo)
	suite.Assertions.NotNil(carts)
	suite.Assertions.NotEmpty(carts)
}

func (suite *CartServiceTestSuite) Test_AddToCart_InvalidProductId_ShoudReturnBadRequest() {
	addCart := dto.AddCart{
		ProductId: 1000,
		UserId:    2,
	}

	err := suite.service.AddToCart(addCart)

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
}

func (suite *CartServiceTestSuite) Test_AddToCart_InvalidUserId_ShoudReturnBadRequest() {
	addCart := dto.AddCart{
		ProductId: 1,
		UserId:    -2000,
	}

	err := suite.service.AddToCart(addCart)

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
}

func (suite *CartServiceTestSuite) Test_AddToCart_AnErrorOccuredInProductRepository_ShoudReturnInternalServerError() {
	service := CreateCartService(suite.cartRepository, repositories.NewErrorProductRepository())
	addCart := dto.AddCart{
		ProductId: 1,
		UserId:    2,
	}

	err := service.AddToCart(addCart)

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *CartServiceTestSuite) Test_AddToCart_AnErrorOccuredInCartRepository_ShoudReturnInternalServerError() {
	service := CreateCartService(repositories.NewErrorCartRepository(), suite.productRepository)
	addCart := dto.AddCart{
		ProductId: 1,
		UserId:    2,
	}

	err := service.AddToCart(addCart)

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *CartServiceTestSuite) Test_RemoveFromCart_ValidCartIdAndUserId_ShoudRemoveProductFromCart() {
	userId := 1
	cartId := 1

	err := suite.service.RemoveFromCart(int64(cartId), int64(userId))

	carts, errRepo := suite.service.GetMyCart(int64(userId))
	suite.Assertions.Nil(err)
	suite.Assertions.Nil(errRepo)
	suite.Assertions.NotNil(carts)
	suite.Assertions.NotEmpty(carts)
	exists := false
	for _, cart := range carts {
		if cart.Id == int64(cartId) {
			exists = true
		}
	}
	suite.Assertions.False(exists)
	suite.Assertions.Equal(2, len(carts))
}

func (suite *CartServiceTestSuite) Test_RemoveFromCart_InvalidCartId_ShoudReturnNotFound() {
	userId := 1
	cartId := 10000

	err := suite.service.RemoveFromCart(int64(cartId), int64(userId))

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusNotFound)
}

func (suite *CartServiceTestSuite) Test_RemoveFromCart_InvalidUserId_ShoudReturnBadRequest() {
	userId := 10000
	cartId := 1

	err := suite.service.RemoveFromCart(int64(cartId), int64(userId))

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusBadRequest)
}

func (suite *CartServiceTestSuite) Test_RemoveFromCart_AnErrorOccuredInCartRepository_ShoudReturnInternalServerError() {
	service := CreateCartService(repositories.NewErrorCartRepository(), suite.productRepository)
	userId := 1
	cartId := 1

	err := service.RemoveFromCart(int64(cartId), int64(userId))

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}

func (suite *CartServiceTestSuite) Test_GetMyCart_ValidUserId_ShoudShowAllProductsInCart() {
	carts, err := suite.service.GetMyCart(1)

	suite.Assertions.Nil(err)
	suite.Assertions.NotNil(carts)
	suite.Assertions.NotEmpty(carts)
	suite.Assertions.Equal(3, len(carts))
}

func (suite *CartServiceTestSuite) Test_GetMyCart_AnErrorOccuredInCartRepository_ShoudReturnInternalServerError() {
	service := CreateCartService(repositories.NewErrorCartRepository(), suite.productRepository)

	carts, err := service.GetMyCart(1)

	suite.Assertions.NotNil(err)
	suite.Assertions.NotNil(carts)
	suite.Assertions.Empty(carts)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}
