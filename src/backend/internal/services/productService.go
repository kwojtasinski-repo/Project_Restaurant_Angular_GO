package services

import (
	"fmt"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"
	applicationerrors "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/errors"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
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

	productDto.Normalize()

	category, errorRepo := service.categoryRepository.Get(productDto.CategoryId)
	if errorRepo != nil {
		return nil, applicationerrors.InternalError(errorRepo.Error())
	}

	if category == nil {
		return nil, applicationerrors.BadRequest(fmt.Sprintf("'Category' with id %v was not found", productDto.CategoryId))
	}

	product := &entities.Product{
		Name:        productDto.Name,
		Price:       productDto.Price,
		Category:    *category,
		Description: productDto.Description,
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

	productDto.Normalize()

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
		return nil, applicationerrors.BadRequest(fmt.Sprintf("'Product' with id %v was not found", productDto.Id))
	}

	product.Name = productDto.Name
	product.Description = productDto.Description
	product.Price = productDto.Price
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
		return applicationerrors.BadRequest(fmt.Sprintf("'Product' with id %v was not found", id))
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
