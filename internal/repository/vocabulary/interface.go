package vocabulary

import "learn/internal/entity"

type Repository interface {
	GetById(id int) (*entity.Vocabulary, error)
	Create(user *entity.Vocabulary) error
	Update(user *entity.Vocabulary) error
	Delete(id int) error
	List(limit, offset int, categoryId int, word string) ([]*Vocabularies, int, error)
	CheckExistsByWord(word string, categoryId int) (bool, error)
	GetVocabulariesByIds(ids []int) ([]*entity.Vocabulary, error)
}
