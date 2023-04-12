package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
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
	categoryService := createCategoryService()
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

	categoryService := createCategoryService()
	if category, err := categoryService.Get(categoryId); err != nil {
		writeErrorResponse(context, *err)
	} else {
		context.IndentedJSON(http.StatusOK, category)
	}
}

func addCategory(context *gin.Context) {
	var newCategory dto.CategoryDto

	if err := context.BindJSON(&newCategory); err != nil {
		return
	}

	categoryService := createCategoryService()
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
		return
	}

	updateCategory.Id = categoryId
	productService := createCategoryService()
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

	categoryService := createCategoryService()
	if err := categoryService.Delete(categoryId); err != nil {
		writeErrorResponse(context, *err)
	}

	context.Writer.WriteHeader(http.StatusNoContent)
}
