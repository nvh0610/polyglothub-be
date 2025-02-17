package user

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	customStatus "learn/internal/common/error"
	"learn/internal/repository"
	"learn/internal/router/payload/request"
	"learn/pkg/resp"
	"learn/pkg/utils"
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
	user, err := u.repo.User().GetById(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.USER_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, user)
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req *request.CreateUserRequest
	if err := utils.BindAndValidate(r, &req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, nil)
		return
	}

	input := ToModelCreateEntity(req)
	err := u.repo.User().Create(input)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}
