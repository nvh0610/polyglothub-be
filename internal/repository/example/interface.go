package example

import "learn/internal/entity"

type Repository interface {
	GetById(id int) (*entity.Example, error)
	List(limit, offset int, vocabularyId int) ([]*entity.Example, int, error)
	Create(example *entity.Example) error
	Update(example *entity.Example) error
	Delete(id int) error
	DeleteByVocabularyId(vocabularyId int) error
	CreateBatch(examples []*entity.Example) error
	UpsertBatch(examples []*entity.Example) error
}
