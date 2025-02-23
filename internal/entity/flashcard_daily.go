package entity

import "time"

type FlashcardDaily struct {
	Id           int       `json:"id" gorm:"id"`
	VocabularyId int       `json:"vocabulary_id" gorm:"vocabulary_id"`
	Topic        string    `json:"topic" gorm:"topic"`
	Date         time.Time `json:"date" gorm:"date"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"updated_at"`
}

func (FlashcardDaily) TableName() string {
	return "flashcard_dailies"
}
