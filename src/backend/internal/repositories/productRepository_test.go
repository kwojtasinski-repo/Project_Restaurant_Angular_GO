package repositories

import (
	"testing"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

var newProductRepository ProductRepository = NewInMemoryProductRepository()

func addTestDataToProductRepo(repo ProductRepository) {
	var testProducts = [3]entities.Product{
		getTestProduct(),
		getTestProduct(),
		getTestProduct(),
	}
	for _, product := range testProducts {
		repo.Add(&product)
	}
}

func Test_ProductRepository_Add(t *testing.T) {
	var testProduct = getTestProduct()
	newProductRepository.Add(&testProduct)

	productAdded, err := newProductRepository.Get(testProduct.Id.Value())
	if productAdded == nil {
		t.Fatalf(`'Product' with id %v shouldnt be null`, testProduct.Id.Value())
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func Test_ProductRepository_Get(t *testing.T) {
	addTestDataToProductRepo(newProductRepository)
	var id int64 = 2

	product, err := newProductRepository.Get(id)

	if product == nil {
		t.Fatalf(`'Product' with id %v shouldnt be null`, product.Id.Value())
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func Test_ProductRepository_Delete(t *testing.T) {
	addTestDataToProductRepo(newProductRepository)
	var product, _ = newProductRepository.Get(1)

	errDelete := newProductRepository.Delete(product)

	productDeleted, errGet := newProductRepository.Get(product.Id.Value())
	if errDelete != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, errDelete)
	}
	if errGet != nil {
		t.Fatalf(`'Error' should be null`)
	}
	if productDeleted == nil {
		t.Fatalf(`'Product' with id %v should not be null`, product.Id.Value())
	}
	if !productDeleted.Deleted {
		t.Fatalf(`'Product' with id %v should be deleted`, product.Id.Value())
	}
}

func Test_ProductRepository_Update(t *testing.T) {
	addTestDataToProductRepo(newProductRepository)
	var product, _ = newProductRepository.Get(2)
	product.Deleted = true
	var description *valueobjects.Description
	description, _ = valueobjects.NewDescription("Test123456789")
	product.Description = *description

	newProductRepository.Update(product)

	var productUpdated, err = newProductRepository.Get(product.Id.Value())
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
	if productUpdated == nil {
		t.Fatalf(`'Product' with id %v shouldnt be null`, product.Id.Value())
	}
	if product.Deleted != productUpdated.Deleted {
		t.Fatalf(`'Product' has different Deleted value expected '%v', actual '%v'`, product.Deleted, productUpdated.Deleted)
	}
	if product.Description.Value() != productUpdated.Description.Value() {
		t.Fatalf(`'Product' has different Description value expected '%v', actual '%v'`, product.Description.Value(), productUpdated.Description.Value())
	}
}

func Test_ProductRepository_GetAll(t *testing.T) {
	repo := NewInMemoryProductRepository()
	product1 := getTestProduct()
	product2 := getTestProduct()
	product2.Deleted = true
	repo.Add(&product1)
	repo.Add(&product2)

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
	if product.Id.Value() != 1 {
		t.Fatalf(`expected 'Product' with id '%v', actual %v`, 1, product.Id.Value())
	}
}
