package dto

import "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"

type UserDto struct {
	Id    int64
	Email string
	Role  string
}

func MapToUserDto(user entities.User) *UserDto {
	return &UserDto{
		Id:    user.Id.Value(),
		Email: user.Email.String(),
		Role:  user.Role,
	}
}
