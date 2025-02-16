package entity

type User struct {
	Id        int    `json:"id" gorm:"id"`
	Username  string `json:"username" gorm:"username"`
	Password  string `json:"password" gorm:"password"`
	Role      string `json:"role" gorm:"role"`
	CreatedAt string `json:"created_at" gorm:"created_at"`
	UpdatedAt string `json:"updated_at" gorm:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
