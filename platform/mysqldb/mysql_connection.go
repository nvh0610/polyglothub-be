package mysqldb

import (
	"cb-manager/pkg/configs"
	"gorm.io/gorm"
)

func NewMysqlAsteriskConnection() (*gorm.DB, error) {
	opts := configs.NewMySQLAsteriskConfig()
	return newMysqlConnection(opts)
}

func NewMysqlCallBotConnection() (*gorm.DB, error) {
	opts := configs.NewMySQLCallBotConfig()
	return newMysqlConnection(opts)
}
