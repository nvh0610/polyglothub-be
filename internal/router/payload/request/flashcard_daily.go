package request

type ConfirmFlashCardDailyRequest struct {
	VocabularyId int    `json:"vocabulary_id" validate:"required"`
	Answer       string `json:"answer" validate:"required"`
	Type         string `json:"type" validate:"required"` //word, meaning, ipa, audio
}
