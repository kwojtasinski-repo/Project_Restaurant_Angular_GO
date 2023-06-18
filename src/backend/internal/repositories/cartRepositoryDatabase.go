package repositories

import (
	"database/sql"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/shopspring/decimal"
)

type cartRepository struct {
	database *sql.DB
}

func CreateCartRepository(database *sql.DB) CartRepository {
	return &cartRepository{
		database: database,
	}
}

func (repo *cartRepository) Add(cart *entities.Cart) error {
	query := "INSERT INTO `carts` (product_id, user_id) VALUES (?, ?);"
	result, err := repo.database.Exec(query, cart.ProductId.Value(), cart.UserId.Value())
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	newId, _ := valueobjects.NewId(id)
	cart.Id = *newId
	return nil
}

func (repo *cartRepository) Delete(cartToDelete entities.Cart) error {
	query := "DELETE FROM `carts` WHERE id = ?;"
	_, err := repo.database.Exec(query, cartToDelete.Id.Value())
	if err != nil {
		return err
	}
	return nil
}

func (repo *cartRepository) Get(id int64) (*entities.Cart, error) {
	// probably join will be necessary
	query := `SELECT c.id, c.product_id, c.user_id 
			  FROM carts c 
			  WHERE id = ?;`
	row := repo.database.QueryRow(query, id)
	var cartId int64
	var productId int64
	var userId int64
	if err := row.Scan(&cartId, &productId, &userId); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	newId, _ := valueobjects.NewId(cartId)
	newProductId, _ := valueobjects.NewId(productId)
	newUserId, _ := valueobjects.NewId(userId)
	cart := &entities.Cart{
		Id:        *newId,
		ProductId: *newProductId,
		UserId:    *newUserId,
	}
	return cart, nil
}

func (repo *cartRepository) GetAllByUser(userId int64) ([]entities.Cart, error) {
	carts := make([]entities.Cart, 0)
	query := `SELECT c.id, c.product_id, c.user_id,
					 p.id, p.name, p.description, p.price 
			  FROM carts c 
			  JOIN products p ON p.id = c.product_id 
			  WHERE c.user_id = ?`
	rows, err := repo.database.Query(query, userId)
	if err != nil {
		return carts, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartId int64
		var productId int64
		var userId int64
		var name string
		var description string
		var price decimal.Decimal
		if err := rows.Scan(&cartId, &productId, &userId, &productId, &name, &description, &price); err != nil {
			return nil, err
		}

		newProductId, _ := valueobjects.NewId(productId)
		newName, _ := valueobjects.NewName(name)
		newDescription, _ := valueobjects.NewDescription(description)
		newPrice, _ := valueobjects.NewPrice(price)
		product := entities.Product{
			Id:          *newProductId,
			Name:        *newName,
			Description: *newDescription,
			Price:       *newPrice,
		}
		newCartId, _ := valueobjects.NewId(cartId)
		newUserId, _ := valueobjects.NewId(userId)
		cart := entities.Cart{
			Id:        *newCartId,
			ProductId: *newProductId,
			UserId:    *newUserId,
			Product:   product,
		}
		carts = append(carts, cart)
	}

	return carts, nil
}

func (repo *cartRepository) DeleteCartByUserId(userId int64) error {
	query := "DELETE FROM `carts` WHERE user_id = ?;"
	_, err := repo.database.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}
