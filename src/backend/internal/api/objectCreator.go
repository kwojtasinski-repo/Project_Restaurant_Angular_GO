package api

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/services"
)

var inMemoryCategoryRepository = repositories.NewInMemoryCategoryRepository()
var inMemoryProductRepository = repositories.NewInMemoryProductRepository()
var inMemoryCartRepository = repositories.NewInMemoryCartRepository()
var inMemoryOrderRepository = repositories.NewInMemoryOrderRepository()
var inMemoryUserRepository = repositories.NewInMemoryUserRepository()
var passwordHasher = services.CreatePassworHasherService()

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
	return services.CreateUserService(inMemoryUserRepository, passwordHasher)
}

func createCartService() services.CartService {
	return services.CreateCartService(inMemoryCartRepository, inMemoryProductRepository, inMemoryUserRepository)
}
