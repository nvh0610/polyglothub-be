package entity

import "time"

type Category struct {
	Id        int       `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	UserID    int       `json:"user_id" gorm:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}
