package user_flashcard_log

import (
	"gorm.io/gorm"
	"learn/internal/entity"
)

type Implement struct {
	db *gorm.DB
}

func NewUserFlashcardLogRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) Create(userFlashcardLog *entity.UserFlashcardLog) error {
	return u.db.Create(userFlashcardLog).Error
}

func (u *Implement) GetByUserIdAndDateAndVocabularyIdAndIsCorrect(userId int, date string, vocabularyId int, isCorrect bool) (*entity.UserFlashcardLog, error) {
	var userFlashcardLog *entity.UserFlashcardLog
	return userFlashcardLog, u.db.First(&userFlashcardLog, "user_id = ? AND date = ? AND vocabulary_id = ? AND is_correct = ?", userId, date, vocabularyId, isCorrect).Error
}
