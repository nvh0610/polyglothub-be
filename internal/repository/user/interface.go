package user

import "learn/internal/entity"

type Repository interface {
	GetUserById(id int) (entity.User, error)
}
