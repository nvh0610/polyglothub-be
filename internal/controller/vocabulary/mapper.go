package vocabulary

import (
	"learn/internal/entity"
	"learn/internal/router/payload/request"
	"strings"
)

func ToModelCreateVocabularyEntity(vocabulary *request.CreateVocabularyRequest) *entity.Vocabulary {
	return &entity.Vocabulary{
		Word:        strings.ToLower(vocabulary.Word),
		Meaning:     vocabulary.Meaning,
		Ipa:         vocabulary.IPA,
		Type:        vocabulary.Type,
		Url:         vocabulary.Url,
		CategoryID:  vocabulary.CategoryId,
		Description: vocabulary.Description,
	}

}

func ToModelUpdateVocabularyEntity(vocabulary *request.UpdateVocabularyRequest, vocabularyEntity *entity.Vocabulary) *entity.Vocabulary {
	vocabularyEntity.Word = strings.ToLower(vocabulary.Word)
	vocabularyEntity.Meaning = vocabulary.Meaning
	vocabularyEntity.Ipa = vocabulary.IPA
	vocabularyEntity.Type = vocabulary.Type
	vocabularyEntity.Url = vocabulary.Url
	vocabularyEntity.CategoryID = vocabulary.CategoryId
	vocabularyEntity.Description = vocabulary.Description
	return vocabularyEntity
}
