package response

import "learn/internal/repository/vocabulary"

type VocabularyResponse struct {
	ID          int                `json:"id"`
	Word        string             `json:"word"`
	Meaning     string             `json:"meaning"`
	IPA         string             `json:"ipa"`
	Type        string             `json:"type"`
	Url         string             `json:"url"`
	Description string             `json:"description"`
	Examples    []*ExampleResponse `json:"examples"`
	CreatedAt   string             `json:"created_at"`
	UpdatedAt   string             `json:"updated_at"`
}

type ListVocabularyResponse struct {
	Vocabularies       []*VocabularyResponse `json:"vocabularies"`
	PaginationResponse PaginationResponse    `json:"pagination"`
}

func ToVocabularyResponse(vocabulary *vocabulary.Vocabularies) *VocabularyResponse {
	return &VocabularyResponse{
		ID:          vocabulary.Id,
		Word:        vocabulary.Word,
		Meaning:     vocabulary.Meaning,
		IPA:         vocabulary.IPA,
		Type:        vocabulary.Type,
		Url:         vocabulary.Url,
		Description: vocabulary.Description,
		Examples:    ToExamplesResponse(vocabulary.Examples),
		CreatedAt:   vocabulary.CreatedAt.Format("2006-01-02"),
		UpdatedAt:   vocabulary.UpdatedAt.Format("2006-01-02"),
	}
}

func ToListVocabularyResponse(vocabularies []*vocabulary.Vocabularies) []*VocabularyResponse {
	var response []*VocabularyResponse
	for _, item := range vocabularies {
		response = append(response, ToVocabularyResponse(item))
	}

	return response
}
