package repositories

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/shopspring/decimal"
)

func getTestProduct() entities.Product {
	product, _ := entities.NewProduct(int64(rand.Intn(1000000-1)+1), "Product"+uuid.NewString(), decimal.New(100, 0), "Description#123456789"+uuid.NewString(), &entities.Category{
		Id:   int64(rand.Intn(1000000-1) + 1),
		Name: "Category#" + uuid.NewString(),
	})
	return *product
}
