package api

import (
	"database/sql"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/config"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/services"
)

var passwordHasher = services.CreatePassworHasherService()
var configuration config.Config
var objectsPerRequest = make(map[string]interface{})

func InitObjectCreator(configFile config.Config) {
	configuration = configFile
}

func CreateDatabaseConnection() (*sql.DB, error) {
	service := objectsPerRequest["database"]

	if service != nil {
		return service.(*sql.DB), nil
	}

	database, err := sql.Open("mysql", configuration.Database.Username+":"+configuration.Database.Password+"@tcp(localhost:3306)/"+configuration.Database.Name+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	objectsPerRequest["database"] = database

	return database, nil
}

func ResetObjectCreator() {
	objectsPerRequest = make(map[string]interface{})
}

func createSessionService() (services.SessionService, error) {
	service := objectsPerRequest["services.SessionService"]

	if service != nil {
		return service.(services.SessionService), nil
	}

	sessionRepository, err := createSessionRepository()
	if err != nil {
		return nil, err
	}

	userRepository, err := createUserRepository()
	if err != nil {
		return nil, err
	}

	sessionService := services.CreateSessionService(sessionRepository, userRepository)
	objectsPerRequest["services.SessionService"] = sessionService
	return sessionService, nil
}

func createProductService() (services.ProductService, error) {
	service := objectsPerRequest["services.ProductService"]

	if service != nil {
		return service.(services.ProductService), nil
	}

	productRepository, err := createProductRepository()
	if err != nil {
		return nil, err
	}

	categoryRepository, err := createCategoryRepository()
	if err != nil {
		return nil, err
	}

	productService := services.CreateProductService(productRepository, categoryRepository)
	objectsPerRequest["services.ProductService"] = productService
	return productService, nil
}

func createCategoryService() (services.CategoryService, error) {
	service := objectsPerRequest["services.CategoryService"]

	if service != nil {
		return service.(services.CategoryService), nil
	}

	categoryRepository, err := createCategoryRepository()
	if err != nil {
		return nil, err
	}

	categoryService := services.CreateCategoryService(categoryRepository)
	objectsPerRequest["services.CategoryService"] = categoryService
	return categoryService, nil
}

func createOrderService() (services.OrderService, error) {
	service := objectsPerRequest["services.OrderService"]

	if service != nil {
		return service.(services.OrderService), nil
	}

	orderRepository, err := createOrderRepository()
	if err != nil {
		return nil, err
	}

	cartRepository, err := createCartRepository()
	if err != nil {
		return nil, err
	}

	productRepository, err := createProductRepository()
	if err != nil {
		return nil, err
	}

	orderService := services.CreateOrderService(orderRepository, cartRepository, productRepository)
	objectsPerRequest["services.OrderService"] = orderService
	return orderService, nil
}

func createUserService() (services.UserService, error) {
	service := objectsPerRequest["services.UserService"]

	if service != nil {
		return service.(services.UserService), nil
	}

	sessionService, err := createSessionService()
	if err != nil {
		return nil, err
	}

	userRepository, err := createUserRepository()
	if err != nil {
		return nil, err
	}

	userService := services.CreateUserService(userRepository, passwordHasher, sessionService)
	objectsPerRequest["services.UserService"] = userService
	return userService, nil
}

func createCartService() (services.CartService, error) {
	service := objectsPerRequest["services.CartService"]

	if service != nil {
		return service.(services.CartService), nil
	}

	cartRepository, err := createCartRepository()
	if err != nil {
		return nil, err
	}

	productRepository, err := createProductRepository()
	if err != nil {
		return nil, err
	}

	userRepository, err := createUserRepository()
	if err != nil {
		return nil, err
	}

	cartService := services.CreateCartService(cartRepository, productRepository, userRepository)
	objectsPerRequest["services.CartService"] = cartService
	return cartService, nil
}

func createCategoryRepository() (repositories.CategoryRepository, error) {
	repo := objectsPerRequest["repositories.CategoryRepository"]

	if repo != nil {
		return repo.(repositories.CategoryRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	categoryRepository := repositories.CreateCategoryRepository(*databaseConnection)
	objectsPerRequest["repositories.CategoryRepository"] = categoryRepository
	return categoryRepository, nil
}

func createProductRepository() (repositories.ProductRepository, error) {
	repo := objectsPerRequest["repositories.ProductRepository"]

	if repo != nil {
		return repo.(repositories.ProductRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	productRepository := repositories.CreateProductRepository(*databaseConnection)
	objectsPerRequest["repositories.ProductRepository"] = productRepository
	return productRepository, nil
}

func createCartRepository() (repositories.CartRepository, error) {
	repo := objectsPerRequest["repositories.CartRepository"]

	if repo != nil {
		return repo.(repositories.CartRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	cartRepository := repositories.CreateCartRepository(*databaseConnection)
	objectsPerRequest["repositories.CartRepository"] = cartRepository
	return cartRepository, nil
}

func createSessionRepository() (repositories.SessionRepository, error) {
	repo := objectsPerRequest["repositories.OrderRepository"]

	if repo != nil {
		return repo.(repositories.SessionRepository), nil
	}

	database, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	sessionRepository := repositories.CreateSessionRepository(*database)
	objectsPerRequest["repositories.SessionRepository"] = sessionRepository
	return sessionRepository, nil
}

func createOrderRepository() (repositories.OrderRepository, error) {
	repo := objectsPerRequest["repositories.OrderRepository"]

	if repo != nil {
		return repo.(repositories.OrderRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	orderRepository := repositories.CreateOrderRepository(*databaseConnection)
	objectsPerRequest["repositories.OrderRepository"] = orderRepository
	return orderRepository, nil
}

func createUserRepository() (repositories.UserRepository, error) {
	repo := objectsPerRequest["repositories.UserRepository"]

	if repo != nil {
		return repo.(repositories.UserRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	userRepository := repositories.CreateUserRepository(*databaseConnection)
	objectsPerRequest["repositories.UserRepository"] = userRepository
	return userRepository, nil
}
