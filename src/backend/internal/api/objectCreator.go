package api

import (
	"database/sql"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/config"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/services"
)

var inMemoryCategoryRepository = repositories.NewInMemoryCategoryRepository()
var inMemoryProductRepository = repositories.NewInMemoryProductRepository()
var inMemoryCartRepository = repositories.NewInMemoryCartRepository()
var inMemoryOrderRepository = repositories.NewInMemoryOrderRepository()
var inMemoryUserRepository = createInMemoryUserRepo()
var inMemorySessionRepository = repositories.NewInMemorySessionRepository()
var passwordHasher = services.CreatePassworHasherService()
var configuration config.Config
var objectsPerRequest = make(map[string]interface{})

func InitObjectCreator(configFile config.Config) {
	configuration = configFile
}

func CreateDatabase() (*sql.DB, error) {
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

func createSessionService() services.SessionService {
	service := objectsPerRequest["services.SessionService"]

	if service != nil {
		return service.(services.SessionService)
	}

	sessionService := services.CreateSessionService(inMemorySessionRepository, inMemoryUserRepository)
	objectsPerRequest["services.SessionService"] = sessionService
	return sessionService
}

func createProductService() services.ProductService {
	service := objectsPerRequest["services.ProductService"]

	if service != nil {
		return service.(services.ProductService)
	}

	productService := services.CreateProductService(inMemoryProductRepository, inMemoryCategoryRepository)
	objectsPerRequest["services.ProductService"] = productService
	return productService
}

func createCategoryService() services.CategoryService {
	service := objectsPerRequest["services.CategoryService"]

	if service != nil {
		return service.(services.CategoryService)
	}

	categoryService := services.CreateCategoryService(inMemoryCategoryRepository)
	objectsPerRequest["services.CategoryService"] = categoryService
	return categoryService
}

func createOrderService() services.OrderService {
	service := objectsPerRequest["services.OrderService"]

	if service != nil {
		return service.(services.OrderService)
	}

	orderService := services.CreateOrderService(inMemoryOrderRepository, inMemoryCartRepository, inMemoryProductRepository)
	objectsPerRequest["services.OrderService"] = orderService
	return orderService
}

func createUserService() services.UserService {
	service := objectsPerRequest["services.UserService"]

	if service != nil {
		return service.(services.UserService)
	}

	userService := services.CreateUserService(inMemoryUserRepository, passwordHasher, createSessionService())
	objectsPerRequest["services.UserService"] = userService
	return userService
}

func createCartService() services.CartService {
	service := objectsPerRequest["services.CartService"]

	if service != nil {
		return service.(services.CartService)
	}

	cartService := services.CreateCartService(inMemoryCartRepository, inMemoryProductRepository, inMemoryUserRepository)
	objectsPerRequest["services.CartService"] = cartService
	return cartService
}

func createPassworHasherService() services.PasswordHasherService {
	service := objectsPerRequest["services.PasswordHasherService"]

	if service != nil {
		return service.(services.PasswordHasherService)
	}

	passwordHasherService := services.CreatePassworHasherService()
	objectsPerRequest["services.PasswordHasherService"] = passwordHasherService
	return passwordHasherService
}

func createCategoryRepository() repositories.CategoryRepository {
	repo := objectsPerRequest["repositories.CategoryRepository"]

	if repo != nil {
		return repo.(repositories.CategoryRepository)
	}

	objectsPerRequest["repositories.CategoryRepository"] = nil
	return nil
}

func createProductRepository() repositories.ProductRepository {
	repo := objectsPerRequest["repositories.ProductRepository"]

	if repo != nil {
		return repo.(repositories.ProductRepository)
	}

	objectsPerRequest["repositories.ProductRepository"] = nil
	return nil
}

func createCartRepository() repositories.CartRepository {
	repo := objectsPerRequest["repositories.CartRepository"]

	if repo != nil {
		return repo.(repositories.CartRepository)
	}

	objectsPerRequest["repositories.CartRepository"] = nil
	return nil
}

func createSessionRepository() (repositories.SessionRepository, error) {
	repo := objectsPerRequest["repositories.OrderRepository"]

	if repo != nil {
		return repo.(repositories.SessionRepository), nil
	}

	database, err := CreateDatabase()
	if err != nil {
		return nil, err
	}

	sessionRepository := repositories.CreateSessionRepository(*database)
	objectsPerRequest["repositories.SessionRepository"] = sessionRepository
	return sessionRepository, nil
}

func createOrderRepository() repositories.OrderRepository {
	repo := objectsPerRequest["repositories.OrderRepository"]

	if repo != nil {
		return repo.(repositories.OrderRepository)
	}

	objectsPerRequest["repositories.OrderRepository"] = nil
	return nil
}

func createInUserRepository() repositories.UserRepository {
	repo := objectsPerRequest["repositories.UserRepository"]

	if repo != nil {
		return repo.(repositories.UserRepository)
	}

	objectsPerRequest["repositories.UserRepository"] = nil
	return nil
}

func createInMemoryUserRepo() repositories.UserRepository {
	repo := objectsPerRequest["repositories.UserRepository"]

	if repo != nil {
		return repo.(repositories.UserRepository)
	}

	passwordHasher := services.CreatePassworHasherService()
	password, _ := passwordHasher.HashPassword("admin")
	id, _ := valueobjects.NewId(1)
	email, _ := valueobjects.NewEmailAddress("admin@admin.com")
	user := &entities.User{
		Id:       *id,
		Email:    *email,
		Password: password,
		Role:     "admin",
	}
	userRepo := repositories.NewInMemoryUserRepository()
	userRepo.Add(user)
	objectsPerRequest["repositories.UserRepository"] = userRepo
	return userRepo
}
