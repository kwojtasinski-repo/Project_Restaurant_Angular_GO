package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
)

func AddProductEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Product Endpoints")
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/products", addProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)
}

func getProducts(context *gin.Context) {
	productService, errCreateObject := createProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	if products, err := productService.GetAll(); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, products)
	}
	ResetObjectCreator()
}

func getProduct(context *gin.Context) {
	id := context.Param("id")
	productId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		ResetObjectCreator()
		return
	}

	productService, errCreateObject := createProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	if products, err := productService.Get(productId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, products)
	}
	ResetObjectCreator()
}

func addProduct(context *gin.Context) {
	var newProduct dto.AddProductDto

	if err := context.BindJSON(&newProduct); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Product"})
		ResetObjectCreator()
		return
	}

	productService, errCreateObject := createProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	dto, err := productService.Add(&newProduct)

	if err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
	ResetObjectCreator()
}

func updateProduct(context *gin.Context) {
	id := context.Param("id")
	productId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		ResetObjectCreator()
		return
	}

	var updateProduct dto.UpdateProductDto

	if err := context.BindJSON(&updateProduct); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Product"})
		ResetObjectCreator()
		return
	}

	updateProduct.Id = productId
	productService, errCreateObject := createProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	dto, errUpdate := productService.Update(&updateProduct)

	if errUpdate != nil {
		writeErrorResponse(context, *errUpdate)
		ResetObjectCreator()
		return
	}

	context.IndentedJSON(http.StatusOK, dto)
	ResetObjectCreator()
}

func deleteProduct(context *gin.Context) {
	id := context.Param("id")
	productId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		ResetObjectCreator()
		return
	}

	productService, errCreateObject := createProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	if err := productService.Delete(productId); err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
	}

	context.Writer.WriteHeader(http.StatusNoContent)
	ResetObjectCreator()
}
