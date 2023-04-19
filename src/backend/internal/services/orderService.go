package services

import (
	"fmt"
	"strings"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
)

type OrderService interface {
	Add(dto.AddOrderDto) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus)
	AddFromCart(userId int64) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus)
}

type orderService struct {
	repo        repositories.OrderRepository
	cartRepo    repositories.CartRepository
	productRepo repositories.ProductRepository
}

func CreateOrderService(orderRepository repositories.OrderRepository, cartRepository repositories.CartRepository,
	productRepository repositories.ProductRepository) OrderService {
	return &orderService{
		repo:        orderRepository,
		cartRepo:    cartRepository,
		productRepo: productRepository,
	}
}

func (service *orderService) Add(addOrderDto dto.AddOrderDto) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus) {
	userId, err := valueobjects.NewId(addOrderDto.UserId)
	if err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	var order *entities.Order
	order, err = entities.NewOrder(0, entities.User{
		Id: *userId,
	}, make([]entities.OrderProduct, 0))

	if err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	if len(addOrderDto.ProductIds) == 0 {
		err = service.repo.Add(order)
		if err != nil {
			return nil, applicationerrors.InternalError(err.Error())
		}
		return dto.MapToOrderDetailsDto(*order), nil
	}

	var errors strings.Builder
	for _, productId := range addOrderDto.ProductIds {
		product, err := service.productRepo.Get(productId)
		if err != nil {
			return nil, applicationerrors.InternalError(err.Error())
		}

		if product == nil {
			fmt.Fprintf(&errors, "'Product' with id was not found %v ", productId)
			continue
		}

		err = order.AddProduct(entities.OrderProduct{
			Name:      product.Name,
			Price:     product.Price,
			ProductId: product.Id,
		})
		if err != nil {
			errors.WriteString(err.Error())
		}
	}

	if errors.Len() > 0 {
		return nil, applicationerrors.BadRequest(errors.String())
	}

	err = service.repo.Add(order)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	return dto.MapToOrderDetailsDto(*order), nil
}

func (service *orderService) AddFromCart(userId int64) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus) {
	userIdValue, err := valueobjects.NewId(userId)
	if err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	productsInCart, err := service.cartRepo.GetAllByUser(userId)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}
	if len(productsInCart) == 0 {
		return nil, applicationerrors.BadRequest("'Cart' is empty, add something before create an 'Order'")
	}

	var order *entities.Order
	order, err = entities.NewOrder(0, entities.User{
		Id: *userIdValue,
	}, make([]entities.OrderProduct, 0))

	if err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	var errors strings.Builder
	for _, productInCart := range productsInCart {
		product, err := service.productRepo.Get(productInCart.ProductId.Value())
		if err != nil {
			return nil, applicationerrors.InternalError(err.Error())
		}

		if product == nil {
			fmt.Fprintf(&errors, "'Product' with id was not found %v ", productInCart.ProductId.Value())
			continue
		}

		order.AddProduct(entities.OrderProduct{
			Name:      product.Name,
			Price:     product.Price,
			ProductId: product.Id,
		})
	}

	if errors.Len() > 0 {
		return nil, applicationerrors.BadRequest(errors.String())
	}

	err = service.repo.Add(order)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}
	err = service.cartRepo.DeleteCartByUserId(userId)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	return dto.MapToOrderDetailsDto(*order), nil
}
