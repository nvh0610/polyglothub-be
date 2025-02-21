package vocabulary

import (
	"gorm.io/gorm"
	"learn/internal/entity"
	"time"
)

type Implement struct {
	db *gorm.DB
}

func NewVocabularyRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

type Vocabularies struct {
	Id         int               `json:"id" gorm:"id"`
	Word       string            `json:"word" gorm:"word"`
	Meaning    string            `json:"meaning" gorm:"meaning"`
	IPA        string            `json:"ipa" gorm:"ipa"`
	Type       string            `json:"type" gorm:"type"`
	Url        string            `json:"url" gorm:"url"`
	CategoryID int               `json:"category_id" gorm:"category_id"`
	CreatedAt  time.Time         `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at" gorm:"updated_at"`
	Examples   []*entity.Example `json:"examples" gorm:"examples"`
}

func (u *Implement) List(limit, offset int, categoryId int, word string) ([]*Vocabularies, int, error) {
	var vocabularies []*Vocabularies
	var count int64

	query := u.db.Where("category_id = ?", categoryId)

	if word != "" {
		query = query.Where("word LIKE ?", "%"+word+"%")
	}

	err := query.Model(&entity.Vocabulary{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	query = u.db.Limit(limit).Offset(offset)
	query.Joins("JOIN examples ON examples.vocabulary_id = vocabularies.id")
	err = query.Find(&vocabularies).Error
	if err != nil {
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
