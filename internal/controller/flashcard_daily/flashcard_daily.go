package flashcard_daily

import (
	"errors"
	"gorm.io/gorm"
	customStatus "learn/internal/common/error"
	"learn/internal/entity"
	"learn/internal/repository"
	"learn/internal/router/payload/request"
	"learn/internal/router/payload/response"
	"learn/pkg/logger"
	"learn/pkg/resp"
	"learn/pkg/utils"
	"net/http"
	"strings"
	"time"
)

type FlashcardDailyController struct {
	repo repository.Registry
}

func NewFlashcardDailyController(flashcardDailyRepo repository.Registry) Controller {
	return &FlashcardDailyController{
		repo: flashcardDailyRepo,
	}
}

func (f *FlashcardDailyController) GetFlashCardDaily(w http.ResponseWriter, r *http.Request) {
	userId, _ := utils.GetUserIdAndRoleFromContext(r)
	dateNow := r.URL.Query().Get("date")
	if dateNow == "" {
		dateNow = time.Now().Format("2006-01-02")
	} else {
		if _, err := time.Parse("2006-01-02", dateNow); err != nil {
			dateNow = time.Now().Format("2006-01-02")
		}
	}

	flashCards, err := f.repo.FlashCardDaily().GetFlashcardDaily(dateNow)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	vocabularyIds := make([]int, 0)
	for _, flashCard := range flashCards {
		_, err = f.repo.UserFlashCardLog().GetByUserIdAndDateAndVocabularyIdAndIsCorrect(userId, dateNow, flashCard.VocabularyId, true)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				vocabularyIds = append(vocabularyIds, flashCard.VocabularyId)
			} else {
				resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
			}
		}
	}

	vocabularies, err := f.repo.Vocabulary().GetVocabulariesByIds(vocabularyIds)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, response.ToFlashcardDailyResponse(vocabularies))
}

func (f *FlashcardDailyController) ConfirmFlashCardDaily(w http.ResponseWriter, r *http.Request) {
	userId, _ := utils.GetUserIdAndRoleFromContext(r)
	req := &request.ConfirmFlashCardDailyRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	dateNow := time.Now().Format("2006-01-02")
	if _, err := f.repo.FlashCardDaily().GetByVocabularyIdAndDate(req.VocabularyId, dateNow); err != nil {
		handleError(w, err, customStatus.FLASHCARD_DAILY_NOT_FOUND)
		return
	}

	vocabulary, err := f.repo.Vocabulary().GetById(req.VocabularyId)
	if err != nil {
		handleError(w, err, customStatus.VOCABULARY_NOT_FOUND)
		return
	}

	isCorrect := strings.EqualFold(vocabulary.Word, req.Answer)

	flashCardLog, err := f.repo.UserFlashCardLog().GetByUserIdAndDateAndVocabularyIdAndIsCorrect(userId, dateNow, req.VocabularyId, true)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
			return
		}
	}

	if flashCardLog.Id != 0 {
		resp.Return(w, http.StatusBadRequest, customStatus.FLASHCARD_LOG_EXIST, nil)
		return
	}

	go f.processFlashcardLog(userId, req.VocabularyId, req.Answer, isCorrect, dateNow)

	if !isCorrect {
		resp.Return(w, http.StatusBadRequest, customStatus.WRONG_ANSWER, nil)
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (f *FlashcardDailyController) processFlashcardLog(userId, vocabularyId int, answer string, isCorrect bool, dateNow string) {
	err := f.repo.UserFlashCardLog().Create(&entity.UserFlashcardLog{
		UserID:       userId,
		VocabularyID: vocabularyId,
		Answer:       answer,
		IsCorrect:    isCorrect,
		Date:         getTodayDate(),
	})
	if err != nil {
		logError(err)
		return
	}

	dailyWord, err := f.repo.UserDailyWordStatistics().GetByUserIdAndDate(userId, dateNow)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = f.repo.UserDailyWordStatistics().Create(&entity.UserDailyWordStatistics{
				UserID:         userId,
				Date:           getTodayDate(),
				CorrectAnswers: boolToInt(isCorrect),
				WrongAnswers:   boolToInt(!isCorrect),
			})
		}
	} else {
		dailyWord.CorrectAnswers += boolToInt(isCorrect)
		dailyWord.WrongAnswers += boolToInt(!isCorrect)
		err = f.repo.UserDailyWordStatistics().Update(dailyWord)
	}

	if err != nil {
		logError(err)
	}
}

func handleError(w http.ResponseWriter, err error, notFoundStatus int) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Return(w, http.StatusNotFound, notFoundStatus, nil)
	} else {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
	}
}

func getTodayDate() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
}

func boolToInt(val bool) int {
	if val {
		return 1
	}
	return 0
}

func logError(err error) {
	logger.ErrorF("[FlashcardDailyController] Error: %v", err)
}
