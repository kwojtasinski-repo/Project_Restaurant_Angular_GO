package repositories

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

type CartRepository interface {
	Add(*entities.Cart) error
	Delete(entities.Cart) error
	Get(int64) (*entities.Cart, error)
	GetByUser(int64) ([]entities.Cart, error)
}

type inMemoryCartRepository struct {
	carts []entities.Cart
}

func NewInMemoryCartRepository() CartRepository {
	return &inMemoryCartRepository{
		carts: make([]entities.Cart, 0),
	}
}

func (repo *inMemoryCartRepository) Add(cart *entities.Cart) error {
	var length int = len(repo.carts)
	if length == 0 {
		cart.Id = 1
		repo.carts = append(repo.carts, *cart)
		return nil
	}

	lastElement := repo.carts[length-1]
	cart.Id = lastElement.Id + 1
	repo.carts = append(repo.carts, *cart)
	return nil
}

func (repo *inMemoryCartRepository) Delete(cartToDelete entities.Cart) error {
	for index, cart := range repo.carts {
		if cart.Id == cartToDelete.Id {
			repo.carts = append(repo.carts[:index], repo.carts[index+1:]...)
			return nil
		}
	}

	return nil
}

func (repo *inMemoryCartRepository) Get(id int64) (*entities.Cart, error) {
	for _, cart := range repo.carts {
		if cart.Id == id {
			return &cart, nil
		}
	}

	return nil, nil
}

func (repo *inMemoryCartRepository) GetByUser(userId int64) ([]entities.Cart, error) {
	carts := make([]entities.Cart, 0)

	for _, cart := range repo.carts {
		if cart.UserId == userId {
			carts = append(carts, cart)
		}
	}

	return carts, nil
}
