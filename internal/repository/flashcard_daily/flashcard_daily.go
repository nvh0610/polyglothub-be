package flashcard_daily

import (
	"gorm.io/gorm"
	"learn/internal/entity"
)

type Implement struct {
	db *gorm.DB
}

func NewFlashcardDailyRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) CreateBatch(flashcardDailies []*entity.FlashcardDaily) error {
	return u.db.Create(flashcardDailies).Error
}

func (u *Implement) GetFlashcardDaily(date string) ([]*entity.FlashcardDaily, error) {
	var flashcardDailies []*entity.FlashcardDaily
	return flashcardDailies, u.db.Where("date = ?", date).Find(&flashcardDailies).Error
}

func (u *Implement) GetByVocabularyIdAndDate(vocabularyId int, date string) (*entity.FlashcardDaily, error) {
	var flashcardDaily *entity.FlashcardDaily
	return flashcardDaily, u.db.First(&flashcardDaily, "vocabulary_id = ? AND date = ?", vocabularyId, date).Error
}
