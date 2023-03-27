package repositories

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

type CategoryRepository interface {
	Add(*entities.Category) error
	Update(entities.Category) error
	Delete(entities.Category) error
	Get(int64) (*entities.Category, error)
	GetAll() ([]entities.Category, error)
}

type inMemoryCategoryRepository struct {
	categories []entities.Category
}

func NewInMemoryCategoryRepository() CategoryRepository {
	return &inMemoryCategoryRepository{
		categories: make([]entities.Category, 0),
	}
}

func (repo *inMemoryCategoryRepository) Add(category *entities.Category) error {
	var length int = len(repo.categories)
	if length == 0 {
		category.Id = 1
		repo.categories = append(repo.categories, *category)
		return nil
	}

	lastElement := repo.categories[length-1]
	category.Id = lastElement.Id + 1
	repo.categories = append(repo.categories, *category)
	return nil
}

func (repo *inMemoryCategoryRepository) Update(categoryToUpdate entities.Category) error {
	for index, category := range repo.categories {
		if category.Id == categoryToUpdate.Id {
			category.Name = categoryToUpdate.Name
			repo.categories[index] = category
		}
	}
	return nil
}

func (repo *inMemoryCategoryRepository) Delete(categoryToDelete entities.Category) error {
	for index, category := range repo.categories {
		if category.Id == categoryToDelete.Id {
			category.Deleted = true
			repo.categories[index] = category
			return nil
		}
	}

	return nil
}

func (repo *inMemoryCategoryRepository) Get(id int64) (*entities.Category, error) {
	for _, category := range repo.categories {
		if category.Id == id {
			return &category, nil
		}
	}

	return nil, nil
}

func (repo *inMemoryCategoryRepository) GetAll() ([]entities.Category, error) {
	categories := make([]entities.Category, 0)

	for _, category := range repo.categories {
		if !category.Deleted {
			categories = append(categories, category)
		}
	}

	return categories, nil
}
