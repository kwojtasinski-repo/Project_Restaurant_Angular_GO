package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TODO refactor and use TestSuite
var testProductService = CreateProductService(repositories.NewInMemoryProductRepository(), createTestInMemoryCategoryRepository())

func createTestInMemoryCategoryRepository() repositories.CategoryRepository {
	repo := repositories.NewInMemoryCategoryRepository()
	repo.Add(&entities.Category{
		Id:   1,
		Name: "Category#1",
	})
	repo.Add(&entities.Category{
		Id:   2,
		Name: "Category#2",
	})
	repo.Add(&entities.Category{
		Id:   3,
		Name: "Category#3",
	})
	return repo
}

func addTestProduct(service ProductService) *dto.ProductDetailsDto {
	dto, _ := service.Add(&dto.AddProductDto{
		Name:        "Name#1",
		Description: "",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	})
	return dto
}

func Test_AddProduct_ValidProduct_ShouldReturnDto(t *testing.T) {
	addProduct := &dto.AddProductDto{
		Name:        "Name#1",
		Description: "",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := testProductService.Add(addProduct)

	assert.Nil(t, err)
	assert.NotNil(t, dto)
	assert.Equal(t, addProduct.Name, dto.Name)
	assert.Equal(t, addProduct.Description, dto.Description)
	assert.Equal(t, addProduct.Price, dto.Price)
	assert.Equal(t, addProduct.CategoryId, dto.Category.Id)
}

func Test_AddProduct_TooShortProductNameAndNegativePrice_ShouldReturnError(t *testing.T) {
	addProduct := &dto.AddProductDto{
		Name:        "",
		Description: "",
		CategoryId:  1,
		Price:       decimal.New(-100, 0),
	}

	dto, err := testProductService.Add(addProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' should have at least 3 characters")
	assert.Contains(t, err.Message, "'Price' cannot be negative")
}

func Test_AddProduct_WhiteSpacesOnProductNameAndDescription_ShouldReturnError(t *testing.T) {
	addProduct := &dto.AddProductDto{
		Name:        "                                                                                                                                                                                                                                                                        ",
		Description: "                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        ",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := testProductService.Add(addProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' should have at least 3 characters")
}

func Test_AddProduct_TooLongProductNameAndDescription_ShouldReturnError(t *testing.T) {
	addProduct := &dto.AddProductDto{
		Name:        "NameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameName1",
		Description: "DescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescri1",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := testProductService.Add(addProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' cannot have more than 200 characters")
	assert.Contains(t, err.Message, "'Description' cannot have more than 5000 characters")
}

func Test_AddProduct_CategoryNotExists_ShouldReturnError(t *testing.T) {
	addProduct := &dto.AddProductDto{
		Name:        "Name#1",
		Description: "",
		CategoryId:  21,
		Price:       decimal.New(100, 0),
	}

	dto, err := testProductService.Add(addProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, fmt.Sprintf("'Category' with id %v was not found", addProduct.CategoryId))
}

func Test_AddProduct_NilAddProduct_ShouldReturnError(t *testing.T) {
	dto, err := testProductService.Add(nil)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "invalid 'Product'")
}

func Test_GetProduct_ValidId_ShouldReturnDto(t *testing.T) {
	dtoAdded := addTestProduct(testProductService)

	dto, err := testProductService.Get(dtoAdded.Id)

	assert.Nil(t, err)
	assert.NotNil(t, dto)
}

func Test_GetProduct_InvalidId_ShouldReturnNilProductAndError(t *testing.T) {
	dto, err := testProductService.Get(2000)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusNotFound, err.Status)
}

func Test_GetAllProducts_ShouldReturnFilledCollection(t *testing.T) {
	dtoAdded := addTestProduct(testProductService)

	dtos, err := testProductService.GetAll()

	assert.NotNil(t, dtoAdded)
	assert.Nil(t, err)
	assert.NotEmpty(t, dtos)
	exists := false
	for i := 0; i < len(dtos); i++ {
		if dtos[i].Id == dtoAdded.Id {
			exists = true
			break
		}
	}
	assert.True(t, exists)
}

func Test_UpdateProduct_ValidProduct_ShouldReturnDto(t *testing.T) {
	dtoAdded := addTestProduct(testProductService)
	updateDto := &dto.UpdateProductDto{
		Id:          dtoAdded.Id,
		Name:        "Abc1235467436",
		Description: "Description12345465475477686799670",
		CategoryId:  2,
		Price:       decimal.New(22025, -2),
	}

	dtoUpdate, err := testProductService.Update(updateDto)

	assert.Nil(t, err)
	assert.NotNil(t, dtoUpdate)
	assert.Equal(t, updateDto.Name, dtoUpdate.Name)
	assert.Equal(t, updateDto.Description, dtoUpdate.Description)
	assert.Equal(t, updateDto.Price, dtoUpdate.Price)
	assert.Equal(t, updateDto.CategoryId, dtoUpdate.Category.Id)
}

func Test_UpdateAndGetProduct_ValidProduct_ShouldReturnDto(t *testing.T) {
	dtoAdded := addTestProduct(testProductService)
	updateDto := &dto.UpdateProductDto{
		Id:          dtoAdded.Id,
		Name:        "Abc1235467436",
		Description: "Description12345465475477686799670",
		CategoryId:  2,
		Price:       decimal.New(22025, -2),
	}

	dtoUpdate, err := testProductService.Update(updateDto)
	dtoAfterUpdate, errGet := testProductService.Get(dtoUpdate.Id)

	assert.Nil(t, err)
	assert.NotNil(t, dtoUpdate)
	assert.Equal(t, updateDto.Name, dtoUpdate.Name)
	assert.Equal(t, updateDto.Description, dtoUpdate.Description)
	assert.Equal(t, updateDto.Price, dtoUpdate.Price)
	assert.Equal(t, updateDto.CategoryId, dtoUpdate.Category.Id)
	assert.Nil(t, errGet)
	assert.NotNil(t, dtoAfterUpdate)
	assert.Equal(t, dtoAfterUpdate.Name, dtoUpdate.Name)
	assert.Equal(t, dtoAfterUpdate.Description, dtoUpdate.Description)
	assert.Equal(t, dtoAfterUpdate.Price, dtoUpdate.Price)
	assert.Equal(t, dtoAfterUpdate.Category.Id, dtoUpdate.Category.Id)
}

func Test_UpdateProduct_TooShortProductNameAndNegativePrice_ShouldReturnError(t *testing.T) {
	updateProduct := &dto.UpdateProductDto{
		Name:        "",
		Description: "",
		CategoryId:  1,
		Price:       decimal.New(-100, 0),
	}

	dto, err := testProductService.Update(updateProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' should have at least 3 characters")
	assert.Contains(t, err.Message, "'Price' cannot be negative")
}

func Test_UpdateProduct_WhiteSpacesOnProductNameAndDescription_ShouldReturnError(t *testing.T) {
	updateProduct := &dto.UpdateProductDto{
		Name:        "                                                                                                                                                                                                                                                                        ",
		Description: "                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        ",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := testProductService.Update(updateProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' should have at least 3 characters")
}

func Test_UpdateProduct_TooLongProductNameAndDescription_ShouldReturnError(t *testing.T) {
	updateProduct := &dto.UpdateProductDto{
		Name:        "NameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameName1",
		Description: "DescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescri1",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := testProductService.Update(updateProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' cannot have more than 200 characters")
	assert.Contains(t, err.Message, "'Description' cannot have more than 5000 characters")
}

func Test_UpdateProduct_CategoryNotExists_ShouldReturnError(t *testing.T) {
	updateProduct := &dto.UpdateProductDto{
		Name:        "Name#1",
		Description: "",
		CategoryId:  21,
		Price:       decimal.New(100, 0),
	}

	dto, err := testProductService.Update(updateProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, fmt.Sprintf("'Category' with id %v was not found", updateProduct.CategoryId))
}

func Test_UpdateProduct_ProductNotExists_ShouldReturnError(t *testing.T) {
	updateProduct := &dto.UpdateProductDto{
		Id:          1000,
		Name:        "Name#1",
		Description: "",
		CategoryId:  3,
		Price:       decimal.New(100, 0),
	}

	dto, err := testProductService.Update(updateProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, fmt.Sprintf("'Product' with id %v was not found", updateProduct.Id))
}

func Test_UpdateProduct_NilUpdateProduct_ShouldReturnError(t *testing.T) {
	dto, err := testProductService.Update(nil)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "invalid 'Product'")
}

func Test_DeleteProduct_ValidId_ShouldDelete(t *testing.T) {
	dtoAdded := addTestProduct(testProductService)

	err := testProductService.Delete(dtoAdded.Id)

	dtoAfterDelete, errGet := testProductService.Get(dtoAdded.Id)
	assert.Nil(t, err)
	assert.Nil(t, dtoAfterDelete)
	assert.NotNil(t, errGet)
	assert.Equal(t, http.StatusNotFound, errGet.Status)
}

func Test_DeleteProduct_InvalidId_ShouldDelete(t *testing.T) {
	var id int64 = 255

	err := testProductService.Delete(id)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, fmt.Sprintf("'Product' with id %v was not found", id))
}
