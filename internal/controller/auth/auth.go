package auth

import (
	"learn/internal/repository"
	"net/http"
)

type AuthController struct {
	repo repository.Registry
}

func NewAuthController(userRepo repository.Registry) Controller {
	return &AuthController{
		repo: userRepo,
	}
}

func (a AuthController) Login(w http.ResponseWriter, r *http.Request) {
}
