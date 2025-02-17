package user

import (
	"learn/internal/entity"
	"learn/internal/router/payload/request"
)

func ToModelCreateEntity(user *request.CreateUserRequest) *entity.User {
	return &entity.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
}
