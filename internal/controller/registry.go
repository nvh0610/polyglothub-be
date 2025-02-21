package controller

import (
	"github.com/redis/go-redis/v9"
	"learn/internal/controller/auth"
	"learn/internal/controller/category"
	"learn/internal/controller/user"
	"learn/internal/repository"
)

type RegistryController struct {
	UserCtrl     user.Controller
	AuthCtrl     auth.Controller
	CategoryCtrl category.Controller
}

func NewRegistryController(repo repository.Registry, redis *redis.Client) *RegistryController {
	return &RegistryController{
		UserCtrl:     user.NewUserController(repo),
		AuthCtrl:     auth.NewAuthController(repo, redis),
		CategoryCtrl: category.NewUserController(repo),
	}
}
