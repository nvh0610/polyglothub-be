package user_flashcard_log

import "learn/internal/entity"

type Repository interface {
	Create(userFlashcardLog *entity.UserFlashcardLog) error
	GetByUserIdAndDateAndVocabularyIdAndIsCorrect(userId int, date string, vocabularyId int, isCorrect bool) (*entity.UserFlashcardLog, error)
}
