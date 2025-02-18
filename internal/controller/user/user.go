package user

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	customStatus "learn/internal/common/error"
	"learn/internal/repository"
	"learn/internal/router/payload/request"
	"learn/internal/router/payload/response"
	"learn/pkg/password"
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

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, response.ToDetailUserResponse(&user))
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := &request.CreateUserRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	isValidRole := IsValidRole(req.Role)
	if !isValidRole {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, nil)
		return
	}

	userExist, err := u.repo.User().CheckExistsByUsername(req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if userExist {
		resp.Return(w, http.StatusBadRequest, customStatus.USER_EXIST, nil)
		return
	}

	pass, err := password.HashPassword(req.Password)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, nil)
		return
	}
	req.Password = pass

	input := ToModelCreateEntity(req)
	err = u.repo.User().Create(input)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.CREATE_USER_FAILED, nil)
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	req := &request.UpdateUserRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	user, err := u.repo.User().GetById(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.USER_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	input := ToModelUpdateEntity(req, &user)
	err = u.repo.User().Update(input)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.UPDATE_USER_FAILED, nil)
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	err := u.repo.User().Delete(idInt)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (u *UserController) ListUser(w http.ResponseWriter, r *http.Request) {
	page, limit := utils.SetDefaultPagination(r.URL.Query())
	offset := (page - 1) * limit

	users, total, err := u.repo.User().List(limit, offset)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	data := response.ListUserResponse{
		Users: response.ToListUserResponse(users),
		PaginationResponse: response.PaginationResponse{
			TotalPage: utils.CalculatorTotalPage(total, limit),
			Limit:     limit,
			Page:      page,
		},
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, data)
}
