package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
)

func AddCartEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Cart Endpoints")
	router.GET("/carts/my", getMyCart)
	router.POST("/carts", addCart)
	router.DELETE("/carts/:id", deleteCart)
}

func getMyCart(context *gin.Context) {
	userId := context.Keys["userId"].(dto.IdObject)
	cartService, errCreateObject := createCartService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	if myCart, err := cartService.GetMyCart(userId.ValueInt); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, myCart)
	}
}

func addCart(context *gin.Context) {
	var newCart dto.AddCart

	if err := context.BindJSON(&newCart); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": applicationerrors.InvalidCart})
		return
	}

	userId := context.Keys["userId"].(dto.IdObject)
	newCart.UserId = userId
	cartService, errCreateObject := createCartService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	err := cartService.AddToCart(newCart)

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.Writer.WriteHeader(http.StatusCreated)
}

func deleteCart(context *gin.Context) {
	id := context.Param("id")
	cartId, errorConvertCart := dto.NewIdObject(id)
	userId := context.Keys["userId"].(dto.IdObject)
	if errorConvertCart != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": applicationerrors.InvalidId})
		return
	}

	cartService, errCreateObject := createCartService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	if err := cartService.RemoveFromCart(cartId.ValueInt, userId.ValueInt); err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.Writer.WriteHeader(http.StatusNoContent)
}
