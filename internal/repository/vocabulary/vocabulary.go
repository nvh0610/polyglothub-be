package vocabulary

import (
	"gorm.io/gorm"
	"learn/internal/entity"
)

type Implement struct {
	db *gorm.DB
}

func NewVocabularyRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) List(limit, offset int, categoryId int) ([]*entity.Vocabulary, int, error) {
	var vocabularies []*entity.Vocabulary
	var count int64

	query := u.db.Limit(limit).Offset(offset)
	if categoryId != 0 {
		query = query.Where("category_id = ? OR user_id = 0", categoryId)
	}

	err := query.Find(&vocabularies).Error
	if err != nil {
		return nil, 0, err
	}

	if err = u.db.Model(&entity.Vocabulary{}).Where("category_id = ? OR category_id = 0", categoryId).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return vocabularies, int(count), err
}

func (u *Implement) Create(vocabulary *entity.Vocabulary) error {
	return u.db.Create(vocabulary).Error
}

func (u *Implement) Update(vocabulary *entity.Vocabulary) error {
	return u.db.Save(vocabulary).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.Vocabulary{Id: id}).Error
}

func (u *Implement) GetById(id int) (*entity.Vocabulary, error) {
	var vocabulary *entity.Vocabulary
	return vocabulary, u.db.First(&vocabulary, "id = ?", id).Error
}
