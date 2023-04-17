package services

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductServiceTestSuite struct {
	suite.Suite
	service ProductService
}

func (suite *ProductServiceTestSuite) SetupTest() {
	log.Println("---- Setup ProductServiceTestSuite Before Each Test ----")
	suite.service = CreateProductService(repositories.NewInMemoryProductRepository(), suite.createTestInMemoryCategoryRepository())
}

func (suite *ProductServiceTestSuite) createTestInMemoryCategoryRepository() repositories.CategoryRepository {
	repo := repositories.NewInMemoryCategoryRepository()
	id1, _ := valueobjects.NewId(1)
	id2, _ := valueobjects.NewId(2)
	id3, _ := valueobjects.NewId(3)
	category1, _ := valueobjects.NewName("Category#1")
	category2, _ := valueobjects.NewName("Category#2")
	category3, _ := valueobjects.NewName("Category#3")
	repo.Add(&entities.Category{
		Id:   *id1,
		Name: *category1,
	})
	repo.Add(&entities.Category{
		Id:   *id2,
		Name: *category2,
	})
	repo.Add(&entities.Category{
		Id:   *id3,
		Name: *category3,
	})
	return repo
}

func (suite *ProductServiceTestSuite) addTestProduct() *dto.ProductDetailsDto {
	dto, _ := suite.service.Add(&dto.AddProductDto{
		Name:        "Name#1",
		Description: "",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	})
	return dto
}

func TestProductServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}

func (suite *ProductServiceTestSuite) Test_AddProduct_ValidProduct_ShouldReturnDto() {
	addProduct := &dto.AddProductDto{
		Name:        "Name#1",
		Description: "",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := suite.service.Add(addProduct)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dto)
	assert.Equal(suite.T(), addProduct.Name, dto.Name)
	assert.Equal(suite.T(), addProduct.Description, dto.Description)
	assert.Equal(suite.T(), addProduct.Price.StringFixedBank(2), dto.Price)
	assert.Equal(suite.T(), addProduct.CategoryId, dto.Category.Id)
}

func (suite *ProductServiceTestSuite) Test_AddProduct_TooShortProductNameAndNegativePrice_ShouldReturnError() {
	addProduct := &dto.AddProductDto{
		Name:        "",
		Description: "",
		CategoryId:  1,
		Price:       decimal.New(-100, 0),
	}

	dto, err := suite.service.Add(addProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' should have at least 3 characters")
	assert.Contains(suite.T(), err.Message, "'Price' cannot be negative")
}

func (suite *ProductServiceTestSuite) Test_AddProduct_WhiteSpacesOnProductNameAndDescription_ShouldReturnError() {
	addProduct := &dto.AddProductDto{
		Name:        "                                                                                                                                                                                                                                                                        ",
		Description: "                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        ",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := suite.service.Add(addProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' should have at least 3 characters")
}

func (suite *ProductServiceTestSuite) Test_AddProduct_TooLongProductNameAndDescription_ShouldReturnError() {
	addProduct := &dto.AddProductDto{
		Name:        "NameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameName1",
		Description: "DescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescri1",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := suite.service.Add(addProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' cannot have more than 200 characters")
	assert.Contains(suite.T(), err.Message, "'Description' cannot have more than 5000 characters")
}

func (suite *ProductServiceTestSuite) Test_AddProduct_CategoryNotExists_ShouldReturnError() {
	addProduct := &dto.AddProductDto{
		Name:        "Name#1",
		Description: "",
		CategoryId:  21,
		Price:       decimal.New(100, 0),
	}

	dto, err := suite.service.Add(addProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, fmt.Sprintf("'Category' with id %v was not found", addProduct.CategoryId))
}

func (suite *ProductServiceTestSuite) Test_AddProduct_NilAddProduct_ShouldReturnError() {
	dto, err := suite.service.Add(nil)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "invalid 'Product'")
}

func (suite *ProductServiceTestSuite) Test_GetProduct_ValidId_ShouldReturnDto() {
	dtoAdded := suite.addTestProduct()

	dto, err := suite.service.Get(dtoAdded.Id)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dto)
}

func (suite *ProductServiceTestSuite) Test_GetProduct_InvalidId_ShouldReturnNilProductAndError() {
	dto, err := suite.service.Get(2000)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusNotFound, err.Status)
}

func (suite *ProductServiceTestSuite) Test_GetAllProducts_ShouldReturnFilledCollection() {
	dtoAdded := suite.addTestProduct()

	dtos, err := suite.service.GetAll()

	assert.NotNil(suite.T(), dtoAdded)
	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), dtos)
	exists := false
	for i := 0; i < len(dtos); i++ {
		if dtos[i].Id == dtoAdded.Id {
			exists = true
			break
		}
	}
	assert.True(suite.T(), exists)
}

func (suite *ProductServiceTestSuite) Test_UpdateProduct_ValidProduct_ShouldReturnDto() {
	dtoAdded := suite.addTestProduct()
	updateDto := &dto.UpdateProductDto{
		Id:          dtoAdded.Id,
		Name:        "Abc1235467436",
		Description: "Description12345465475477686799670",
		CategoryId:  2,
		Price:       decimal.New(22025, -2),
	}

	dtoUpdate, err := suite.service.Update(updateDto)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dtoUpdate)
	assert.Equal(suite.T(), updateDto.Name, dtoUpdate.Name)
	assert.Equal(suite.T(), updateDto.Description, dtoUpdate.Description)
	assert.Equal(suite.T(), updateDto.Price.StringFixedBank(2), dtoUpdate.Price)
	assert.Equal(suite.T(), updateDto.CategoryId, dtoUpdate.Category.Id)
}

