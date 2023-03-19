package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CategoryServiceTestSuite struct {
	suite.Suite
	service CategoryService
}

func (suite *CategoryServiceTestSuite) SetupTest() {
	fmt.Println("---- Setup Before Each Test ----")
	suite.service = CreateCategoryService(repositories.NewInMemoryCategoryRepository())
}

func (suite *CategoryServiceTestSuite) addTestCategory() *dto.CategoryDto {
	dto, _ := suite.service.Add(&dto.CategoryDto{
		Name: "Name#1",
	})
	return dto
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_ValidCategory_ShouldReturnDto(t *testing.T) {
	addProduct := &dto.CategoryDto{
		Name: "Name#1",
	}

	dto, err := suite.service.Add(addProduct)

	assert.Nil(t, err)
	assert.NotNil(t, dto)
	assert.Equal(t, addProduct.Name, dto.Name)
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_TooShortCategoryName_ShouldReturnError(t *testing.T) {
	addProduct := &dto.CategoryDto{
		Name: "",
	}

	dto, err := suite.service.Add(addProduct)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' should have at least 3 characters")
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_WhiteSpacesOnName_ShouldReturnError(t *testing.T) {
	category := &dto.CategoryDto{
		Name: "                                                                                                                                                                                                                                                                        ",
	}

	dto, err := suite.service.Add(category)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' should have at least 3 characters")
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_TooLongProductName_ShouldReturnError(t *testing.T) {
	category := &dto.CategoryDto{
		Name: "NameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameName1",
	}

	dto, err := suite.service.Add(category)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' cannot have more than 200 characters")
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_Nil_ShouldReturnError(t *testing.T) {
	dto, err := suite.service.Add(nil)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "invalid 'Category'")
}

func (suite *CategoryServiceTestSuite) Test_GetCategory_ValidId_ShouldReturnDto(t *testing.T) {
	dtoAdded := suite.addTestCategory()

	dto, err := suite.service.Get(dtoAdded.Id)

	assert.Nil(t, err)
	assert.NotNil(t, dto)
}

func (suite *CategoryServiceTestSuite) Test_GetCategory_InvalidId_ShouldReturnNilProductAndError(t *testing.T) {
	dto, err := suite.service.Get(2000)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusNotFound, err.Status)
}

func (suite *CategoryServiceTestSuite) Test_GetAllCategories_ShouldReturnFilledCollection(t *testing.T) {
	dtoAdded := suite.addTestCategory()

	dtos, err := suite.service.GetAll()

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

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_ValidCategory_ShouldReturnDto(t *testing.T) {
	dtoAdded := suite.addTestCategory()
	updateDto := &dto.CategoryDto{
		Id:   dtoAdded.Id,
		Name: "Abc1235467436",
	}

	dtoUpdate, err := suite.service.Update(updateDto)

	assert.Nil(t, err)
	assert.NotNil(t, dtoUpdate)
	assert.Equal(t, updateDto.Name, dtoUpdate.Name)
}

func (suite *CategoryServiceTestSuite) Test_UpdateAndGetCategory_ValidCategory_ShouldReturnDto(t *testing.T) {
	dtoAdded := suite.addTestCategory()
	updateDto := &dto.CategoryDto{
		Id:   dtoAdded.Id,
		Name: "Abc1235467436",
	}

	dtoUpdate, err := suite.service.Update(updateDto)
	dtoAfterUpdate, errGet := suite.service.Get(dtoUpdate.Id)

	assert.Nil(t, err)
	assert.NotNil(t, dtoUpdate)
	assert.Equal(t, updateDto.Name, dtoUpdate.Name)
	assert.Nil(t, errGet)
	assert.NotNil(t, dtoAfterUpdate)
	assert.Equal(t, dtoAfterUpdate.Name, dtoUpdate.Name)
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_TooShortCategoryName_ShouldReturnError(t *testing.T) {
	updateCategory := &dto.CategoryDto{
		Name: "",
	}

	dto, err := suite.service.Update(updateCategory)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' should have at least 3 characters")
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_WhiteSpacesOnName_ShouldReturnError(t *testing.T) {
	updateCategory := &dto.CategoryDto{
		Name: "                                                                                                                                                                                                                                                                        ",
	}

	dto, err := suite.service.Update(updateCategory)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' should have at least 3 characters")
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_TooLongName_ShouldReturnError(t *testing.T) {
	updateCategory := &dto.CategoryDto{
		Name: "NameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameName1",
	}

	dto, err := suite.service.Update(updateCategory)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.NotEmpty(t, err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "'Name' cannot have more than 200 characters")
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_CategoryNotExists_ShouldReturnError(t *testing.T) {
	updateCategory := &dto.CategoryDto{
		Id:   1000,
		Name: "Name#1",
	}

	dto, err := suite.service.Update(updateCategory)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, fmt.Sprintf("'Category' with id %v was not found", updateCategory.Id))
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_NilCategory_ShouldReturnError(t *testing.T) {
	dto, err := suite.service.Update(nil)

	assert.NotNil(t, err)
	assert.Nil(t, dto)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, "invalid 'Category'")
}

func (suite *CategoryServiceTestSuite) Test_DeleteCategory_ValidId_ShouldDelete(t *testing.T) {
	dtoAdded := suite.addTestCategory()

	err := suite.service.Delete(dtoAdded.Id)

	dtoAfterDelete, errGet := suite.service.Get(dtoAdded.Id)
	assert.Nil(t, err)
	assert.Nil(t, dtoAfterDelete)
	assert.NotNil(t, errGet)
	assert.Equal(t, http.StatusNotFound, errGet.Status)
}

func (suite *CategoryServiceTestSuite) Test_DeleteCategory_InvalidId_ShouldDelete(t *testing.T) {
	var id int64 = 255

	err := suite.service.Delete(id)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Contains(t, err.Message, fmt.Sprintf("'Category' with id %v was not found", id))
}
