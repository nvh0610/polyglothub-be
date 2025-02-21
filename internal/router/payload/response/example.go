package response

import "learn/internal/entity"

type ExampleResponse struct {
	Id           int    `json:"id"`
	VocabularyId int    `json:"vocabulary_id"`
	Sentence     string `json:"sentence"`
	Meaning      string `json:"meaning"`
}

type ExamplesResponse struct {
	Examples []ExampleResponse `json:"examples"`
}

func ToExampleResponse(example *entity.Example) *ExampleResponse {
	return &ExampleResponse{
		Id:           example.Id,
		VocabularyId: example.VocabularyID,
		Sentence:     example.Sentence,
		Meaning:      example.Meaning,
	}
}

func ToExamplesResponse(examples []*entity.Example) []*ExampleResponse {
	var responses []*ExampleResponse
	for _, example := range examples {
		responses = append(responses, ToExampleResponse(example))
	}
	return responses
}
