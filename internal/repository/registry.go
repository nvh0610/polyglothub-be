package repository

import (
	"gorm.io/gorm"
	"learn/internal/repository/category"
	"learn/internal/repository/user"
	"learn/internal/repository/vocabulary"
)

type Registry interface {
	User() user.Repository
	Category() category.Repository
	Vocabulary() vocabulary.Repository
}

type mysqlImplement struct {
	userRepo   user.Repository
	category   category.Repository
	vocabulary vocabulary.Repository
}

func (m *mysqlImplement) User() user.Repository {
	return m.userRepo
}

func (m *mysqlImplement) Category() category.Repository {
	return m.category
}

func (m *mysqlImplement) Vocabulary() vocabulary.Repository {
	return m.vocabulary
}

func NewRegistryRepo(db *gorm.DB) Registry {
	return &mysqlImplement{
		userRepo:   user.NewUserRepository(db),
		category:   category.NewCategoryRepository(db),
		vocabulary: vocabulary.NewVocabularyRepository(db),
	}
}
