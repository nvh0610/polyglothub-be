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

type DashboardResponse struct {
	Fullname       string `json:"fullname" gorm:"column:fullname"`
	CorrectAnswers int    `json:"correct_answers" gorm:"column:total_correct_answers"`
	WrongAnswers   int    `json:"wrong_answers" gorm:"column:total_wrong_answers"`
}

func (u *Implement) GetByDate(startDate, endDate string, limit, offset int) ([]*DashboardResponse, int, error) {
	var dashboard []*DashboardResponse
	var total int64

	query := u.db.Table("user_daily_word_statistics AS u").
		Select("u2.fullname, SUM(u.correct_answers) AS total_correct_answers, SUM(u.wrong_answers) AS total_wrong_answers").
		Joins("JOIN `learn-language`.users u2 ON u2.id = u.user_id").
		Where("u.date BETWEEN ? AND ?", startDate, endDate).
		Group("u2.fullname")

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Limit(limit).Offset(offset).Find(&dashboard).Error
	if err != nil {
		return nil, 0, err
	}

	return dashboard, int(total), nil
}
