package user

import "learn/internal/entity"

type Repository interface {
	GetById(id int) (entity.User, error)
	List(limit, offset int) ([]entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id int) error
}
