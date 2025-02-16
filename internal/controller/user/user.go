package user

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	customStatus "learn/internal/common/error"
	"learn/internal/repository"
	"learn/pkg/resp"
	"net/http"
	"strconv"
)

type UserController struct {
	repo repository.Registry
}

func NewUserController(userRepo repository.Registry) Controller {
	return &UserController{
		repo: userRepo,
	}
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	user, err := u.repo.User().GetUserById(idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			resp.Return(w, http.StatusNotFound, customStatus.USER_NOT_FOUND, "user not found")
		}
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, user)
}
