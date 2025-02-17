package entity

import "time"

type UserFlashcardLogs struct {
	Id           int       `json:"id" gorm:"id"`
	UserID       int       `json:"user_id" gorm:"user_id"`
	VocabularyID int       `json:"vocabulary_id" gorm:"vocabulary_id"`
	Answer       string    `json:"answer" gorm:"answer"`
	IsCorrect    bool      `json:"is_correct" gorm:"is_correct"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"updated_at"`
}

func (UserFlashcardLogs) TableName() string {
	return "user_flashcard_logs"
}
