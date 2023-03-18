package repositories

import (
	"testing"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

var categoryRepository CategoryRepository = NewInMemoryCategoryRepository()

func addTestDataToCategoryRepo(repo CategoryRepository) {
	var testCategories = [3]entities.Category{
		{
			Id:   1,
			Name: "Category #1",
		},
		{
			Id:   2,
			Name: "Category #2",
		},
		{
			Id:   3,
			Name: "Category #3",
		},
	}
	for _, category := range testCategories {
		repo.Add(&category)
	}
}

func TestCategoryRepositoryAdd(t *testing.T) {
	var category = entities.Category{
		Id:   0,
		Name: "Category #1",
	}
	categoryRepository.Add(&category)

	categoryAdded, err := categoryRepository.Get(category.Id)
	if categoryAdded == nil {
		t.Fatalf(`'Category' with id %v shouldnt be null`, category.Id)
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func TestCategoryRepositoryGet(t *testing.T) {
	addTestDataToCategoryRepo(categoryRepository)
	var id int64 = 2

	category, err := categoryRepository.Get(id)

	if category == nil {
		t.Fatalf(`'Category' with id %v shouldnt be null`, category.Id)
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func TestCategoryRepositoryDelete(t *testing.T) {
	addTestDataToCategoryRepo(categoryRepository)
	var category, _ = categoryRepository.Get(1)

	errDelete := categoryRepository.Delete(*category)

	categoryDeleted, errGet := categoryRepository.Get(category.Id)
	if errDelete != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, errDelete)
	}
	if errGet != nil {
		t.Fatalf(`'Error' should be null`)
	}
	if categoryDeleted != nil {
		t.Fatalf(`'Category' with id %v should be null`, category.Id)
	}
}

func TestCategoryRepositoryUpdate(t *testing.T) {
	addTestDataToCategoryRepo(categoryRepository)
	var category, _ = categoryRepository.Get(2)
	category.Name = "Abc1234Guid"

	categoryRepository.Update(*category)

	var productUpdated, err = categoryRepository.Get(category.Id)
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
	if productUpdated == nil {
		t.Fatalf(`'Product' with id %v shouldnt be null`, category.Id)
	}
	if category.Name != productUpdated.Name {
		t.Fatalf(`'Category' has different Name value expected '%v', actual '%v'`, category.Name, productUpdated.Name)
	}
}

func TestCategoryRepositoryGetAll(t *testing.T) {
	repo := NewInMemoryCategoryRepository()
	repo.Add(&entities.Category{
		Id:   1,
		Name: "Product #1",
	})
	repo.Add(&entities.Category{
		Id:   2,
		Name: "Product #2",
	})

	products, err := repo.GetAll()

	if len(products) == 0 {
		t.Fatalf(`'Categories' has no elements'`)
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
	if len(products) != 2 {
		t.Fatalf(`'Categories' should have only two elements`)
	}
	var product = products[0]
	if product.Id != 1 {
		t.Fatalf(`expected 'Category' with id '%v', actual %v`, 1, product.Id)
	}
}
