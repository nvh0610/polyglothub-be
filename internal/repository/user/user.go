package user

import (
	"gorm.io/gorm"
	"learn/internal/entity"
)

type Implement struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) GetUserById(id int) (entity.User, error) {
	var user entity.User
	return user, u.db.First(&user, "id = ?", id).Error
}
