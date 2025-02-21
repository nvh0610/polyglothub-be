package repository

import (
	"gorm.io/gorm"
	"learn/internal/repository/category"
	"learn/internal/repository/user"
)

type Registry interface {
	User() user.Repository
	Category() category.Repository
}

type mysqlImplement struct {
	userRepo user.Repository
	category category.Repository
}

func (m *mysqlImplement) User() user.Repository {
	return m.userRepo
}

func (m *mysqlImplement) Category() category.Repository {
	return m.category
}

func NewRegistryRepo(db *gorm.DB) Registry {
	return &mysqlImplement{
		userRepo: user.NewUserRepository(db),
		category: category.NewCategoryRepository(db),
	}
}
