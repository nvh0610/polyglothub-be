package entity

import "time"

type UserFlashcardLog struct {
	Id           int       `json:"id" gorm:"id"`
	UserID       int       `json:"user_id" gorm:"user_id"`
	VocabularyID int       `json:"vocabulary_id" gorm:"vocabulary_id"`
	Answer       string    `json:"answer" gorm:"answer"`
	IsCorrect    bool      `json:"is_correct" gorm:"is_correct"`
	Date         time.Time `json:"date" gorm:"date"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"updated_at"`
}

func (UserFlashcardLog) TableName() string {
	return "user_flashcard_logs"
}
