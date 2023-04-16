package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
)

func AddOrderEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Order Endpoints")
	router.POST("/orders/from-cart", addOrdersFromCart)
	router.POST("/orders", addOrders)
}

func addOrdersFromCart(context *gin.Context) {
	userId := context.Keys["userId"].(int64)
	orderService, errCreateObject := createOrderService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	dto, err := orderService.AddFromCart(userId)

	if err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
	ResetObjectCreator()
}

func addOrders(context *gin.Context) {
	var newOrder dto.AddOrderDto

	if err := context.BindJSON(&newOrder); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Order"})
		ResetObjectCreator()
		return
	}
	userId := context.Keys["userId"].(int64)
	newOrder.UserId = userId

	orderService, errCreateObject := createOrderService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	dto, err := orderService.Add(newOrder)

	if err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
	ResetObjectCreator()
}
