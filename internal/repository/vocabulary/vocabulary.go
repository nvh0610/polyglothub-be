package vocabulary

import (
	"errors"
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
	Id          int        `json:"id" gorm:"column:id"`
	Word        string     `json:"word" gorm:"column:word"`
	Meaning     string     `json:"meaning" gorm:"column:meaning"`
	IPA         string     `json:"ipa" gorm:"column:ipa"`
	Type        string     `json:"type" gorm:"column:type"`
	Url         string     `json:"url" gorm:"column:url"`
	Description string     `json:"description" gorm:"column:description"`
	CategoryID  int        `json:"category_id" gorm:"column:category_id"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at"`
	Examples    []*Example `json:"examples" gorm:"foreignKey:VocabularyID"`
}

type Example struct {
	Id           int       `json:"id" gorm:"column:id"`
	Sentence     string    `json:"sentence" gorm:"column:sentence"`
	Meaning      string    `json:"meaning" gorm:"column:meaning"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
	VocabularyID int       `json:"vocabulary_id" gorm:"column:vocabulary_id"`
}

func (u *Implement) List(limit, offset int, categoryId int, word string) ([]*Vocabularies, int, error) {
	var vocabularies []*Vocabularies
	var count int64

	query := u.db.Model(&Vocabularies{}).Where("category_id = ?", categoryId)

	if word != "" {
		query = query.Where("word LIKE ?", "%"+word+"%")
	}

	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.
		Limit(limit).
		Offset(offset).
		Preload("Examples").Find(&vocabularies).Error

	if err != nil {
		return nil, 0, err
	}

	return vocabularies, int(count), nil
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

func (u *Implement) CheckExistsByWord(word string, categoryId int) (bool, error) {
	var vocabulary entity.Vocabulary
	err := u.db.First(&vocabulary, "word = ? AND category_id = ?", word, categoryId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}

		return false, err
	}

	return false, nil
}

func (u *Implement) GetVocabulariesByIds(ids []int) ([]*entity.Vocabulary, error) {
	var vocabularies []*entity.Vocabulary
	return vocabularies, u.db.Where("id in (?)", ids).Find(&vocabularies).Error
}
