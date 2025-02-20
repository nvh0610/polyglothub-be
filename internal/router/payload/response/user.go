package response

import (
	"learn/internal/entity"
	"time"
)

type DetailUserResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDetailUserResponse(user *entity.User) *DetailUserResponse {
	return &DetailUserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListUserResponse(users []*entity.User) []*DetailUserResponse {
	var res []*DetailUserResponse
	for _, user := range users {
		res = append(res, ToDetailUserResponse(user))
	}
	return res
}

type PaginationResponse struct {
	TotalPage int `json:"total_page"`
	Limit     int `json:"limit"`
	Page      int `json:"page"`
}

type ListUserResponse struct {
	PaginationResponse
	Users []*DetailUserResponse `json:"users"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
