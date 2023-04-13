package api

import (
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
var passwordHasher = services.CreatePassworHasherService()
var inMemorySessionRepository = repositories.NewInMemorySessionRepository()

func createInMemoryUserRepo() repositories.UserRepository {
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
	return userRepo
}

func createProductService() services.ProductService {
	return services.CreateProductService(inMemoryProductRepository, inMemoryCategoryRepository)
}

func createCategoryService() services.CategoryService {
	return services.CreateCategoryService(inMemoryCategoryRepository)
}

func createOrderService() services.OrderService {
	return services.CreateOrderService(inMemoryOrderRepository, inMemoryCartRepository, inMemoryProductRepository)
}

func createUserService() services.UserService {
	return services.CreateUserService(inMemoryUserRepository, passwordHasher, CreateSessionService())
}

func createCartService() services.CartService {
	return services.CreateCartService(inMemoryCartRepository, inMemoryProductRepository, inMemoryUserRepository)
}

func CreateSessionService() services.SessionService {
	return services.CreateSessionService(inMemorySessionRepository, inMemoryUserRepository)
}
