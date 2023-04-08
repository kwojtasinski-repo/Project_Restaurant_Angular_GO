package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/shopspring/decimal"
)

type OrderService interface {
	Add(dto.AddOrderDto) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus)
	AddFromCart(int64) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus)
}

type orderService struct {
	repo        repositories.OderRepository
	cartRepo    repositories.CartRepository
	productRepo repositories.ProductRepository
}

func CreateOrderService(orderRepository repositories.OderRepository, cartRepository repositories.CartRepository,
	productRepository repositories.ProductRepository) OrderService {
	return &orderService{
		repo:        orderRepository,
		cartRepo:    cartRepository,
		productRepo: productRepository,
	}
}

func (service *orderService) Add(addOrderDto dto.AddOrderDto) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus) {
	order := &entities.Order{
		OrderNumber:   uuid.New().String(),
		Price:         decimal.Zero,
		Created:       time.Now(),
		UserId:        addOrderDto.UserId,
		OrderProducts: make([]entities.OrderProduct, 0),
	}
	if len(addOrderDto.ProductIds) == 0 {
		service.repo.Add(order)
		return dto.MapToOrderDetailsDto(*order), nil
	}

	var errors strings.Builder
	totalCost := decimal.Zero
	for _, productId := range addOrderDto.ProductIds {
		product, err := service.productRepo.Get(productId)
		if err != nil {
			return nil, applicationerrors.InternalError(err.Error())
		}

		if product == nil {
			fmt.Fprintf(&errors, "'Product' with id was not found %v ", productId)
			continue
		}

		order.OrderProducts = append(order.OrderProducts, entities.OrderProduct{
			Name:      product.Name(),
			Price:     product.Price(),
			ProductId: product.Id(),
		})
		totalCost = totalCost.Add(product.Price())
	}

	if errors.Len() > 0 {
		return nil, applicationerrors.BadRequest(errors.String())
	}
	order.Price = totalCost

	return dto.MapToOrderDetailsDto(*order), nil
}

func (service *orderService) AddFromCart(userId int64) (*dto.OrderDetailsDto, *applicationerrors.ErrorStatus) {
	productsInCart, err := service.cartRepo.GetAllByUser(userId)
	if err != nil {
		return nil, applicationerrors.InternalError(err.Error())
	}

	order := &entities.Order{
		OrderNumber:   uuid.New().String(),
		Price:         decimal.Zero,
		Created:       time.Now(),
		UserId:        userId,
		OrderProducts: make([]entities.OrderProduct, 0),
	}
	if len(productsInCart) == 0 {
		return nil, applicationerrors.BadRequest("'Cart' is empty, add something before create an 'Order'")
	}

	var errors strings.Builder
	totalCost := decimal.Zero
	for _, productInCart := range productsInCart {
		product, err := service.productRepo.Get(productInCart.ProductId)
		if err != nil {
			return nil, applicationerrors.InternalError(err.Error())
		}

		if product == nil {
			fmt.Fprintf(&errors, "'Product' with id was not found %v ", product.Id())
			continue
		}

		order.OrderProducts = append(order.OrderProducts, entities.OrderProduct{
			Name:      product.Name(),
			Price:     product.Price(),
			ProductId: product.Id(),
		})
		totalCost = totalCost.Add(product.Price())
	}

	if errors.Len() > 0 {
		return nil, applicationerrors.BadRequest(errors.String())
	}
	order.Price = totalCost

	return dto.MapToOrderDetailsDto(*order), nil
}
