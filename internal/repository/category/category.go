package category

import (
	"gorm.io/gorm"
	"learn/internal/entity"
)

type Implement struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}
func (u *Implement) GetById(id int) (*entity.Category, error) {
	var category *entity.Category
	return category, u.db.First(&category, "id = ?", id).Error
}

func (u *Implement) Create(category *entity.Category) error {
	return u.db.Create(category).Error
}

func (u *Implement) Update(category *entity.Category) error {
	return u.db.Save(category).Error
}

func (u *Implement) Delete(id int) error {
	return u.db.Delete(&entity.Category{Id: id}).Error
}

func (u *Implement) List(limit, offset int, userId int) ([]*entity.Category, int, error) {
	var categories []*entity.Category
	var count int64

	query := u.db.Limit(limit).Offset(offset)
	query = query.Where("user_id = ? OR user_id = 0", userId)
	err := query.Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	if err = u.db.Model(&entity.Category{}).Where("user_id = ? OR user_id = 0", userId).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return categories, int(count), err
}

func (u *Implement) GetByIdAndUserId(id int, userId int) (*entity.Category, error) {
	var category *entity.Category
	if userId == 0 {
		return category, u.db.First(&category, "id = ?", id).Error
	}
	return category, u.db.First(&category, "id = ? AND user_id = ?", id, userId).Error
}
