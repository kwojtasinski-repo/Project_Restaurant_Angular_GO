package repositories

import (
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type ProductRepository interface {
	Add(*entities.Product) error
	Update(entities.Product) error
	Delete(entities.Product) error
	Get(int64) (*entities.Product, error)
	GetAll() ([]entities.Product, error)
}

type inMemoryProductRepository struct {
	products []entities.Product
}

func NewInMemoryProductRepository() ProductRepository {
	return &inMemoryProductRepository{
		products: make([]entities.Product, 0),
	}
}

func (repo *inMemoryProductRepository) Add(product *entities.Product) error {
	var length int = len(repo.products)
	if length == 0 {
		newId, _ := valueobjects.NewId(1)
		product.Id = *newId
		repo.products = append(repo.products, *product)
		return nil
	}

	lastElement := repo.products[length-1]
	newId, _ := valueobjects.NewId(lastElement.Id.Value() + 1)
	product.Id = *newId
	repo.products = append(repo.products, *product)
	return nil
}

func (repo *inMemoryProductRepository) Update(productToUpdate entities.Product) error {
	for index, product := range repo.products {
		if product.Id.Value() == productToUpdate.Id.Value() {
			product.Name = productToUpdate.Name
			product.Description = productToUpdate.Description
			product.Price = productToUpdate.Price
			product.Deleted = productToUpdate.Deleted
			product.Category = productToUpdate.Category
			repo.products[index] = product
		}
	}
	return nil
}

func (repo *inMemoryProductRepository) Delete(productToDelete entities.Product) error {
	for index, product := range repo.products {
		if product.Id.Value() == productToDelete.Id.Value() {
			product.Deleted = true
			repo.products[index] = product
			return nil
		}
	}

	return nil
}

func (repo *inMemoryProductRepository) Get(id int64) (*entities.Product, error) {
	for _, product := range repo.products {
		if product.Id.Value() == id {
			return &product, nil
		}
	}

	return nil, nil
}

func (repo *inMemoryProductRepository) GetAll() ([]entities.Product, error) {
	products := make([]entities.Product, 0)

	for _, product := range repo.products {
		if !product.Deleted {
			products = append(products, product)
		}
	}

	return products, nil
}
