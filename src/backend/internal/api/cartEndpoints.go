package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
)

func AddCartEndpoints(router *gin.Engine) {
	log.Println("Setup Cart Endpoints")
	router.GET("/api/carts/:id", getMyCart)
	router.POST("/api/carts", addCart)
	router.DELETE("/api/carts", deleteCart)
}

func getMyCart(context *gin.Context) {
	id := context.Param("id")
	userId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	cartService := createCartService()
	if categories, err := cartService.GetMyCart(userId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, categories)
	}

}

func addCart(context *gin.Context) {
	var newCart dto.AddCart

	if err := context.BindJSON(&newCart); err != nil {
		return
	}

	cartService := createCartService()
	err := cartService.AddToCart(newCart)

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.Writer.WriteHeader(http.StatusCreated)
}

func deleteCart(context *gin.Context) {
	id, errId := context.GetQuery("id")
	userIdString, errUserId := context.GetQuery("userId")
	cartId, errorConvertCart := strconv.ParseInt(id, 10, 64)
	userId, errorConvertUser := strconv.ParseInt(userIdString, 10, 64)

	if !errId || !errUserId || errorConvertCart != nil || errorConvertUser != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	cartService := createCartService()
	if err := cartService.RemoveFromCart(cartId, userId); err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.Writer.WriteHeader(http.StatusNoContent)
}
