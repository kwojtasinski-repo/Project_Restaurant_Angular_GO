package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
)

func AddCategoryEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Category Endpoints")
	router.GET("/categories", getCategories)
	router.GET("/categories/:id", getCategory)
	endpointWithPermissionAdmin := router.Group("/categories", PermissionMiddleware("admin"))
	endpointWithPermissionAdmin.POST("", addCategory)
	endpointWithPermissionAdmin.PUT("/:id", updateCategory)
	endpointWithPermissionAdmin.DELETE("/:id", deleteCategory)
}

func getCategories(context *gin.Context) {
	categoryService, errCreateObject := createCategoryService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	if categories, err := categoryService.GetAll(); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, categories)
	}
}

func getCategory(context *gin.Context) {
	id := context.Param("id")
	categoryId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	categoryService, errCreateObject := createCategoryService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	if category, err := categoryService.Get(categoryId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, category)
	}
}

func addCategory(context *gin.Context) {
	var newCategory dto.CategoryDto

	if err := context.BindJSON(&newCategory); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Category"})
		return
	}

	categoryService, errCreateObject := createCategoryService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	dto, err := categoryService.Add(&newCategory)

	if err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
}

func updateCategory(context *gin.Context) {
	id := context.Param("id")
	categoryId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var updateCategory dto.CategoryDto

	if err := context.BindJSON(&updateCategory); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Category"})
		return
	}

	updateCategory.Id = categoryId
	productService, errCreateObject := createCategoryService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	dto, errUpdate := productService.Update(&updateCategory)

	if errUpdate != nil {
		writeErrorResponse(context, *errUpdate)
		return
	}

	context.IndentedJSON(http.StatusOK, dto)
}

func deleteCategory(context *gin.Context) {
	id := context.Param("id")
	categoryId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	categoryService, errCreateObject := createCategoryService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
		return
	}

	if err := categoryService.Delete(categoryId); err != nil {
		writeErrorResponse(context, *err)
		return
	}

	context.Writer.WriteHeader(http.StatusNoContent)
}