func (suite *ProductServiceTestSuite) Test_UpdateAndGetProduct_ValidProduct_ShouldReturnDto() {
	dtoAdded := suite.addTestProduct()
	updateDto := &dto.UpdateProductDto{
		Id:          dtoAdded.Id,
		Name:        "Abc1235467436",
		Description: "Description12345465475477686799670",
		CategoryId:  2,
		Price:       decimal.New(22025, -2),
	}

	dtoUpdate, err := suite.service.Update(updateDto)
	dtoAfterUpdate, errGet := suite.service.Get(dtoUpdate.Id)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dtoUpdate)
	assert.Equal(suite.T(), updateDto.Name, dtoUpdate.Name)
	assert.Equal(suite.T(), updateDto.Description, dtoUpdate.Description)
	assert.Equal(suite.T(), updateDto.Price.StringFixedBank(2), dtoUpdate.Price)
	assert.Equal(suite.T(), updateDto.CategoryId, dtoUpdate.Category.Id)
	assert.Nil(suite.T(), errGet)
	assert.NotNil(suite.T(), dtoAfterUpdate)
	assert.Equal(suite.T(), dtoAfterUpdate.Name, dtoUpdate.Name)
	assert.Equal(suite.T(), dtoAfterUpdate.Description, dtoUpdate.Description)
	assert.Equal(suite.T(), dtoAfterUpdate.Price, dtoUpdate.Price)
	assert.Equal(suite.T(), dtoAfterUpdate.Category.Id, dtoUpdate.Category.Id)
}

func (suite *ProductServiceTestSuite) Test_UpdateProduct_TooShortProductNameAndNegativePrice_ShouldReturnError() {
	updateProduct := &dto.UpdateProductDto{
		Name:        "",
		Description: "",
		CategoryId:  1,
		Price:       decimal.New(-100, 0),
	}

	dto, err := suite.service.Update(updateProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' should have at least 3 characters")
	assert.Contains(suite.T(), err.Message, "'Price' cannot be negative")
}

func (suite *ProductServiceTestSuite) Test_UpdateProduct_WhiteSpacesOnProductNameAndDescription_ShouldReturnError() {
	updateProduct := &dto.UpdateProductDto{
		Name:        "                                                                                                                                                                                                                                                                        ",
		Description: "                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        ",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := suite.service.Update(updateProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' should have at least 3 characters")
}

func (suite *ProductServiceTestSuite) Test_UpdateProduct_TooLongProductNameAndDescription_ShouldReturnError() {
	updateProduct := &dto.UpdateProductDto{
		Name:        "NameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameName1",
		Description: "DescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescri1",
		CategoryId:  1,
		Price:       decimal.New(100, 0),
	}

	dto, err := suite.service.Update(updateProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' cannot have more than 200 characters")
	assert.Contains(suite.T(), err.Message, "'Description' cannot have more than 5000 characters")
}

func (suite *ProductServiceTestSuite) Test_UpdateProduct_CategoryNotExists_ShouldReturnError() {
	updateProduct := &dto.UpdateProductDto{
		Name:        "Name#1",
		Description: "",
		CategoryId:  21,
		Price:       decimal.New(100, 0),
	}

	dto, err := suite.service.Update(updateProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, fmt.Sprintf("'Category' with id %v was not found", updateProduct.CategoryId))
}

func (suite *ProductServiceTestSuite) Test_UpdateProduct_ProductNotExists_ShouldReturnError() {
	updateProduct := &dto.UpdateProductDto{
		Id:          1000,
		Name:        "Name#1",
		Description: "",
		CategoryId:  3,
		Price:       decimal.New(100, 0),
	}

	dto, err := suite.service.Update(updateProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusNotFound, err.Status)
	assert.Contains(suite.T(), err.Message, fmt.Sprintf("'Product' with id %v was not found", updateProduct.Id))
}

func (suite *ProductServiceTestSuite) Test_UpdateProduct_NilUpdateProduct_ShouldReturnError() {
	dto, err := suite.service.Update(nil)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "invalid 'Product'")
}

func (suite *ProductServiceTestSuite) Test_DeleteProduct_ValidId_ShouldDelete() {
	dtoAdded := suite.addTestProduct()

	err := suite.service.Delete(dtoAdded.Id)

	dtoAfterDelete, errGet := suite.service.Get(dtoAdded.Id)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dtoAfterDelete)
	assert.Equal(suite.T(), true, dtoAfterDelete.Deleted)
	assert.Nil(suite.T(), errGet)
}

func (suite *ProductServiceTestSuite) Test_DeleteProduct_InvalidId_ShouldDelete() {
	var id int64 = 255

	err := suite.service.Delete(id)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusNotFound, err.Status)
	assert.Contains(suite.T(), err.Message, fmt.Sprintf("'Product' with id %v was not found", id))
}
