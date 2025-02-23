package user_daily_word_statistic

import (
	"gorm.io/gorm"
	"learn/internal/entity"
)

type Implement struct {
	db *gorm.DB
}

func NewUserDailyWordStatisticRepository(db *gorm.DB) *Implement {
	return &Implement{db: db}
}

func (u *Implement) Create(userDailyWordStatistic *entity.UserDailyWordStatistics) error {
	return u.db.Create(userDailyWordStatistic).Error
}

func (u *Implement) Update(userDailyWordStatistic *entity.UserDailyWordStatistics) error {
	return u.db.Save(userDailyWordStatistic).Error
}

func (u *Implement) GetByUserIdAndDate(userId int, date string) (*entity.UserDailyWordStatistics, error) {
	var userDailyWordStatistic *entity.UserDailyWordStatistics
	return userDailyWordStatistic, u.db.First(&userDailyWordStatistic, "user_id = ? and date = ?", userId, date).Error
}
