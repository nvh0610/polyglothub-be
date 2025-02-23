package flashcard_daily

import (
	"learn/internal/entity"
)

type Repository interface {
	CreateBatch(flashcardDailies []*entity.FlashcardDaily) error
	GetFlashcardDaily(date string) ([]*entity.FlashcardDaily, error)
	GetByVocabularyIdAndDate(vocabularyId int, date string) (*entity.FlashcardDaily, error)
}
