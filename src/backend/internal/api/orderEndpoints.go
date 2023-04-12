package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
)

func AddOrderEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Order Endpoints")
	router.POST("/orders/from-cart", addOrdersFromCart)
	router.POST("/orders", addOrders)
}

func addOrdersFromCart(context *gin.Context) {
	userId := context.Keys["userId"].(int64)
	orderService := createOrderService()
	dto, err := orderService.AddFromCart(userId)

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
}

func addOrders(context *gin.Context) {
	var newOrder dto.AddOrderDto

	if err := context.BindJSON(&newOrder); err != nil {
		return
	}
	userId := context.Keys["userId"].(int64)
	newOrder.UserId = userId

	orderService := createOrderService()
	dto, err := orderService.Add(newOrder)

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
}
