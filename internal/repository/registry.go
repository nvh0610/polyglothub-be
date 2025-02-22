package repository

import (
	"gorm.io/gorm"
	"learn/internal/repository/category"
	"learn/internal/repository/example"
	"learn/internal/repository/user"
	"learn/internal/repository/vocabulary"
)

type Registry interface {
	User() user.Repository
	Category() category.Repository
	Vocabulary() vocabulary.Repository
	Example() example.Repository
	DoInTx(txFunc func(txRepo Registry) error) error
}

type mysqlImplement struct {
	db         *gorm.DB
	userRepo   user.Repository
	category   category.Repository
	vocabulary vocabulary.Repository
	example    example.Repository
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

func (m *mysqlImplement) Example() example.Repository {
	return m.example
}

func NewRegistryRepo(db *gorm.DB) Registry {
	return &mysqlImplement{
		db:         db,
		userRepo:   user.NewUserRepository(db),
		category:   category.NewCategoryRepository(db),
		vocabulary: vocabulary.NewVocabularyRepository(db),
		example:    example.NewExampleRepository(db),
	}
}

func (m *mysqlImplement) DoInTx(txFunc func(txRepo Registry) error) error {
	tx := m.db.Begin()
	txRepo := &mysqlImplement{
		db:         m.db,
		userRepo:   user.NewUserRepository(tx),
		category:   category.NewCategoryRepository(tx),
		vocabulary: vocabulary.NewVocabularyRepository(tx),
		example:    example.NewExampleRepository(tx),
	}

	err := txFunc(txRepo)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
