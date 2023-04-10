package repositories

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

type OderRepository interface {
	Add(*entities.Order) error
	Delete(entities.Order) error
	Get(int64) (*entities.Order, error)
	GetAllByUser(int64) ([]entities.Order, error)
	Update(*entities.Order) error
}

type inMemoryOrderRepository struct {
	orders []entities.Order
}

func NewInMemoryOrderRepository() OderRepository {
	return &inMemoryOrderRepository{
		orders: make([]entities.Order, 0),
	}
}

func (repo *inMemoryOrderRepository) Add(order *entities.Order) error {
	var length int = len(repo.orders)
	if length == 0 {
		id, _ := valueobjects.NewId(1)
		order.Id = *id
		repo.orders = append(repo.orders, *order)
		return nil
	}

	lastElement := repo.orders[length-1]
	id, _ := valueobjects.NewId(lastElement.Id.Value() + 1)
	order.Id = *id
	repo.orders = append(repo.orders, *order)
	return nil
}

func (repo *inMemoryOrderRepository) Delete(orderToDelete entities.Order) error {
	for index, order := range repo.orders {
		if order.Id == orderToDelete.Id {
			repo.orders = append(repo.orders[:index], repo.orders[index+1:]...)
			return nil
		}
	}

	return nil
}

func (repo *inMemoryOrderRepository) Update(orderToUpdate *entities.Order) error {
	for index, order := range repo.orders {
		if order.Id == orderToUpdate.Id {
			order.OrderNumber = orderToUpdate.OrderNumber
			order.Modified = orderToUpdate.Modified
			order.UserId = orderToUpdate.UserId
			order.Price = orderToUpdate.Price
			order.User = orderToUpdate.User
			order.OrderProducts = orderToUpdate.OrderProducts
			repo.orders[index] = order
		}
	}
	return nil
}

func (repo *inMemoryOrderRepository) Get(id int64) (*entities.Order, error) {
	for _, order := range repo.orders {
		if order.Id.Value() == id {
			return &order, nil
		}
	}
	return nil, nil
}

func (repo *inMemoryOrderRepository) GetAllByUser(userId int64) ([]entities.Order, error) {
	orders := make([]entities.Order, 0)

	for _, order := range repo.orders {
		if order.UserId.Value() == userId {
			orders = append(orders, order)
		}
	}
	return orders, nil
}
