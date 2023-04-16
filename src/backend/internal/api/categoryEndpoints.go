package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
)

func AddCategoryEndpoints(router *gin.RouterGroup) {
	log.Println("Setup Category Endpoints")
	router.GET("/categories", getCategories)
	router.GET("/categories/:id", getCategory)
	router.POST("/categories", addCategory)
	router.PUT("/categories/:id", updateCategory)
	router.DELETE("/categories/:id", deleteCategory)
}

func getCategories(context *gin.Context) {
	categoryService, errCreateObject := createCategoryService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	if categories, err := categoryService.GetAll(); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, categories)
	}

	ResetObjectCreator()
}

func getCategory(context *gin.Context) {
	id := context.Param("id")
	categoryId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		ResetObjectCreator()
		return
	}

	categoryService, errCreateObject := createCategoryService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	if category, err := categoryService.Get(categoryId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, category)
	}
	ResetObjectCreator()
}

func addCategory(context *gin.Context) {
	var newCategory dto.CategoryDto

	if err := context.BindJSON(&newCategory); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Category"})
		ResetObjectCreator()
		return
	}

	categoryService, errCreateObject := createCategoryService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	dto, err := categoryService.Add(&newCategory)

	if err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
		return
	}

	context.IndentedJSON(http.StatusCreated, dto)
	ResetObjectCreator()
}

func updateCategory(context *gin.Context) {
	id := context.Param("id")
	categoryId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		ResetObjectCreator()
		return
	}

	var updateCategory dto.CategoryDto

	if err := context.BindJSON(&updateCategory); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Category"})
		ResetObjectCreator()
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
		ResetObjectCreator()
		return
	}

	context.IndentedJSON(http.StatusOK, dto)
	ResetObjectCreator()
}

func deleteCategory(context *gin.Context) {
	id := context.Param("id")
	categoryId, errorConvert := strconv.ParseInt(id, 10, 64)

	if errorConvert != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		ResetObjectCreator()
		return
	}

	categoryService, errCreateObject := createCategoryService()
	if errCreateObject != nil {
		writeErrorResponse(context, *applicationerrors.InternalError(errCreateObject.Error()))
	}

	if err := categoryService.Delete(categoryId); err != nil {
		writeErrorResponse(context, *err)
		ResetObjectCreator()
	}

	context.Writer.WriteHeader(http.StatusNoContent)
	ResetObjectCreator()
}
