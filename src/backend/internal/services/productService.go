package services

import (
	"fmt"
	"strings"

	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	applicationerrors "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
)

type ProductService interface {
	Update(*dto.UpdateProductDto) (*dto.ProductDetailsDto, *applicationerrors.ErrorStatus)
	Add(*dto.AddProductDto) (*dto.ProductDetailsDto, *applicationerrors.ErrorStatus)
	Delete(int64) *applicationerrors.ErrorStatus
	Get(int64) (*dto.ProductDetailsDto, *applicationerrors.ErrorStatus)
	GetAll() ([]dto.ProductDto, *applicationerrors.ErrorStatus)
}

type productService struct {
	repository         repositories.ProductRepository
	categoryRepository repositories.CategoryRepository
}

func CreateProductService(repo repositories.ProductRepository, categoryRepository repositories.CategoryRepository) ProductService {
	return &productService{
		repository:         repo,
		categoryRepository: categoryRepository,
	}
}

func (service *productService) Add(productDto *dto.AddProductDto) (*dto.ProductDetailsDto, *applicationerrors.ErrorStatus) {
	if productDto == nil {
		return nil, applicationerrors.BadRequest("invalid 'Product'")
	}

	if err := productDto.Validate(); err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	category, errorRepo := service.categoryRepository.Get(productDto.CategoryId)
	if errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	if category == nil {
		return nil, applicationerrors.BadRequest(fmt.Sprintf("'Category' with id %v was not found", productDto.CategoryId))
	}

	product, err := entities.NewProduct(0, productDto.Name, productDto.Price, productDto.Description, category)
	if err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	if errorRepo := service.repository.Add(product); errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}
	return dto.MapToProductDetailsDto(*product), nil
}

func (service *productService) Update(productDto *dto.UpdateProductDto) (*dto.ProductDetailsDto, *applicationerrors.ErrorStatus) {
	if productDto == nil {
		return nil, applicationerrors.BadRequest("invalid 'Product'")
	}

	if err := productDto.Validate(); err != nil {
		return nil, applicationerrors.BadRequest(err.Error())
	}

	category, errorRepo := service.categoryRepository.Get(productDto.CategoryId)
	if errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	if category == nil {
		return nil, applicationerrors.BadRequest(fmt.Sprintf("'Category' with id %v was not found", productDto.CategoryId))
	}

	product, errorRepo := service.repository.Get(productDto.Id)
	if errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	if product == nil {
		return nil, applicationerrors.NotFoundWithMessage(fmt.Sprintf("'Product' with id %v was not found", productDto.Id))
	}

	if product.Deleted {
		return nil, applicationerrors.BadRequest(fmt.Sprintf("'Product' with id %v was deleted", productDto.Id))
	}

	var err error
	var name *valueobjects.Name
	var description *valueobjects.Description
	var price *valueobjects.Price
	var validationErrors strings.Builder
	if name, err = valueobjects.NewName(productDto.Name); err != nil {
		validationErrors.WriteString(err.Error())
	}
	if description, err = valueobjects.NewDescription(productDto.Description); err != nil {
		validationErrors.WriteString(err.Error())
	}
	if price, err = valueobjects.NewPrice(productDto.Price); err != nil {
		validationErrors.WriteString(err.Error())
	}
	if validationErrors.Len() > 0 {
		return nil, applicationerrors.BadRequest(validationErrors.String())
	}
	product.Name = *name
	product.Description = *description
	product.Price = *price
	product.Category = *category

	if errorRepo = service.repository.Update(*product); errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	return dto.MapToProductDetailsDto(*product), nil
}

func (service *productService) Delete(id int64) *applicationerrors.ErrorStatus {
	product, errorRepo := service.repository.Get(id)

	if errorRepo != nil {
		return applicationerrors.InternalError(errorRepo.Error())
	}

	if product == nil {
		return applicationerrors.NotFoundWithMessage(fmt.Sprintf("'Product' with id %v was not found", id))
	}

	if errorRepo = service.repository.Delete(*product); errorRepo != nil {
		return applicationerrors.InternalError(errorRepo.Error())
	}

	return nil
}

func (service *productService) Get(id int64) (*dto.ProductDetailsDto, *applicationerrors.ErrorStatus) {
	product, errorRepo := service.repository.Get(id)

	if errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	if product == nil {
		return nil, applicationerrors.NotFound()
	}

	return dto.MapToProductDetailsDto(*product), nil
}

func (service *productService) GetAll() ([]dto.ProductDto, *applicationerrors.ErrorStatus) {
	products, errorRepo := service.repository.GetAll()

	if errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	productsDto := make([]dto.ProductDto, 0)
	for _, product := range products {
		productsDto = append(productsDto, *dto.MapToProductDto(product))
	}

	return productsDto, nil
}
