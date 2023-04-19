package dto

import "github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/entities"

type UserDto struct {
	Id      int64
	Email   string
	Role    string
	Deleted *bool
}

func MapToUserDto(user entities.User) *UserDto {
	var deleted *bool
	deleted = nil
	if user.Deleted {
		deleted = &user.Deleted
	}
	return &UserDto{
		Id:      user.Id.Value(),
		Email:   user.Email.String(),
		Role:    user.Role,
		Deleted: deleted,
	}
}
