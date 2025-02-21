package request

type ExampleRequest struct {
	VocabularyId int    `json:"vocabulary_id"`
	Sentence     string `json:"sentence" validate:"required"`
	Meaning      string `json:"meaning" validate:"required"`
}
