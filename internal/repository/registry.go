package repository

import (
	"gorm.io/gorm"
	"learn/internal/repository/user"
)

type Registry interface {
	User() user.Repository
}

type mysqlImplement struct {
	userRepo user.Repository
}

func (m *mysqlImplement) User() user.Repository {
	return m.userRepo
}

func NewRegistryRepo(db *gorm.DB) Registry {
	return &mysqlImplement{
		userRepo: user.NewUserRepository(db),
	}
}
