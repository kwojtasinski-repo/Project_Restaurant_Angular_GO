package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
)

func AddOrderEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Order Endpoints")
	router.POST("/orders/from-cart", addOrdersFromCart)
	router.POST("/orders", addOrders)
	router.GET("/orders/:id", getOrder)
	router.GET("/orders/my", getMyOrders)
	router.GET("/orders", getAllOrders)
}

func addOrdersFromCart(context *gin.Context) {
	orderService, errCreateObject := createOrderService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	dto, err := orderService.AddFromCart()

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
}

func addOrders(context *gin.Context) {
	var newOrder dto.AddOrderDto

	if err := context.BindJSON(&newOrder); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Order"})
		return
	}
	userId := context.Keys["userId"].(dto.IdObject)
	newOrder.UserId = userId

	orderService, errCreateObject := createOrderService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	dto, err := orderService.Add(newOrder)

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
}

func getOrder(context *gin.Context) {
	id := context.Param("id")
	orderId, errorConvert := dto.NewIdObject(id)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	orderService, errCreateObject := createOrderService()

	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	dto, err := orderService.Get(orderId.ValueInt)

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusOK, dto)
}

func getMyOrders(context *gin.Context) {
	orderService, errCreateObject := createOrderService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	dtos, err := orderService.GetMyOrders()

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusOK, dtos)
}

func getAllOrders(context *gin.Context) {
	orderService, errCreateObject := createOrderService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	dtos, err := orderService.GetAll()

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusOK, dtos)
}
