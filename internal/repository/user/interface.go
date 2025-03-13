package user

import "learn/internal/entity"

type Repository interface {
	GetById(id int) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id int) error
	CheckExistsByUsername(username string) (bool, error)
	List(limit, offset int) ([]*entity.User, int, error)
	GetByUsername(username string) (*entity.User, error)
}
