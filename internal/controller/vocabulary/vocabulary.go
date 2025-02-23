package vocabulary

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	customStatus "learn/internal/common/error"
	"learn/internal/controller/example"
	"learn/internal/controller/user"
	"learn/internal/repository"
	"learn/internal/router/payload/request"
	"learn/internal/router/payload/response"
	"learn/pkg/resp"
	"learn/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

type VocabularyController struct {
	repo repository.Registry
}

func NewVocabularyController(vocabularyRepo repository.Registry) Controller {
	return &VocabularyController{
		repo: vocabularyRepo,
	}
}

func (v *VocabularyController) CreateVocabulary(w http.ResponseWriter, r *http.Request) {
	req := &request.CreateVocabularyRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	userId, role := utils.GetUserIdAndRoleFromContext(r)

	category, err := v.repo.Category().GetById(req.CategoryId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.CATEGORY_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if !user.IsValidAdminRole(role) && category.UserID != userId {
		resp.Return(w, http.StatusForbidden, customStatus.FORBIDDEN, nil)
		return
	}

	exist, err := v.repo.Vocabulary().CheckExistsByWord(strings.ToLower(req.Word), req.CategoryId)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if !exist {
		resp.Return(w, http.StatusBadRequest, customStatus.VOCABULARY_ALREADY_EXISTS, nil)
		return
	}

	err = v.repo.DoInTx(func(txRepo repository.Registry) error {
		inputVocabulary := ToModelCreateVocabularyEntity(req)
		err = v.repo.Vocabulary().Create(inputVocabulary)
		if err != nil {
			return err
		}

		inputExample := example.ToModelExampleEntities(req.Examples, inputVocabulary.Id)

		err = v.repo.Example().CreateBatch(inputExample)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (v *VocabularyController) UpdateVocabulary(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	req := &request.UpdateVocabularyRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	userId, role := utils.GetUserIdAndRoleFromContext(r)

	vocabulary, err := v.repo.Vocabulary().GetById(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.VOCABULARY_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	category, err := v.repo.Category().GetById(vocabulary.CategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.CATEGORY_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if !user.IsValidAdminRole(role) && category.UserID != userId {
		resp.Return(w, http.StatusForbidden, customStatus.FORBIDDEN, nil)
		return
	}

	err = v.repo.DoInTx(func(txRepo repository.Registry) error {
		inputVocabulary := ToModelUpdateVocabularyEntity(req, vocabulary)
		inputVocabulary.Id = idInt
		err = v.repo.Vocabulary().Update(inputVocabulary)
		if err != nil {
			return err
		}

		inputExample := example.ToModelExampleEntities(req.Examples, inputVocabulary.Id)

		err = v.repo.Example().UpsertBatch(inputExample)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (v *VocabularyController) DeleteVocabulary(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	userId, role := utils.GetUserIdAndRoleFromContext(r)

	vocabulary, err := v.repo.Vocabulary().GetById(idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.VOCABULARY_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	category, err := v.repo.Category().GetById(vocabulary.CategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.CATEGORY_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if !user.IsValidAdminRole(role) && category.UserID != userId {
		resp.Return(w, http.StatusForbidden, customStatus.FORBIDDEN, nil)
		return
	}

	err = v.repo.DoInTx(func(txRepo repository.Registry) error {
		err = v.repo.Vocabulary().Delete(idInt)
		if err != nil {
			return err
		}

		err = v.repo.Example().DeleteByVocabularyId(idInt)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (v *VocabularyController) ListVocabulary(w http.ResponseWriter, r *http.Request) {
	page, limit := utils.SetDefaultPagination(r.URL.Query())
	userId, _ := utils.GetUserIdAndRoleFromContext(r)
	categoryId, _ := strconv.Atoi(r.URL.Query().Get("category_id"))
	word := r.URL.Query().Get("word")
	offset := (page - 1) * limit

	category, err := v.repo.Category().GetById(categoryId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.CATEGORY_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if category.UserID != 0 && category.UserID != userId {
		resp.Return(w, http.StatusForbidden, customStatus.FORBIDDEN, nil)
		return
	}

	vocabularies, total, err := v.repo.Vocabulary().List(limit, offset, categoryId, word)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	data := response.ListVocabularyResponse{
		Vocabularies: response.ToListVocabularyResponse(vocabularies),
		PaginationResponse: response.PaginationResponse{
			TotalPage: utils.CalculatorTotalPage(total, limit),
			Limit:     limit,
			Page:      page,
		},
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, data)
}
