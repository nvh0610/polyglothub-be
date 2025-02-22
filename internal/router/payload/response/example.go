package response

import (
	"learn/internal/repository/vocabulary"
)

type ExampleResponse struct {
	Id           int    `json:"id"`
	VocabularyId int    `json:"vocabulary_id"`
	Sentence     string `json:"sentence"`
	Meaning      string `json:"meaning"`
}

type ExamplesResponse struct {
	Examples []ExampleResponse `json:"examples"`
}

func ToExampleResponse(example *vocabulary.Example) *ExampleResponse {
	return &ExampleResponse{
		Id:           example.Id,
		VocabularyId: example.VocabularyID,
		Sentence:     example.Sentence,
		Meaning:      example.Meaning,
	}
}

func ToExamplesResponse(examples []*vocabulary.Example) []*ExampleResponse {
	var responses []*ExampleResponse
	for _, example := range examples {
		responses = append(responses, ToExampleResponse(example))
	}
	return responses
}
