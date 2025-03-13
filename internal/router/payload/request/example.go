package request

type ExampleRequest struct {
	Id           int    `json:"id"`
	VocabularyId int    `json:"vocabulary_id"`
	Sentence     string `json:"sentence" validate:"required"`
	Meaning      string `json:"meaning" validate:"required"`
}
