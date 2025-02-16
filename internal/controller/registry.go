package controller

import (
	"learn/internal/controller/user"
	"learn/internal/repository"
)

type RegistryController struct {
	UserCtrl user.Controller
}

func NewRegistryController(repo repository.Registry) *RegistryController {
	return &RegistryController{
		UserCtrl: user.NewUserController(repo),
	}
}
