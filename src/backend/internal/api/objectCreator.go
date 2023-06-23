package api

import (
	"database/sql"
	"errors"
	"log"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/config"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/services"
	"github.com/speps/go-hashids/v2"
	"golang.org/x/sync/syncmap"
)

var passwordHasher = services.CreatePassworHasherService()
var configuration config.Config
var objectsPerRequest = syncmap.Map{}
var sqlDb *sql.DB
var hashId *hashids.HashID

func InitObjectCreator(configFile config.Config) {
	configuration = configFile
}

func CreateDatabaseConnection() (*sql.DB, error) {
	if sqlDb != nil {
		return sqlDb, nil
	}

	var err error
	sqlDb, err = sql.Open("mysql", configuration.Database.Username+":"+configuration.Database.Password+"@tcp(localhost:3306)/"+configuration.Database.Name+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	log.Println("setting limits connections to database to 5")
	sqlDb.SetMaxOpenConns(5)
	sqlDb.SetMaxIdleConns(5)

	return sqlDb, nil
}

func ResetObjectCreator() {
	objectsPerRequest = syncmap.Map{}
}

func createSessionService() (services.SessionService, error) {
	service, exists := objectsPerRequest.Load("services.SessionService")

	if exists {
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
	objectsPerRequest.Store("services.SessionService", sessionService)
	return sessionService, nil
}

func CreateProductService() (services.ProductService, error) {
	service, exists := objectsPerRequest.Load("services.ProductService")

	if exists {
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
	objectsPerRequest.Store("services.ProductService", productService)
	return productService, nil
}

func CreateCategoryService() (services.CategoryService, error) {
	service, exists := objectsPerRequest.Load("services.CategoryService")

	if exists {
		return service.(services.CategoryService), nil
	}

	categoryRepository, err := createCategoryRepository()
	if err != nil {
		return nil, err
	}

	categoryService := services.CreateCategoryService(categoryRepository)
	objectsPerRequest.Store("services.CategoryService", categoryService)
	return categoryService, nil
}

func createOrderService() (services.OrderService, error) {
	service, exists := objectsPerRequest.Load("services.OrderService")

	if exists {
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

	sessionProvider, err := getSessionProvider()
	if err != nil {
		return nil, err
	}

	orderService := services.CreateOrderService(orderRepository, cartRepository, productRepository, *sessionProvider)
	objectsPerRequest.Store("services.OrderService", orderService)
	return orderService, nil
}

func CreateUserService() (services.UserService, error) {
	service, exists := objectsPerRequest.Load("services.UserService")

	if exists {
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
	objectsPerRequest.Store("services.UserService", userService)
	return userService, nil
}

func createCartService() (services.CartService, error) {
	service, exists := objectsPerRequest.Load("services.CartService")

	if exists {
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

	cartService := services.CreateCartService(cartRepository, productRepository)
	objectsPerRequest.Store("services.CartService", cartService)
	return cartService, nil
}

func createCategoryRepository() (repositories.CategoryRepository, error) {
	repo, exists := objectsPerRequest.Load("repositories.CategoryRepository")

	if exists {
		return repo.(repositories.CategoryRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	categoryRepository := repositories.CreateCategoryRepository(databaseConnection)
	objectsPerRequest.Store("repositories.CategoryRepository", categoryRepository)
	return categoryRepository, nil
}

func createProductRepository() (repositories.ProductRepository, error) {
	repo, exists := objectsPerRequest.Load("repositories.ProductRepository")

	if exists {
		return repo.(repositories.ProductRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	productRepository := repositories.CreateProductRepository(databaseConnection)
	objectsPerRequest.Store("repositories.ProductRepository", productRepository)
	return productRepository, nil
}

func createCartRepository() (repositories.CartRepository, error) {
	repo, exists := objectsPerRequest.Load("repositories.CartRepository")

	if exists {
		return repo.(repositories.CartRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	cartRepository := repositories.CreateCartRepository(databaseConnection)
	objectsPerRequest.Store("repositories.CartRepository", cartRepository)
	return cartRepository, nil
}

func createSessionRepository() (repositories.SessionRepository, error) {
	repo, exists := objectsPerRequest.Load("repositories.OrderRepository")

	if exists {
		return repo.(repositories.SessionRepository), nil
	}

	database, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	sessionRepository := repositories.CreateSessionRepository(database)
	objectsPerRequest.Store("repositories.SessionRepository", sessionRepository)
	return sessionRepository, nil
}

func createOrderRepository() (repositories.OrderRepository, error) {
	repo, exists := objectsPerRequest.Load("repositories.OrderRepository")

	if exists {
		return repo.(repositories.OrderRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	orderRepository := repositories.CreateOrderRepository(databaseConnection)
	objectsPerRequest.Store("repositories.OrderRepository", orderRepository)
	return orderRepository, nil
}

func createUserRepository() (repositories.UserRepository, error) {
	repo, exists := objectsPerRequest.Load("repositories.UserRepository")

	if exists {
		return repo.(repositories.UserRepository), nil
	}

	databaseConnection, err := CreateDatabaseConnection()
	if err != nil {
		return nil, err
	}

	userRepository := repositories.CreateUserRepository(databaseConnection)
	objectsPerRequest.Store("repositories.UserRepository", userRepository)
	return userRepository, nil
}

func addSessionProvider(sessionDto *dto.SessionDto) error {
	if sessionDto == nil {
		return errors.New("invalid session")
	}

	objectsPerRequest.Store("sessionProvider", sessionDto)
	return nil
}

func getSessionProvider() (*dto.SessionDto, error) {
	sessionProvider, exists := objectsPerRequest.Load("sessionProvider")

	if !exists {
		return nil, errors.New("'Session Provider' is nil, check if is added to 'objectCreator'")
	}

	return sessionProvider.(*dto.SessionDto), nil
}

func CreateHashId() (*hashids.HashID, error) {
	if hashId != nil {
		return hashId, nil
	}

	data := hashids.NewData()
	data.Salt = configuration.Server.IdSalt
	data.MinLength = 7
	hashIdLocal, err := hashids.NewWithData(data)
	if err != nil {
		return nil, err
	}

	hashId = hashIdLocal
	return hashIdLocal, nil
}
