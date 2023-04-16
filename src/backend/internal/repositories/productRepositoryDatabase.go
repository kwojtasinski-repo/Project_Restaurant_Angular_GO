package repositories

import (
	"database/sql"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/shopspring/decimal"
)

type productRepository struct {
	database sql.DB
}

func CreateProductRepository(database sql.DB) ProductRepository {
	return &productRepository{
		database: database,
	}
}

func (repo *productRepository) Add(product *entities.Product) error {
	query := "INSERT INTO `products` (name, description, price, category_id, deleted) VALUES (?, ?, ?, ?, ?);"
	result, err := repo.database.Exec(query, product.Name.Value(), product.Description.Value(), product.Price.Value(), product.Category.Id.Value(), product.Deleted)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	newId, _ := valueobjects.NewId(id)
	product.Id = *newId
	return nil
}

func (repo *productRepository) Update(productToUpdate entities.Product) error {
	query := "UPDATE `products` SET name = ?, description = ?, price = ?, category_id = ?, deleted = ? WHERE id = ?;"
	_, err := repo.database.Exec(query, productToUpdate.Name.Value(), productToUpdate.Description.Value(), productToUpdate.Price.Value(), productToUpdate.Category.Id.Value(), productToUpdate.Deleted, productToUpdate.Id.Value())
	if err != nil {
		return err
	}
	return nil
}

func (repo *productRepository) Delete(productToDelete entities.Product) error {
	query := "UPDATE `products` SET deleted = ? WHERE id = ?;"
	_, err := repo.database.Exec(query, true, productToDelete.Id.Value())
	if err != nil {
		return err
	}
	return nil
}

func (repo *productRepository) Get(id int64) (*entities.Product, error) {
	query := `SELECT p.id, p.name, p.description, p.price, p.category_id, p.deleted, 
					 c.id, c.name, c.deleted 
			  FROM products p 
			  INNER JOIN categories c ON c.id = p.category_id 
			  WHERE p.id = ?;`
	row := repo.database.QueryRow(query, id)
	var productId int64
	var productName string
	var productDescription string
	var productPrice decimal.Decimal
	var productDeleted bool
	var categoryId int64
	var categoryName string
	var categoryDeleted bool
	if err := row.Scan(&productId, &productName, &productDescription, &productPrice, &categoryId, &productDeleted, &categoryId, &categoryName, &categoryDeleted); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	newCategoryId, _ := valueobjects.NewId(categoryId)
	newCategoryName, _ := valueobjects.NewName(categoryName)
	category := &entities.Category{
		Id:      *newCategoryId,
		Name:    *newCategoryName,
		Deleted: categoryDeleted,
	}

	newProductId, _ := valueobjects.NewId(productId)
	newProductName, _ := valueobjects.NewName(productName)
	newProductDescription, _ := valueobjects.NewDescription(productDescription)
	newProductPrice, _ := valueobjects.NewPrice(productPrice)
	product := &entities.Product{
		Id:          *newProductId,
		Name:        *newProductName,
		Description: *newProductDescription,
		Price:       *newProductPrice,
		Deleted:     productDeleted,
		Category:    *category,
	}
	return product, nil
}

func (repo *productRepository) GetAll() ([]entities.Product, error) {
	products := make([]entities.Product, 0)
	query := `SELECT p.id, p.name, p.description, p.price, p.category_id
			  FROM products p
			  WHERE p.deleted = false;`
	rows, err := repo.database.Query(query)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name string
		var description string
		var price decimal.Decimal
		var categoryId int64
		if err := rows.Scan(&id, &name, &description, &price, &categoryId); err != nil {
			return nil, err
		}

		newCategoryId, _ := valueobjects.NewId(categoryId)
		category := &entities.Category{
			Id: *newCategoryId,
		}
		product, _ := entities.NewProduct(id, name, price, description, category)
		products = append(products, *product)
	}

	return products, nil
}
