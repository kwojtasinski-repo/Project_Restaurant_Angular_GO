package services

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CategoryServiceTestSuite struct {
	suite.Suite
	service CategoryService
}

func (suite *CategoryServiceTestSuite) SetupTest() {
	log.Println("---- Setup CategoryServiceTestSuite Before Each Test ----")
	suite.service = CreateCategoryService(repositories.NewInMemoryCategoryRepository())
}

func (suite *CategoryServiceTestSuite) addTestCategory() *dto.CategoryDto {
	dto, _ := suite.service.Add(&dto.CategoryDto{
		Name: "Name#1",
	})
	return dto
}

func TestCategoryServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryServiceTestSuite))
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_ValidCategory_ShouldReturnDto() {
	addProduct := &dto.CategoryDto{
		Name: "Name#1",
	}

	dto, err := suite.service.Add(addProduct)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dto)
	assert.Equal(suite.T(), addProduct.Name, dto.Name)
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_TooShortCategoryName_ShouldReturnError() {
	addProduct := &dto.CategoryDto{
		Name: "",
	}

	dto, err := suite.service.Add(addProduct)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' should have at least 3 characters")
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_WhiteSpacesOnName_ShouldReturnError() {
	category := &dto.CategoryDto{
		Name: "                                                                                                                                                                                                                                                                        ",
	}

	dto, err := suite.service.Add(category)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' should have at least 3 characters")
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_TooLongProductName_ShouldReturnError() {
	category := &dto.CategoryDto{
		Name: "NameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameName1",
	}

	dto, err := suite.service.Add(category)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' cannot have more than 200 characters")
}

func (suite *CategoryServiceTestSuite) Test_AddCategory_Nil_ShouldReturnError() {
	dto, err := suite.service.Add(nil)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "invalid 'Category'")
}

func (suite *CategoryServiceTestSuite) Test_GetCategory_ValidId_ShouldReturnDto() {
	dtoAdded := suite.addTestCategory()

	dto, err := suite.service.Get(dtoAdded.Id.ValueInt)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dto)
}

func (suite *CategoryServiceTestSuite) Test_GetCategory_InvalidId_ShouldReturnNilProductAndError() {
	dto, err := suite.service.Get(2000)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusNotFound, err.Status)
}

func (suite *CategoryServiceTestSuite) Test_GetAllCategories_ShouldReturnFilledCollection() {
	dtoAdded := suite.addTestCategory()

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

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_ValidCategory_ShouldReturnDto() {
	dtoAdded := suite.addTestCategory()
	updateDto := &dto.CategoryDto{
		Id:   dtoAdded.Id,
		Name: "Abc1235467436",
	}

	dtoUpdate, err := suite.service.Update(updateDto)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dtoUpdate)
	assert.Equal(suite.T(), updateDto.Name, dtoUpdate.Name)
}

func (suite *CategoryServiceTestSuite) Test_UpdateAndGetCategory_ValidCategory_ShouldReturnDto() {
	dtoAdded := suite.addTestCategory()
	updateDto := &dto.CategoryDto{
		Id:   dtoAdded.Id,
		Name: "Abc1235467436",
	}

	dtoUpdate, err := suite.service.Update(updateDto)
	dtoAfterUpdate, errGet := suite.service.Get(dtoUpdate.Id.ValueInt)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dtoUpdate)
	assert.Equal(suite.T(), updateDto.Name, dtoUpdate.Name)
	assert.Nil(suite.T(), errGet)
	assert.NotNil(suite.T(), dtoAfterUpdate)
	assert.Equal(suite.T(), dtoAfterUpdate.Name, dtoUpdate.Name)
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_TooShortCategoryName_ShouldReturnError() {
	updateCategory := &dto.CategoryDto{
		Name: "",
	}

	dto, err := suite.service.Update(updateCategory)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' should have at least 3 characters")
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_WhiteSpacesOnName_ShouldReturnError() {
	updateCategory := &dto.CategoryDto{
		Name: "                                                                                                                                                                                                                                                                        ",
	}

	dto, err := suite.service.Update(updateCategory)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' should have at least 3 characters")
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_TooLongName_ShouldReturnError() {
	updateCategory := &dto.CategoryDto{
		Name: "NameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameNameName1",
	}

	dto, err := suite.service.Update(updateCategory)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.NotEmpty(suite.T(), err.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "'Name' cannot have more than 200 characters")
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_CategoryNotExists_ShouldReturnError() {
	updateCategory := &dto.CategoryDto{
		Id:   dto.IdObject{ValueInt: 1000},
		Name: "Name#1",
	}

	dto, err := suite.service.Update(updateCategory)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusNotFound, err.Status)
	assert.Contains(suite.T(), err.Message, fmt.Sprintf("'Category' with id %v was not found", updateCategory.Id))
}

func (suite *CategoryServiceTestSuite) Test_UpdateCategory_NilCategory_ShouldReturnError() {
	dto, err := suite.service.Update(nil)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), dto)
	assert.Equal(suite.T(), http.StatusBadRequest, err.Status)
	assert.Contains(suite.T(), err.Message, "invalid 'Category'")
}

func (suite *CategoryServiceTestSuite) Test_DeleteCategory_ValidId_ShouldDelete() {
	dtoAdded := suite.addTestCategory()

	err := suite.service.Delete(dtoAdded.Id.ValueInt)

	dtoAfterDelete, errGet := suite.service.Get(dtoAdded.Id.ValueInt)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), dtoAfterDelete)
	assert.Equal(suite.T(), true, dtoAfterDelete.Deleted)
	assert.Nil(suite.T(), errGet)
}

func (suite *CategoryServiceTestSuite) Test_DeleteCategory_InvalidId_ShouldDelete() {
	var id int64 = 255

	err := suite.service.Delete(id)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusNotFound, err.Status)
	assert.Contains(suite.T(), err.Message, fmt.Sprintf("'Category' with id %v was not found", id))
}

func (suite *CategoryServiceTestSuite) Test_DeleteCategory_AnErrorOccuredInCategoryRepository_ShouldReturnInternalServerError() {
	service := CreateCategoryService(repositories.NewErrorCategoryRepository())

	err := service.Delete(1)

	suite.Assertions.NotNil(err)
	suite.Assertions.Equal(err.Status, http.StatusInternalServerError)
}
