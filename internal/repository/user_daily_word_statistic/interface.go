package user_daily_word_statistic

import "learn/internal/entity"

type Repository interface {
	Create(userDailyWordStatistic *entity.UserDailyWordStatistics) error
	Update(userDailyWordStatistic *entity.UserDailyWordStatistics) error
	GetByUserIdAndDate(userId int, date string) (*entity.UserDailyWordStatistics, error)
}
