package repositories

import (
	"testing"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
)

var cartRepository CartRepository = NewInMemoryCartRepository()

func addTestDataToCartRepo(repo CartRepository) {
	var testProduct = getTestProduct()
	id1, _ := valueobjects.NewId(1)
	id2, _ := valueobjects.NewId(2)
	id3, _ := valueobjects.NewId(3)
	var testCarts = [3]entities.Cart{
		{
			Id:        *id1,
			ProductId: testProduct.Id,
			Product:   testProduct,
			UserId:    *id1,
		},
		{
			Id:        *id2,
			ProductId: testProduct.Id,
			Product:   testProduct,
			UserId:    *id1,
		},
		{
			Id:        *id3,
			ProductId: testProduct.Id,
			Product:   testProduct,
			UserId:    *id1,
		},
	}
	for _, cart := range testCarts {
		repo.Add(&cart)
	}
}

func Test_CartRepository_Add(t *testing.T) {
	var testProduct = getTestProduct()
	id, _ := valueobjects.NewId(1)
	cart := &entities.Cart{
		Id:        *id,
		ProductId: testProduct.Id,
		Product:   testProduct,
		UserId:    *id,
	}

	cartRepository.Add(cart)

	cartAdded, err := cartRepository.Get(cart.Id.Value())
	if cartAdded == nil {
		t.Fatalf(`'Cart' with id %v shouldnt be null`, cart.Id.Value())
	}
	if err != nil {
		t.Fatalf(`'Error' should be null, and text contains %v`, err)
	}
}

func Test_CartRepository_Get(t *testing.T) {
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

func Test_CartRepository_Delete(t *testing.T) {
	addTestDataToCartRepo(cartRepository)
	var cart, _ = cartRepository.Get(1)

	errDelete := cartRepository.Delete(*cart)

	cartDeleted, errGet := cartRepository.Get(cart.Id.Value())
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
