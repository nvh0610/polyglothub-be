package example

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"learn/internal/entity"
)

type Implement struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) GetById(id int) (*entity.Example, error) {
	var example *entity.Example
	return example, u.db.First(&example, "id = ?", id).Error
}

func (u *Implement) List(limit, offset int, vocabularyId int) ([]*entity.Example, int, error) {
	var examples []*entity.Example
	var count int64
	err := u.db.Limit(limit).Offset(offset).Where("vocabulary_id = ?", vocabularyId).Find(&examples).Error
	if err != nil {
		return nil, 0, err
	}

	if err = u.db.Model(&entity.Example{}).Where("vocabulary_id = ?", vocabularyId).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return examples, int(count), err
}

func (u *Implement) Create(example *entity.Example) error {
	return u.db.Create(example).Error
}

func (u *Implement) Update(example *entity.Example) error {
	return u.db.Save(example).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.Example{Id: id}).Error
}

func (u *Implement) CreateBatch(examples []*entity.Example) error {
	return u.db.Create(examples).Error
}

func (u *Implement) UpsertBatch(examples []*entity.Example) error {
	return u.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"sentence", "meaning"}),
	}).Create(examples).Error
}

func (u *Implement) DeleteByVocabularyId(vocabularyId int) error {
	return u.db.Where("vocabulary_id = ?", vocabularyId).Delete(&entity.Example{}).Error
}
