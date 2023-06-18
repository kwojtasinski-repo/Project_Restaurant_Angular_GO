package repositories

import (
	"testing"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

var newCategoryRepository CategoryRepository = NewInMemoryCategoryRepository()

func addTestDataToCategoryRepo(repo CategoryRepository) {
	id1, _ := valueobjects.NewId(1)
	id2, _ := valueobjects.NewId(1)
	id3, _ := valueobjects.NewId(1)
	categoryName1, _ := valueobjects.NewName("Category #1")
	categoryName2, _ := valueobjects.NewName("Category #2")
	categoryName3, _ := valueobjects.NewName("Category #3")
	var testCategories = [3]entities.Category{
		{
			Id:   *id1,
			Name: *categoryName1,
		},
		{
			Id:   *id2,
			Name: *categoryName2,
		},
		{
			Id:   *id3,
			Name: *categoryName3,
		},
	}
	for _, category := range testCategories {
		repo.Add(&category)
	}
}

func Test_CategoryRepository_Add(t *testing.T) {
	id, _ := valueobjects.NewId(1)
	categoryName, _ := valueobjects.NewName("Category #1")
	var category = entities.Category{
		Id:   *id,
		Name: *categoryName,
	}

	newCategoryRepository.Add(&category)

	categoryAdded, err := newCategoryRepository.Get(category.Id.Value())
	if categoryAdded == nil {
		t.Fatalf(`'Category' with id %v shouldnt be null`, category.Id.Value())
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func Test_CategoryRepository_Get(t *testing.T) {
	addTestDataToCategoryRepo(newCategoryRepository)
	var id int64 = 2

	category, err := newCategoryRepository.Get(id)

	if category == nil {
		t.Fatalf(`'Category' with id %v shouldnt be null`, category.Id.Value())
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func Test_CategoryRepository_Delete(t *testing.T) {
	addTestDataToCategoryRepo(newCategoryRepository)
	var category, _ = newCategoryRepository.Get(1)

	errDelete := newCategoryRepository.Delete(*category)

	categoryDeleted, errGet := newCategoryRepository.Get(category.Id.Value())
	if errDelete != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, errDelete)
	}
	if errGet != nil {
		t.Fatalf(`'Error' should be null`)
	}
	if categoryDeleted == nil {
		t.Fatalf(`'Category' with id %v should not be null`, category.Id.Value())
	}
	if !categoryDeleted.Deleted {
		t.Fatalf(`'Category' with id %v should be deleted`, category.Id.Value())
	}
}

func Test_CategoryRepository_Update(t *testing.T) {
	addTestDataToCategoryRepo(newCategoryRepository)
	var category, _ = newCategoryRepository.Get(2)
	categoryName, _ := valueobjects.NewName("Abc1234Guid")
	category.Name = *categoryName

	newCategoryRepository.Update(*category)

	var productUpdated, err = newCategoryRepository.Get(category.Id.Value())
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
	if productUpdated == nil {
		t.Fatalf(`'Product' with id %v shouldnt be null`, category.Id.Value())
	}
	if category.Name != productUpdated.Name {
		t.Fatalf(`'Category' has different Name value expected '%v', actual '%v'`, category.Name, productUpdated.Name)
	}
}

func Test_CategoryRepository_GetAll(t *testing.T) {
	repo := NewInMemoryCategoryRepository()
	id1, _ := valueobjects.NewId(1)
	id2, _ := valueobjects.NewId(2)
	categoryName1, _ := valueobjects.NewName("Product #1")
	categoryName2, _ := valueobjects.NewName("Product #2")
	repo.Add(&entities.Category{
		Id:   *id1,
		Name: *categoryName1,
	})
	repo.Add(&entities.Category{
		Id:   *id2,
		Name: *categoryName2,
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
	if product.Id.Value() != 1 {
		t.Fatalf(`expected 'Category' with id '%v', actual %v`, 1, product.Id.Value())
	}
}
