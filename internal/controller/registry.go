package controller

import (
	"learn/internal/controller/auth"
	"learn/internal/controller/user"
	"learn/internal/repository"
)

type RegistryController struct {
	UserCtrl user.Controller
	AuthCtrl auth.Controller
}

func NewRegistryController(repo repository.Registry) *RegistryController {
	return &RegistryController{
		UserCtrl: user.NewUserController(repo),
		AuthCtrl: auth.NewAuthController(repo),
	}
}
