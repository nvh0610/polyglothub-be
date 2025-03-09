package response

import (
	"learn/internal/entity"
	"learn/internal/repository/user_daily_word_statistic"
	"math"
)

type FlashcardDailyResponse struct {
	VocabularyId int    `json:"vocabulary_id"`
	Word         string `json:"word"`
	Meaning      string `json:"meaning"`
	IPA          string `json:"ipa"`
	Type         string `json:"type"`
	Url          string `json:"url"`
	Description  string `json:"description"`
}

type FlashcardDailiesResponse struct {
	Flashcards []*FlashcardDailyResponse `json:"flashcards"`
}

func ToFlashcardDailyResponse(vocabulary []*entity.Vocabulary) *FlashcardDailiesResponse {
	var flashcards []*FlashcardDailyResponse
	for _, item := range vocabulary {
		flashcards = append(flashcards, &FlashcardDailyResponse{
			VocabularyId: item.Id,
			Word:         item.Word,
			Meaning:      item.Meaning,
			IPA:          item.Ipa,
			Type:         item.Type,
			Url:          item.Url,
			Description:  item.Description,
		})
	}

	return &FlashcardDailiesResponse{
		Flashcards: flashcards,
	}
}

type DashboardFlashcardResponse struct {
	Username      string `json:"username"`
	NumberCorrect int    `json:"number_correct"`
	NumberWrong   int    `json:"number_wrong"`
	Percent       int    `json:"percent"`
}

type DashboardFlashcardsResponse struct {
	Report             []*DashboardFlashcardResponse `json:"report"`
	PaginationResponse PaginationResponse            `json:"pagination"`
}

func ToDashboardFlashcardResponse(dashboard []*user_daily_word_statistic.DashboardResponse) []*DashboardFlashcardResponse {
	var report []*DashboardFlashcardResponse
	for _, item := range dashboard {
		report = append(report, &DashboardFlashcardResponse{
			Username:      item.Fullname,
			NumberCorrect: item.CorrectAnswers,
			NumberWrong:   item.WrongAnswers,
			Percent:       int(math.Round(float64(item.CorrectAnswers) * 100 / float64(item.CorrectAnswers+item.WrongAnswers))),
		})
	}

	return report
}
