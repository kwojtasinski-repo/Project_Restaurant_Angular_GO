package entities

import "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/dto"

type Category struct {
	Id      int64
	Name    string
	Deleted bool // soft delete
}

func (product *Category) MapToDto() *dto.CategoryDto {
	return &dto.CategoryDto{
		Id:   product.Id,
		Name: product.Name,
	}
}
