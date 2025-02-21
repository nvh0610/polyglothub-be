package vocabulary

import "learn/internal/repository"

type VocabularyController struct {
	repo repository.Registry
}

func NewVocabularyController(vocabularyRepo repository.Registry) Controller {
	return &VocabularyController{
		repo: vocabularyRepo,
	}
}
