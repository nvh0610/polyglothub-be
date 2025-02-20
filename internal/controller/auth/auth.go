package auth

import (
	"errors"
	"gorm.io/gorm"
	customStatus "learn/internal/common/error"
	"learn/internal/repository"
	"learn/internal/router/payload/request"
	"learn/internal/router/payload/response"
	"learn/pkg/jwt"
	"learn/pkg/password"
	"learn/pkg/resp"
	"learn/pkg/utils"
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

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req := &request.LoginUserRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	user, err := a.repo.User().GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.USER_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if password.CheckPassword(req.Password, user.Password) != nil {
		resp.Return(w, http.StatusUnauthorized, customStatus.WRONG_PASSWORD, nil)
		return
	}

	token, err := jwt.GenerateToken(user.Id, user.Role)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, response.LoginResponse{
		AccessToken: token,
	})
}

func (a *AuthController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(int)

	req := &request.ChangePasswordRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	user, err := a.repo.User().GetById(userId)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if password.CheckPassword(req.OldPassword, user.Password) != nil {
		resp.Return(w, http.StatusUnauthorized, customStatus.WRONG_PASSWORD, nil)
		return
	}

	hashedPassword, err := password.HashPassword(req.NewPassword)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	user.Password = hashedPassword
	err = a.repo.User().Update(user)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}
