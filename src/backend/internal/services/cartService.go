package services

import (
	"fmt"
	"log"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
)

type CartService interface {
	AddToCart(addCart dto.AddCart) *applicationerrors.ErrorStatus
	RemoveFromCart(cartId, userId int64) *applicationerrors.ErrorStatus
	GetMyCart(userId int64) ([]dto.CartDto, *applicationerrors.ErrorStatus)
}

type cartService struct {
	repo              repositories.CartRepository
	productRepository repositories.ProductRepository
	userRepository    repositories.UserRepository
}

func CreateCartService(cartRepository repositories.CartRepository, productRepository repositories.ProductRepository, userRepository repositories.UserRepository) CartService {
	return &cartService{
		repo:              cartRepository,
		productRepository: productRepository,
		userRepository:    userRepository,
	}
}

func (service *cartService) AddToCart(addCart dto.AddCart) *applicationerrors.ErrorStatus {
	user, errorStatus := getUser(service.userRepository, addCart.UserId)
	if errorStatus != nil {
		return errorStatus
	}

	product, errorStatus := getProduct(service.productRepository, addCart.ProductId)
	if errorStatus != nil {
		return errorStatus
	}

	cart, err := entities.NewCart(0, *user, *product)
	if err != nil {
		return applicationerrors.BadRequest(err.Error())
	}

	log.Printf("'User' with id '%v' add 'Product' with id '%v'", user.Id, product.Id)
	err = service.repo.Add(cart)
	if err != nil {
		return applicationerrors.InternalError(err.Error())
	}

	return nil
}

func (service *cartService) RemoveFromCart(cartId, userId int64) *applicationerrors.ErrorStatus {
	user, errorStatus := getUser(service.userRepository, userId)
	if errorStatus != nil {
		return errorStatus
	}

	cart, errorStatus := getCart(service.repo, cartId)
	if errorStatus != nil {
		return errorStatus
	}

	log.Printf("'User' with id '%v' removes 'ItemCart' with id '%v'", user.Id, cart.Id)
	err := service.repo.Delete(*cart)
	if err != nil {
		return applicationerrors.InternalError(err.Error())
	}

	return nil
}

func (service *cartService) GetMyCart(userId int64) ([]dto.CartDto, *applicationerrors.ErrorStatus) {
	carts := make([]dto.CartDto, 0)
	user, errorStatus := getUser(service.userRepository, userId)
	if errorStatus != nil {
		return carts, errorStatus
	}

	log.Printf("'User' with id '%v' get all items from cart", user.Id)
	cartsInRepo, err := service.repo.GetAllByUser(userId)
	if err != nil {
		return carts, applicationerrors.InternalError(err.Error())
	}

	for _, cart := range cartsInRepo {
		carts = append(carts, dto.CartDto{
			Id:      cart.Id.Value(),
			Product: *dto.MapToProductDto(cart.Product),
			UserId:  userId,
		})
	}

	return carts, nil
}

func getUser(userRepo repositories.UserRepository, userId int64) (*entities.User, *applicationerrors.ErrorStatus) {
	var user *entities.User
	var err error
	user, err = userRepo.Get(userId)

	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	if user == nil {
		return user, applicationerrors.BadRequest(fmt.Sprintf("'User' with id '%v' was not found", userId))
	}

	return user, nil
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
