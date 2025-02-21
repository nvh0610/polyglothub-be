package vocabulary

import (
	"learn/internal/entity"
	"learn/internal/router/payload/request"
)

func ToModelCreateVocabularyEntity(vocabulary *request.CreateVocabularyRequest) *entity.Vocabulary {
	return &entity.Vocabulary{
		Word:       vocabulary.Word,
		Meaning:    vocabulary.Meaning,
		IPA:        vocabulary.IPA,
		Type:       vocabulary.Type,
		Url:        vocabulary.Url,
		CategoryID: vocabulary.CategoryId,
	}

}

func ToModelUpdateVocabularyEntity(vocabulary *request.UpdateVocabularyRequest) *entity.Vocabulary {
	return &entity.Vocabulary{
		Word:       vocabulary.Word,
		Meaning:    vocabulary.Meaning,
		IPA:        vocabulary.IPA,
		Type:       vocabulary.Type,
		Url:        vocabulary.Url,
		CategoryID: vocabulary.CategoryId,
	}
}
