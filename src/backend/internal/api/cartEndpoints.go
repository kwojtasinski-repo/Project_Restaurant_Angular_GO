package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
)

func AddCartEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Cart Endpoints")
	router.GET("/carts/my", getMyCart)
	router.POST("/carts", addCart)
	router.DELETE("/carts/:id", deleteCart)
}

func getMyCart(context *gin.Context) {
	userId := context.Keys["userId"].(int64)
	cartService := createCartService()
	if categories, err := cartService.GetMyCart(userId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, categories)
	}

	ResetObjectCreator()
}

func addCart(context *gin.Context) {
	var newCart dto.AddCart

	if err := context.BindJSON(&newCart); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Cart"})
		ResetObjectCreator()
		return
	}

	userId := context.Keys["userId"].(int64)
	newCart.UserId = userId
	cartService := createCartService()
	err := cartService.AddToCart(newCart)

	if err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
		return
	}

	context.Writer.WriteHeader(http.StatusCreated)
	ResetObjectCreator()
}

func deleteCart(context *gin.Context) {
	id := context.Param("id")
	cartId, errorConvertCart := strconv.ParseInt(id, 10, 64)
	userId := context.Keys["userId"].(int64)
	if errorConvertCart != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		ResetObjectCreator()
		return
	}

	cartService := createCartService()
	if err := cartService.RemoveFromCart(cartId, userId); err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
		return
	}

	context.Writer.WriteHeader(http.StatusNoContent)
	ResetObjectCreator()
}
