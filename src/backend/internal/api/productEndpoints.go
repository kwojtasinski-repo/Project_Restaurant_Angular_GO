package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
)

func AddProductEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Product Endpoints")
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	endpointWithPermissionAdmin := router.Group("/products", PermissionMiddleware("admin"))
	endpointWithPermissionAdmin.POST("", addProduct)
	endpointWithPermissionAdmin.PUT("/:id", updateProduct)
	endpointWithPermissionAdmin.DELETE("/:id", deleteProduct)
}

func getProducts(context *gin.Context) {
	productService, errCreateObject := CreateProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	if products, err := productService.GetAll(); err != nil {
		writeErrorResponse(context, *err)
		return
	} else {
		context.IndentedJSON(http.StatusOK, products)
	}
}

func getProduct(context *gin.Context) {
	id := context.Param("id")
	productId, errorConvert := dto.NewIdObject(id)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	productService, errCreateObject := CreateProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	if products, err := productService.Get(productId.ValueInt); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, products)
	}
}

func addProduct(context *gin.Context) {
	var newProduct dto.AddProductDto

	if err := context.BindJSON(&newProduct); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Product"})
		return
	}

	productService, errCreateObject := CreateProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	dto, err := productService.Add(&newProduct)

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
}

func updateProduct(context *gin.Context) {
	id := context.Param("id")
	productId, errorConvert := dto.NewIdObject(id)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var updateProduct dto.UpdateProductDto

	if err := context.BindJSON(&updateProduct); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Product"})
		return
	}

	updateProduct.Id = *productId
	productService, errCreateObject := CreateProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	dto, errUpdate := productService.Update(&updateProduct)

	if errUpdate != nil {
		writeErrorResponse(context, *errUpdate)
		return
	}

	context.IndentedJSON(http.StatusOK, dto)
}

func deleteProduct(context *gin.Context) {
	id := context.Param("id")
	productId, errorConvert := dto.NewIdObject(id)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	productService, errCreateObject := CreateProductService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	if err := productService.Delete(productId.ValueInt); err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.Writer.WriteHeader(http.StatusNoContent)
}
