package request

type CreateVocabularyRequest struct {
	Word        string           `json:"word" validate:"required"`
	Meaning     string           `json:"meaning" validate:"required"`
	IPA         string           `json:"ipa" validate:"required"`
	Type        string           `json:"type" validate:"required"`
	Url         string           `json:"url" validate:"required"`
	Description string           `json:"description" validate:"required"`
	CategoryId  int              `json:"category_id" validate:"required"`
	Examples    []ExampleRequest `json:"examples" validate:"required"`
}

type UpdateVocabularyRequest struct {
	Word        string           `json:"word" validate:"required"`
	Meaning     string           `json:"meaning" validate:"required"`
	IPA         string           `json:"ipa" validate:"required"`
	Type        string           `json:"type" validate:"required"`
	Url         string           `json:"url" validate:"required"`
	Description string           `json:"description" validate:"required"`
	CategoryId  int              `json:"category_id" validate:"required"`
	Examples    []ExampleRequest `json:"examples" validate:"required"`
}
