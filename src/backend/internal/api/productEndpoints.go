package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
)

func AddProductEndpoints(router *gin.Engine) {
	router.GET("/api/products", getProducts)
	router.GET("/api/products/:id", getProduct)
	router.POST("/api/products", addProduct)
	router.PUT("/api/products/:id", updateProduct)
	router.DELETE("/api/products/:id", deleteProduct)
}

func getProducts(context *gin.Context) {
	productService := createProductService()
	if products, err := productService.GetAll(); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, products)
	}

}

func getProduct(context *gin.Context) {
	id := context.Param("id")
	productId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	productService := createProductService()
	if products, err := productService.Get(productId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, products)
	}
}

func addProduct(context *gin.Context) {
	var newProduct dto.AddProductDto

	if err := context.BindJSON(&newProduct); err != nil {
		return
	}

	productService := createProductService()
	dto, err := productService.Add(&newProduct)

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
}

func updateProduct(context *gin.Context) {
	id := context.Param("id")
	productId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var updateProduct dto.UpdateProductDto

	if err := context.BindJSON(&updateProduct); err != nil {
		return
	}

	updateProduct.Id = productId
	productService := createProductService()
	dto, errUpdate := productService.Update(&updateProduct)

	if errUpdate != nil {
		writeErrorResponse(context, *errUpdate)
		return
	}

	context.IndentedJSON(http.StatusOK, dto)
}

func deleteProduct(context *gin.Context) {
	id := context.Param("id")
	productId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	productService := createProductService()
	if err := productService.Delete(productId); err != nil {
		writeErrorResponse(context, *err)
	}

	context.Writer.WriteHeader(http.StatusNoContent)
}
