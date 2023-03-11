package repositories

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/shopspring/decimal"
)

func getTestProduct() entities.Product {
	return entities.Product{
		Id:          1,
		Name:        "Product#1",
		Description: "Description#123456789",
		Price:       decimal.New(100, 0),
		Deleted:     false,
	}
}
