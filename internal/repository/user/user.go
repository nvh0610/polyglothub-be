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

func (u *Implement) GetById(id int) (entity.User, error) {
	var user entity.User
	return user, u.db.First(&user, "id = ?", id).Error
}

func (u *Implement) List(limit, offset int) ([]entity.User, error) {
	var users []entity.User
	return users, u.db.Limit(limit).Offset(offset).Find(&users).Error
}

func (u *Implement) Create(user *entity.User) error {
	return u.db.Create(user).Error
}

func (u *Implement) Update(user *entity.User) error {
	return u.db.Save(user).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.User{Id: id}).Error
}
