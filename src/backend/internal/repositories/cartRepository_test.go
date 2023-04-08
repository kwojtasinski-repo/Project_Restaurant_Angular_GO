package repositories

import (
	"testing"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
)

var cartRepository CartRepository = NewInMemoryCartRepository()

func addTestDataToCartRepo(repo CartRepository) {
	var testProduct = getTestProduct()
	var testCarts = [3]entities.Cart{
		{
			Id:        1,
			ProductId: testProduct.Id.Value(),
			Product:   testProduct,
			UserId:    1,
		},
		{
			Id:        2,
			ProductId: testProduct.Id.Value(),
			Product:   testProduct,
			UserId:    1,
		},
		{
			Id:        3,
			ProductId: testProduct.Id.Value(),
			Product:   testProduct,
			UserId:    1,
		},
	}
	for _, cart := range testCarts {
		repo.Add(&cart)
	}
}

func TestCartRepositoryAdd(t *testing.T) {
	var testProduct = getTestProduct()
	cart := &entities.Cart{
		Id:        0,
		ProductId: testProduct.Id.Value(),
		Product:   testProduct,
		UserId:    1,
	}

	cartRepository.Add(cart)

	cartAdded, err := cartRepository.Get(cart.Id)
	if cartAdded == nil {
		t.Fatalf(`'Cart' with id %v shouldnt be null`, cart.Id)
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func TestCartRepositoryGet(t *testing.T) {
	addTestDataToCartRepo(cartRepository)
	var id int64 = 2

	cart, err := cartRepository.Get(id)

	if cart == nil {
		t.Fatalf(`'Cart' with id %v shouldnt be null`, cart.Id)
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func TestCartRepositoryDelete(t *testing.T) {
	addTestDataToCartRepo(cartRepository)
	var cart, _ = cartRepository.Get(1)

	errDelete := cartRepository.Delete(*cart)

	cartDeleted, errGet := cartRepository.Get(cart.Id)
	if errDelete != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, errDelete)
	}
	if errGet != nil {
		t.Fatalf(`'Error' should be null`)
	}
	if cartDeleted != nil {
		t.Fatalf(`'Cart' with id %v should be null`, cart.Id)
	}
}
