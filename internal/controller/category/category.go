package category

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	customStatus "learn/internal/common/error"
	"learn/internal/controller/user"
	"learn/internal/repository"
	"learn/internal/router/payload/request"
	"learn/internal/router/payload/response"
	"learn/pkg/resp"
	"learn/pkg/utils"
	"net/http"
	"strconv"
)

type CategoryController struct {
	repo repository.Registry
}

func NewCategoryController(categoryRepo repository.Registry) Controller {
	return &CategoryController{
		repo: categoryRepo,
	}
}

func (c *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	req := &request.CreateCategoryRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}
	input := ToModelCategoryEntity(req)
	userId, role := utils.GetUserIdAndRoleFromContext(r)
	if !user.IsValidAdminRole(role) {
		input.UserID = userId
	}

	err := c.repo.Category().Create(input)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.CREATE_CATEGORY_FAILED, nil)
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (c *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	userId, role := utils.GetUserIdAndRoleFromContext(r)

	req := &request.UpdateCategoryRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	category, err := c.repo.Category().GetById(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.CATEGORY_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if category.UserID != userId && !user.IsValidAdminRole(role) {
		resp.Return(w, http.StatusForbidden, customStatus.FORBIDDEN, nil)
		return
	}

	category.Name = req.Name
	err = c.repo.Category().Update(category)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.UPDATE_CATEGORY_FAILED, nil)
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (c *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	userId, role := utils.GetUserIdAndRoleFromContext(r)

	category, err := c.repo.Category().GetById(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.CATEGORY_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if category.UserID != userId && !user.IsValidAdminRole(role) {
		resp.Return(w, http.StatusForbidden, customStatus.FORBIDDEN, nil)
		return
	}

	err = c.repo.Category().Delete(idInt)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (c *CategoryController) ListCategory(w http.ResponseWriter, r *http.Request) {
	page, limit := utils.SetDefaultPagination(r.URL.Query())
	offset := (page - 1) * limit
	userID := r.Context().Value("user_id").(int)

	categories, total, err := c.repo.Category().List(limit, offset, userID)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	data := response.ListCategoryResponse{
		Categories: response.ToListCategoryResponse(categories),
		PaginationResponse: response.PaginationResponse{
			TotalPage: utils.CalculatorTotalPage(total, limit),
			Limit:     limit,
			Page:      page,
		},
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, data)
}
