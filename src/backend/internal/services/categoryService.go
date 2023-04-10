package services

import (
	"fmt"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
)

type CategoryService interface {
	Add(*dto.CategoryDto) (*dto.CategoryDto, *applicationerrors.ErrorStatus)
	Update(*dto.CategoryDto) (*dto.CategoryDto, *applicationerrors.ErrorStatus)
	Delete(int64) *applicationerrors.ErrorStatus
	Get(int64) (*dto.CategoryDetailsDto, *applicationerrors.ErrorStatus)
	GetAll() ([]dto.CategoryDto, *applicationerrors.ErrorStatus)
}

type categoryService struct {
	repository repositories.CategoryRepository
}

func CreateCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		repository: repo,
	}
}

func (service *categoryService) Add(categoryDto *dto.CategoryDto) (*dto.CategoryDto, *applicationerrors.ErrorStatus) {
	if categoryDto == nil {
		return nil, applicationerrors.BadRequest("invalid 'Category'")
	}

	if err := categoryDto.Validate(); err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	name, err := valueobjects.NewName(categoryDto.Name)
	if err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	category := &entities.Category{
		Name: *name,
	}

	if errorRepo := service.repository.Add(category); errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}
	return dto.MapToCategoryDto(*category), nil
}

func (service *categoryService) Update(categoryDto *dto.CategoryDto) (*dto.CategoryDto, *applicationerrors.ErrorStatus) {
	if categoryDto == nil {
		return nil, applicationerrors.BadRequest("invalid 'Category'")
	}

	if err := categoryDto.Validate(); err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	name, err := valueobjects.NewName(categoryDto.Name)
	if err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	category, errorRepo := service.repository.Get(categoryDto.Id)
	if errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	if category == nil {
		return nil, applicationerrors.BadRequest(fmt.Sprintf("'Category' with id %v was not found", categoryDto.Id))
	}

	if category.Deleted {
		return nil, applicationerrors.BadRequest(fmt.Sprintf("'Category' with id %v was deleted", categoryDto.Id))
	}

	category.Name = *name

	if errorRepo = service.repository.Update(*category); errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	return dto.MapToCategoryDto(*category), nil
}

func (service *categoryService) Delete(id int64) *applicationerrors.ErrorStatus {
	product, errorRepo := service.repository.Get(id)

	if errorRepo != nil {
		return applicationerrors.InternalError(errorRepo.Error())
	}

	if product == nil {
		return applicationerrors.BadRequest(fmt.Sprintf("'Category' with id %v was not found", id))
	}

	if errorRepo = service.repository.Delete(*product); errorRepo != nil {
		return applicationerrors.InternalError(errorRepo.Error())
	}

	return nil
}

func (service *categoryService) Get(id int64) (*dto.CategoryDetailsDto, *applicationerrors.ErrorStatus) {
	category, errorRepo := service.repository.Get(id)

	if errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	if category == nil {
		return nil, applicationerrors.NotFound()
	}

	return dto.MapToCategoryDetailsDto(*category), nil
}

func (service *categoryService) GetAll() ([]dto.CategoryDto, *applicationerrors.ErrorStatus) {
	categories, errorRepo := service.repository.GetAll()

	if errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	categoriesDto := make([]dto.CategoryDto, 0)
	for _, category := range categories {
		categoriesDto = append(categoriesDto, *dto.MapToCategoryDto(category))
	}

	return categoriesDto, nil
}
