package entity

import "time"

type Vocabulary struct {
	Id          int       `json:"id" gorm:"id"`
	Word        string    `json:"word" gorm:"word"`
	Meaning     string    `json:"meaning" gorm:"meaning"`
	Ipa         string    `json:"ipa" gorm:"ipa"`
	Type        string    `json:"type" gorm:"type"`
	Url         string    `json:"url" gorm:"url"`
	Description string    `json:"description" gorm:"description"`
	CategoryID  int       `json:"category_id" gorm:"category_id"`
	Topic       string    `json:"topic" gorm:"topic"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at;autoUpdateTime"`
}

func (Vocabulary) TableName() string {
	return "vocabularies"
}
