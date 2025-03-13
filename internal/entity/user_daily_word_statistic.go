package entity

import "time"

type UserDailyWordStatistics struct {
	Id             int       `json:"id" gorm:"id"`
	UserID         int       `json:"user_id" gorm:"user_id"`
	CorrectAnswers int       `json:"correct_answers" gorm:"correct_answers"`
	WrongAnswers   int       `json:"wrong_answers" gorm:"wrong_answers"`
	Date           time.Time `json:"date" gorm:"date"`
	CreatedAt      time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"updated_at"`
}

func (UserDailyWordStatistics) TableName() string {
	return "user_daily_word_statistics"
}
