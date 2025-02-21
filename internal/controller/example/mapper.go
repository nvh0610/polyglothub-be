package example

import (
	"learn/internal/entity"
	"learn/internal/router/payload/request"
)

func ToModelExampleEntity(request request.ExampleRequest, vocabularyId int) *entity.Example {
	if request.VocabularyId == 0 {
		request.VocabularyId = vocabularyId
	}

	return &entity.Example{
		Id:           request.Id,
		VocabularyID: request.VocabularyId,
		Sentence:     request.Sentence,
		Meaning:      request.Meaning,
	}
}

func ToModelExampleEntities(requests []request.ExampleRequest, vocabularyId int) []*entity.Example {
	var examples []*entity.Example
	for _, input := range requests {
		examples = append(examples, ToModelExampleEntity(input, vocabularyId))
	}
	return examples
}
