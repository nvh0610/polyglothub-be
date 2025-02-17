package entity

import "time"

type Vocabulary struct {
	Id        int       `json:"id" gorm:"id"`
	Word      string    `json:"word" gorm:"word"`
	Meaning   string    `json:"meaning" gorm:"meaning"`
	IPA       string    `json:"ipa" gorm:"ipa"`
	Type      string    `json:"type" gorm:"type"`
	Url       string    `json:"url" gorm:"url"`
	Category  int       `json:"category" gorm:"category"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (Vocabulary) TableName() string {
	return "vocabularies"
}
