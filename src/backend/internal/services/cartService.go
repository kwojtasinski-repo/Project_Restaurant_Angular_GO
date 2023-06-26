package services

import (
	"fmt"
	"log"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
)

type CartService interface {
	AddToCart(addCart dto.AddCart) *applicationerrors.ErrorStatus
	RemoveFromCart(cartId, userId int64) *applicationerrors.ErrorStatus
	GetMyCart(userId int64) ([]dto.CartDto, *applicationerrors.ErrorStatus)
}

type cartService struct {
	repo              repositories.CartRepository
	productRepository repositories.ProductRepository
}

func CreateCartService(cartRepository repositories.CartRepository, productRepository repositories.ProductRepository) CartService {
	return &cartService{
		repo:              cartRepository,
		productRepository: productRepository,
	}
}

func (service *cartService) AddToCart(addCart dto.AddCart) *applicationerrors.ErrorStatus {
	product, errorStatus := getProduct(service.productRepository, addCart.ProductId.ValueInt)
	if errorStatus != nil {
		return errorStatus
	}

	userId, err := valueobjects.NewId(addCart.UserId.ValueInt)
	if err != nil {
		return applicationerrors.BadRequest(applicationerrors.InvalidUserId)
	}

	cart, err := entities.NewCart(0, entities.User{Id: *userId}, *product)
	if err != nil {
		return applicationerrors.BadRequest(err.Error())
	}

	log.Printf("'User' with id '%v' add 'Product' with id '%v'\n", userId.Value(), product.Id)
	err = service.repo.Add(cart)
	if err != nil {
		return applicationerrors.InternalError(err.Error())
	}

	return nil
}

func (service *cartService) RemoveFromCart(cartId, userId int64) *applicationerrors.ErrorStatus {
	cart, errorStatus := getCart(service.repo, cartId)
	if errorStatus != nil {
		return errorStatus
	}

	log.Printf("'User' with id '%v' removes 'ItemCart' with id '%v'\n", userId, cart.Id)
	err := service.repo.Delete(*cart)
	if err != nil {
		return applicationerrors.InternalError(err.Error())
	}

	return nil
}

func (service *cartService) GetMyCart(userId int64) ([]dto.CartDto, *applicationerrors.ErrorStatus) {
	carts := make([]dto.CartDto, 0)
	log.Printf("'User' with id '%v' get all items from cart\n", userId)
	cartsInRepo, err := service.repo.GetAllByUser(userId)
	if err != nil {
		return carts, applicationerrors.InternalError(err.Error())
	}

	for _, cart := range cartsInRepo {
		carts = append(carts, dto.CartDto{
			Id:      dto.IdObject{ValueInt: cart.Id.Value()},
			Product: *dto.MapToProductDto(cart.Product),
			UserId:  dto.IdObject{ValueInt: userId},
		})
	}

	return carts, nil
}

func getProduct(productRepo repositories.ProductRepository, productId int64) (*entities.Product, *applicationerrors.ErrorStatus) {
	var product *entities.Product
	var err error
	product, err = productRepo.Get(productId)

	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	if product == nil {
		return product, applicationerrors.BadRequest(fmt.Sprintf("'Product' with id '%v' was not found", productId))
	}

	return product, nil
}

func getCart(cartRepo repositories.CartRepository, cartId int64) (*entities.Cart, *applicationerrors.ErrorStatus) {
	var cart *entities.Cart
	var err error
	cart, err = cartRepo.Get(cartId)

	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	if cart == nil {
		return cart, applicationerrors.NotFoundWithMessage(fmt.Sprintf("'Cart' with id '%v' was not found", cartId))
	}

	return cart, nil
}
