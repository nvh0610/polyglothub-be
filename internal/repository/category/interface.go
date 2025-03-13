package category

import "learn/internal/entity"

type Repository interface {
	GetById(id int) (*entity.Category, error)
	GetByUserId(userId int) ([]*entity.Category, error)
	Create(user *entity.Category) error
	Update(user *entity.Category) error
	Delete(id int) error
	List(limit, offset int, userId int) ([]*entity.Category, int, error)
}
