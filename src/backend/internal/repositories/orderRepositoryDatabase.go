package repositories

import (
	"database/sql"
	"time"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/shopspring/decimal"
)

type orderRepository struct {
	database *sql.DB
}

func CreateOrderRepository(database *sql.DB) OrderRepository {
	return &orderRepository{
		database: database,
	}
}

func (repo *orderRepository) Add(order *entities.Order) error {
	query := "INSERT INTO `orders` (order_number, price, created, modified, user_id) VALUES (?, ?, ?, ?, ?);"
	result, err := repo.database.Exec(query, order.OrderNumber.Value(), order.Price.Value(), order.Created, order.Modified, order.UserId.Value())
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	newId, _ := valueobjects.NewId(id)
	order.Id = *newId

	for _, orderProduct := range order.OrderProducts {
		queryOrderProduct := "INSERT INTO `order_products` (name, price, product_id, order_id) VALUES (?, ?, ?, ?);"
		result, err := repo.database.Exec(queryOrderProduct, orderProduct.Name.Value(), orderProduct.Price.Value(), orderProduct.ProductId.Value(), id)
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		newId, _ := valueobjects.NewId(id)
		orderProduct.Id = *newId
	}

	return nil
}

func (repo *orderRepository) Delete(orderToDelete entities.Order) error {
	query := "DELETE FROM `orders` WHERE id = ?;"
	_, err := repo.database.Exec(query, orderToDelete.Id.Value())
	if err != nil {
		return err
	}

	return nil
}

func (repo *orderRepository) Update(orderToUpdate *entities.Order) error {
	query := "UPDATE `orders` SET order_number = ?, price = ?, SET modified = ? WHERE id = ?;"
	_, err := repo.database.Exec(query, orderToUpdate.OrderNumber.Value(), orderToUpdate.Price.Value(), orderToUpdate.Modified, orderToUpdate.Id.Value())
	if err != nil {
		return err
	}
	return nil
}

func (repo *orderRepository) Get(id int64) (*entities.Order, error) {
	query := `SELECT o.id, o.order_number, o.price, o.created, o.modified, o.user_id, 
					 op.id, op.name, op.price, op.product_id, 
					 u.email 
			  FROM orders o 
			  INNER JOIN order_products op ON op.order_id = o.id 
			  LEFT JOIN users u ON u.id = o.user_id 
			  WHERE o.id = ?;`

	rows, err := repo.database.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderId int64
	var orderNumber string
	var price decimal.Decimal
	var created time.Time
	var modified *time.Time
	var userId int64
	var newUserId *valueobjects.Id
	var newEmail *valueobjects.EmailAddress
	orderProducts := make([]entities.OrderProduct, 0)
	for rows.Next() {
		var orderProductId int64
		var productName string
		var productPrice decimal.Decimal
		var productId int64
		var email string
		if err := rows.Scan(&orderId, &orderNumber, &price, &created, &modified, &userId,
			&orderProductId, &productName, &productPrice, &productId, &email); err != nil {
			return nil, err
		}

		newUserId, _ = valueobjects.NewId(userId)
		newEmail, _ = valueobjects.NewEmailAddress(email)

		newId, _ := valueobjects.NewId(orderProductId)
		newProductName, _ := valueobjects.NewName(productName)
		newProductPrice, _ := valueobjects.NewPrice(productPrice)
		newProductId, _ := valueobjects.NewId(productId)
		orderProduct := entities.OrderProduct{
			Id:        *newId,
			Name:      *newProductName,
			Price:     *newProductPrice,
			ProductId: *newProductId,
		}
		orderProducts = append(orderProducts, orderProduct)
	}

	user := entities.User{
		Id:    *newUserId,
		Email: *newEmail,
	}
	newId, _ := valueobjects.NewId(orderId)
	newPrice, _ := valueobjects.NewPrice(price)
	order := &entities.Order{
		Id:            *newId,
		OrderNumber:   valueobjects.NewOrderNumberWithGiven(orderNumber),
		Price:         *newPrice,
		Created:       created,
		Modified:      modified,
		User:          user,
		UserId:        user.Id,
		OrderProducts: orderProducts,
	}

	return order, nil
}

func (repo *orderRepository) GetAllByUser(userId int64) ([]entities.Order, error) {
	orders := make([]entities.Order, 0)
	query := "SELECT id, order_number, price, created, modified, user_id FROM `orders` WHERE user_id = ?;"
	rows, err := repo.database.Query(query, userId)
	if err != nil {
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var orderNumber string
		var price decimal.Decimal
		var created time.Time
		var modified *time.Time
		var userId int64
		if err := rows.Scan(&id, &orderNumber, &price, &created, &modified, &userId); err != nil {
			return nil, err
		}

		newId, _ := valueobjects.NewId(id)
		newPrice, _ := valueobjects.NewPrice(price)
		newUserId, _ := valueobjects.NewId(userId)
		order := entities.Order{
			Id:          *newId,
			OrderNumber: valueobjects.NewOrderNumberWithGiven(orderNumber),
			Price:       *newPrice,
			Created:     created,
			Modified:    modified,
			UserId:      *newUserId,
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (repo *orderRepository) GetAll() ([]entities.Order, error) {
	orders := make([]entities.Order, 0)
	query := "SELECT id, order_number, price, created, modified, user_id FROM `orders`;"
	rows, err := repo.database.Query(query)
	if err != nil {
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var orderNumber string
		var price decimal.Decimal
		var created time.Time
		var modified *time.Time
		var userId int64
		if err := rows.Scan(&id, &orderNumber, &price, &created, &modified, &userId); err != nil {
			return nil, err
		}

		newId, _ := valueobjects.NewId(id)
		newPrice, _ := valueobjects.NewPrice(price)
		newUserId, _ := valueobjects.NewId(userId)
		order := entities.Order{
			Id:          *newId,
			OrderNumber: valueobjects.NewOrderNumberWithGiven(orderNumber),
			Price:       *newPrice,
			Created:     created,
			Modified:    modified,
			UserId:      *newUserId,
		}
		orders = append(orders, order)
	}

	return orders, nil
}
