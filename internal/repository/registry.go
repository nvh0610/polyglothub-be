package repository

import (
	"gorm.io/gorm"
	"learn/internal/repository/category"
	"learn/internal/repository/example"
	"learn/internal/repository/flashcard_daily"
	"learn/internal/repository/user"
	"learn/internal/repository/user_daily_word_statistic"
	"learn/internal/repository/user_flashcard_log"
	"learn/internal/repository/vocabulary"
)

type Registry interface {
	User() user.Repository
	Category() category.Repository
	Vocabulary() vocabulary.Repository
	Example() example.Repository
	FlashCardDaily() flashcard_daily.Repository
	UserFlashCardLog() user_flashcard_log.Repository
	UserDailyWordStatistics() user_daily_word_statistic.Repository
	DoInTx(txFunc func(txRepo Registry) error) error
}

type mysqlImplement struct {
	db                      *gorm.DB
	userRepo                user.Repository
	category                category.Repository
	vocabulary              vocabulary.Repository
	example                 example.Repository
	flashCardDaily          flashcard_daily.Repository
	userFlashCardLog        user_flashcard_log.Repository
	userDailyWordStatistics user_daily_word_statistic.Repository
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

func (m *mysqlImplement) FlashCardDaily() flashcard_daily.Repository {
	return m.flashCardDaily
}

func (m *mysqlImplement) UserFlashCardLog() user_flashcard_log.Repository {
	return m.userFlashCardLog
}

func (m *mysqlImplement) UserDailyWordStatistics() user_daily_word_statistic.Repository {
	return m.userDailyWordStatistics
}

func NewRegistryRepo(db *gorm.DB) Registry {
	return &mysqlImplement{
		db:                      db,
		userRepo:                user.NewUserRepository(db),
		category:                category.NewCategoryRepository(db),
		vocabulary:              vocabulary.NewVocabularyRepository(db),
		example:                 example.NewExampleRepository(db),
		flashCardDaily:          flashcard_daily.NewFlashcardDailyRepository(db),
		userFlashCardLog:        user_flashcard_log.NewUserFlashcardLogRepository(db),
		userDailyWordStatistics: user_daily_word_statistic.NewUserDailyWordStatisticRepository(db),
	}
}

func (m *mysqlImplement) DoInTx(txFunc func(txRepo Registry) error) error {
	tx := m.db.Begin()
	txRepo := &mysqlImplement{
		db:                      m.db,
		userRepo:                user.NewUserRepository(tx),
		category:                category.NewCategoryRepository(tx),
		vocabulary:              vocabulary.NewVocabularyRepository(tx),
		example:                 example.NewExampleRepository(tx),
		flashCardDaily:          flashcard_daily.NewFlashcardDailyRepository(tx),
		userFlashCardLog:        user_flashcard_log.NewUserFlashcardLogRepository(tx),
		userDailyWordStatistics: user_daily_word_statistic.NewUserDailyWordStatisticRepository(tx),
	}

	err := txFunc(txRepo)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
