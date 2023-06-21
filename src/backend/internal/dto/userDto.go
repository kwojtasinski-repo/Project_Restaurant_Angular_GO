package dto

import "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"

type UserDto struct {
	Id      IdObject `json:"id"`
	Email   string   `json:"email"`
	Role    string   `json:"role"`
	Deleted *bool    `json:"deleted"`
}

func MapToUserDto(user entities.User) *UserDto {
	var deleted *bool
	deleted = nil
	if user.Deleted {
		deleted = &user.Deleted
	}
	return &UserDto{
		Id:      IdObject{ValueInt: user.Id.Value()},
		Email:   user.Email.String(),
		Role:    user.Role,
		Deleted: deleted,
	}
}
