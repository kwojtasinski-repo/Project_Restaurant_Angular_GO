package repositories

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/shopspring/decimal"
)

const timeStoreInCache = time.Minute * 5

func getTestProduct() entities.Product {
	id, _ := valueobjects.NewId(int64(rand.Intn(1000000-1) + 1))
	categoryName, _ := valueobjects.NewName("Category#" + uuid.NewString())
	product, _ := entities.NewProduct(int64(rand.Intn(1000000-1)+1), "Product"+uuid.NewString(), decimal.New(100, 0), "Description#123456789"+uuid.NewString(), &entities.Category{
		Id:   *id,
		Name: *categoryName,
	})
	return *product
}
