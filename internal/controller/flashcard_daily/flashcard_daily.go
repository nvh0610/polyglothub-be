package flashcard_daily

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	customStatus "learn/internal/common/error"
	"learn/internal/entity"
	"learn/internal/repository"
	"learn/internal/router/payload/request"
	"learn/internal/router/payload/response"
	"learn/job/schedule"
	"learn/pkg/config"
	"learn/pkg/logger"
	"learn/pkg/resp"
	"learn/pkg/utils"
	"net/http"
	"strings"
	"time"
)

const (
	MaxFlashCardDaily              = 5
	TimeCronJobFetchFlashCardDaily = "@every 0h3m0s"
)

type FlashcardDailyController struct {
	repo           repository.Registry
	vocabularyRepo repository.Registry
	redis          *redis.Client
	maxFlashCard   int
	timeCronJob    string
}

func NewFlashcardDailyController(repo repository.Registry, redis *redis.Client) Controller {
	return &FlashcardDailyController{
		repo:           repo,
		vocabularyRepo: repo,
		redis:          redis,
		maxFlashCard:   config.MaxFlashCardDaily(),
		timeCronJob:    config.CronJobFetchFlashCardDaily(),
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
	switch req.Type {
	case "word":
		isCorrect = strings.EqualFold(vocabulary.Meaning, req.Answer)
		req.Answer = vocabulary.Word
	case "meaning":
		isCorrect = strings.EqualFold(vocabulary.Word, req.Answer)
	case "ipa":
		isCorrect = strings.EqualFold(vocabulary.Word, req.Answer)
	case "audio":
		isCorrect = strings.EqualFold(vocabulary.Word, req.Answer)
	}

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

	//if !isCorrect {
	//	resp.Return(w, http.StatusBadRequest, customStatus.WRONG_ANSWER, nil)
	//	return
	//}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (f *FlashcardDailyController) GetAllFlashCard(w http.ResponseWriter, r *http.Request) {
	userId, _ := utils.GetUserIdAndRoleFromContext(r)
	page, limit := utils.SetDefaultPagination(r.URL.Query())
	offset := (page - 1) * limit
	var categoryIds []int
	option := r.URL.Query().Get("option")
	if option == "private" {
		categories, err := f.repo.Category().GetByUserId(userId)
		if err != nil {
			resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
			return
		}

		for _, category := range categories {
			categoryIds = append(categoryIds, category.Id)
		}
	}

	vocabularies, total, err := f.repo.Vocabulary().List(limit, offset, 0, "", categoryIds, "")
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

func (f *FlashcardDailyController) CronJobDailyFlashcard() {
	_, _ = schedule.RegisterScheduler(config.CronJobFetchFlashCardDaily(), func() {
		logger.Info("cron job: fetch flashcard daily")
		dateNow := time.Now().Format("2006-01-02")
		flashCards, err := f.repo.FlashCardDaily().GetFlashcardDaily(dateNow)
		if err != nil {
			logError(err)
			return
		}

		if len(flashCards) >= f.maxFlashCard {
			return
		}

		maxId, err := f.repo.Vocabulary().GetMaxId()
		if err != nil {
			logError(err)
			return
		}

		usedIDs := f.redis.LRange(context.Background(), config.REDIS_FLASHCARD_DAILY, 0, -1).Val()
		var vocabularies []*entity.Vocabulary
		var allVocabIds []int

		for len(vocabularies) < f.maxFlashCard {
			remainingCount := f.maxFlashCard - len(vocabularies)

			randomIDs := utils.GenerateRandomNumbers(remainingCount, maxId, usedIDs)
			if len(randomIDs) == 0 {
				for i := 0; i < remainingCount; i++ {
					oldID, err := f.redis.RPop(context.Background(), config.REDIS_FLASHCARD_DAILY).Int()
					if err != nil {
						break
					}
					randomIDs = append(randomIDs, oldID)
				}
				if len(randomIDs) == 0 {
					break
				}
			}

			newVocabs, err := f.repo.Vocabulary().GetVocabulariesByIds(randomIDs)
			if err != nil {
				logError(err)
				return
			}

			vocabularies = append(vocabularies, newVocabs...)

			for _, vocab := range newVocabs {
				allVocabIds = append(allVocabIds, vocab.Id)
			}

			usedIDs = append(usedIDs, utils.IntSliceToStringSlice(randomIDs)...)
		}

		if len(vocabularies) == 0 {
			logError(errors.New("no vocabularies found"))
			return
		}

		entityFlashCards := make([]*entity.FlashcardDaily, len(vocabularies))
		today := getTodayDate()

		for i, vocabulary := range vocabularies {
			entityFlashCards[i] = &entity.FlashcardDaily{
				VocabularyId: vocabulary.Id,
				Date:         today,
			}
		}

		err = f.repo.FlashCardDaily().CreateBatch(entityFlashCards)
		if err != nil {
			logError(err)
			return
		}

		err = f.redis.LPush(context.Background(), config.REDIS_FLASHCARD_DAILY, utils.IntSliceToStringSlice(allVocabIds)).Err()
		if err != nil {
			logError(err)
			return
		}
	})
}

func (f *FlashcardDailyController) GetDashboard(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	page, limit := utils.SetDefaultPagination(r.URL.Query())
	offset := (page - 1) * limit
	if startDate == "" {
		startDate = time.Now().Format("2006-01-02")
	}
	if endDate == "" {
		endDate = time.Now().Format("2006-01-02")
	}

	userDailyWordStatistics, total, err := f.repo.UserDailyWordStatistics().GetByDate(startDate, endDate, limit, offset)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	data := response.DashboardFlashcardsResponse{
		Report: response.ToDashboardFlashcardResponse(userDailyWordStatistics),
		PaginationResponse: response.PaginationResponse{
			TotalPage: utils.CalculatorTotalPage(total, limit),
			Limit:     limit,
			Page:      page,
		},
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, data)
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
