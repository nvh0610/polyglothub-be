package response

import "learn/internal/entity"

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
