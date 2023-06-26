package services

import (
	"fmt"
	"strings"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
)

type OrderService interface {
	Add(dto.AddOrderDto) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus)
	AddFromCart() (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus)
	Get(orderId int64) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus)
	GetAll() ([]dto.OrderDto, *applicationerrors.ErrorStatus)
	GetMyOrders() ([]dto.OrderDto, *applicationerrors.ErrorStatus)
}

type orderService struct {
	repo            repositories.OrderRepository
	cartRepo        repositories.CartRepository
	productRepo     repositories.ProductRepository
	sessionProvider dto.SessionDto
}

func CreateOrderService(orderRepository repositories.OrderRepository, cartRepository repositories.CartRepository,
	productRepository repositories.ProductRepository, sessionProvider dto.SessionDto) OrderService {
	return &orderService{
		repo:            orderRepository,
		cartRepo:        cartRepository,
		productRepo:     productRepository,
		sessionProvider: sessionProvider,
	}
}

func (service *orderService) Add(addOrderDto dto.AddOrderDto) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus) {
	userId, err := valueobjects.NewId(addOrderDto.UserId.ValueInt)
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
		product, err := service.productRepo.Get(productId.ValueInt)
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

func (service *orderService) AddFromCart() (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus) {
	userIdValue, err := valueobjects.NewId(service.sessionProvider.UserId.ValueInt)
	if err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	productsInCart, err := service.cartRepo.GetAllByUser(service.sessionProvider.UserId.ValueInt)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}
	if len(productsInCart) == 0 {
		return nil, applicationerrors.BadRequest(applicationerrors.AddOrderEmptyCart)
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
	err = service.cartRepo.DeleteCartByUserId(service.sessionProvider.UserId.ValueInt)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	return dto.MapToOrderDetailsDto(*order), nil
}

func (service *orderService) Get(orderId int64) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus) {
	order, err := service.repo.Get(orderId)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	if service.sessionProvider.Role != "admin" && order.UserId.Value() != service.sessionProvider.UserId.ValueInt {
		return nil, applicationerrors.NotFound()
	}

	return dto.MapToOrderDetailsDto(*order), nil
}

func (service *orderService) GetAll() ([]dto.OrderDto, *applicationerrors.ErrorStatus) {
	ordersFromRepo, err := service.repo.GetAll()
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	orders := make([]dto.OrderDto, 0)
	for _, order := range ordersFromRepo {
		orders = append(orders, *dto.MapToOrderDto(order))
	}

	return orders, nil
}

func (service *orderService) GetMyOrders() ([]dto.OrderDto, *applicationerrors.ErrorStatus) {
	ordersFromRepo, err := service.repo.GetAllByUser(service.sessionProvider.UserId.ValueInt)

	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	orders := make([]dto.OrderDto, 0)

	for _, order := range ordersFromRepo {
		orders = append(orders, *dto.MapToOrderDto(order))
	}

	return orders, nil
}
