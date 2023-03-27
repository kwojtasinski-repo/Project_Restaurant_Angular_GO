package api

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/services"
)

var inMemoryCategoryRepository = repositories.NewInMemoryCategoryRepository()
var inMemoryProductRepository = repositories.NewInMemoryProductRepository()
var inMemoryCartRepository = repositories.NewInMemoryCartRepository()
var inMemoryOrderRepository = repositories.NewInMemoryOrderRepository()

func createProductService() services.ProductService {
	return services.CreateProductService(inMemoryProductRepository, inMemoryCategoryRepository)
}

func createCategoryService() services.CategoryService {
	return services.CreateCategoryService(inMemoryCategoryRepository)
}

func createOrderService() services.OrderService {
	return services.CreateOrderService(inMemoryOrderRepository, inMemoryCartRepository, inMemoryProductRepository)
}
