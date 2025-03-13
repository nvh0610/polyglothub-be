package entity

import "time"

const (
	USER_ROLE_ADMIN      = "admin"
	USER_ROLE_USER       = "user"
	USER_ROLE_SUPERADMIN = "superadmin"
)

type User struct {
	Id        int       `json:"id" gorm:"id"`
	Username  string    `json:"username" gorm:"username"`
	Password  string    `json:"password" gorm:"password"`
	Fullname  string    `json:"fullname" gorm:"fullname"`
	Role      string    `json:"role" gorm:"role"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
