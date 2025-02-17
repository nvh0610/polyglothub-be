package entity

import "time"

type Example struct {
	Id           int       `json:"id" gorm:"id"`
	VocabularyID int       `json:"vocabulary_id" gorm:"vocabulary_id"`
	Sentence     string    `json:"sentence" gorm:"sentence"`
	Meaning      string    `json:"meaning" gorm:"meaning"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"updated_at"`
}

func (Example) TableName() string {
	return "examples"
}
