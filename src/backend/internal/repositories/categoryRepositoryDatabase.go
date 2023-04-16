package repositories

import (
	"database/sql"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type categoryRepository struct {
	database sql.DB
}

func CreateCategoryRepository(database sql.DB) CategoryRepository {
	return &categoryRepository{
		database: database,
	}
}

func (repo *categoryRepository) Add(category *entities.Category) error {
	query := "INSERT INTO `categories` (name, deleted) VALUES (?, ?);"
	result, err := repo.database.Exec(query, category.Name.Value(), category.Deleted)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	newId, _ := valueobjects.NewId(id)
	category.Id = *newId
	return nil
}

func (repo *categoryRepository) Update(categoryToUpdate entities.Category) error {
	query := "UPDATE `categories` SET name = ?, deleted = ? WHERE id = ?;"
	_, err := repo.database.Exec(query, categoryToUpdate.Name.Value(), categoryToUpdate.Deleted, categoryToUpdate.Id.Value())
	if err != nil {
		return err
	}
	return nil
}

func (repo *categoryRepository) Delete(categoryToDelete entities.Category) error {
	query := "UPDATE `categories` SET deleted = ? WHERE id = ?;"
	_, err := repo.database.Exec(query, true, categoryToDelete.Id.Value())
	if err != nil {
		return err
	}
	return nil
}

func (repo *categoryRepository) Get(id int64) (*entities.Category, error) {
	query := "SELECT id, name, deleted FROM `categories` WHERE id = ?;"
	row := repo.database.QueryRow(query, id)
	var categoryId int64
	var name string
	var deleted bool
	if err := row.Scan(&categoryId, &name, &deleted); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	newId, _ := valueobjects.NewId(categoryId)
	newName, _ := valueobjects.NewName(name)
	category := &entities.Category{
		Id:      *newId,
		Name:    *newName,
		Deleted: deleted,
	}
	return category, nil
}

func (repo *categoryRepository) GetAll() ([]entities.Category, error) {
	categories := make([]entities.Category, 0)

	query := "SELECT id, name FROM `categories` WHERE deleted = false;"
	rows, err := repo.database.Query(query)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var categoryId int64
		var name string
		if err := rows.Scan(&categoryId, &name); err != nil {
			return nil, err
		}

		category, _ := entities.NewCategory(categoryId, name)
		categories = append(categories, *category)
	}

	return categories, nil
}
