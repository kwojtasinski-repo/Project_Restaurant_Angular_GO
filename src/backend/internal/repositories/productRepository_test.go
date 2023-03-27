package repositories

import (
	"testing"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/shopspring/decimal"
)

var productRepository ProductRepository = NewInMemoryProductRepository()

func addTestDataToProductRepo(repo ProductRepository) {
	var testProducts = [3]entities.Product{
		getTestProduct(),
		{
			Id:          2,
			Name:        "Product #2",
			Description: "Description",
			Price:       decimal.New(15555, -2),
			Category: entities.Category{
				Id:   1,
				Name: "Category#1",
			},
			Deleted: false,
		},
		{
			Id:          3,
			Name:        "Product #3",
			Description: "Description3",
			Price:       decimal.New(1555555, -4),
			Category: entities.Category{
				Id:   1,
				Name: "Category#1",
			},
			Deleted: false,
		},
	}
	for _, product := range testProducts {
		repo.Add(&product)
	}
}

func TestProductRepositoryAdd(t *testing.T) {
	var testProduct = getTestProduct()
	productRepository.Add(&testProduct)

	productAdded, err := productRepository.Get(testProduct.Id)
	if productAdded == nil {
		t.Fatalf(`'Product' with id %v shouldnt be null`, testProduct.Id)
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func TestProductRepositoryGet(t *testing.T) {
	addTestDataToProductRepo(productRepository)
	var id int64 = 2

	product, err := productRepository.Get(id)

	if product == nil {
		t.Fatalf(`'Product' with id %v shouldnt be null`, product.Id)
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func TestProductRepositoryDelete(t *testing.T) {
	addTestDataToProductRepo(productRepository)
	var product, _ = productRepository.Get(1)

	errDelete := productRepository.Delete(*product)

	productDeleted, errGet := productRepository.Get(product.Id)
	if errDelete != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, errDelete)
	}
	if errGet != nil {
		t.Fatalf(`'Error' should be null`)
	}
	if productDeleted == nil {
		t.Fatalf(`'Product' with id %v should not be null`, product.Id)
	}
	if !productDeleted.Deleted {
		t.Fatalf(`'Product' with id %v should be deleted`, product.Id)
	}
}

func TestProductRepositoryUpdate(t *testing.T) {
	addTestDataToProductRepo(productRepository)
	var product, _ = productRepository.Get(2)
	product.Deleted = true
	product.Description = "Test123456789"

	productRepository.Update(*product)

	var productUpdated, err = productRepository.Get(product.Id)
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
	if productUpdated == nil {
		t.Fatalf(`'Product' with id %v shouldnt be null`, product.Id)
	}
	if product.Deleted != productUpdated.Deleted {
		t.Fatalf(`'Product' has different Deleted value expected '%v', actual '%v'`, product.Deleted, productUpdated.Deleted)
	}
	if product.Description != productUpdated.Description {
		t.Fatalf(`'Product' has different Description value expected '%v', actual '%v'`, product.Description, productUpdated.Description)
	}
}

func TestProductRepositoryGetAll(t *testing.T) {
	repo := NewInMemoryProductRepository()
	repo.Add(&entities.Product{
		Id:          1,
		Name:        "Product #2",
		Description: "Description",
		Price:       decimal.New(15555, -2),
		Category: entities.Category{
			Id:   1,
			Name: "Category#1",
		},
		Deleted: false,
	})
	repo.Add(&entities.Product{
		Id:          2,
		Name:        "Product #3",
		Description: "Description3",
		Price:       decimal.New(1555555, -4),
		Category: entities.Category{
			Id:   1,
			Name: "Category#1",
		},
		Deleted: true,
	})

	products, err := repo.GetAll()

	if len(products) == 0 {
		t.Fatalf(`'Products' has no elements'`)
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
	if len(products) > 1 {
		t.Fatalf(`'Products' should have only one element`)
	}
	var product = products[0]
	if product.Id != 1 {
		t.Fatalf(`expected 'Product' with id '%v', actual %v`, 1, product.Id)
	}
}
