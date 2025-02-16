package configs

import (
	"fmt"
	"time"
)

type MySQLConfig struct {
	Username    string
	Password    string
	IP          string
	DB          string
	MaxIdleConn int
	MaxOpenConn int
	MaxLifeTime time.Duration
}

const (
	// defaultMaxIdleConns
	// the default maximum number of connections in the idle connection pool.
	defaultMaxIdleConns = 10
	// defaultMaxOpenConns
	// the default maximum number of open connections to the database.
	defaultMaxOpenConns = 100
	// defaultConnMaxLifetime
	// the default maximum amount of time a connection may be reused.
	defaultConnMaxLifetime = time.Hour
)

func (*MySQLConfig) IPEnv() string {
	return StringEnv(`MYSQL_IP`)
}

func (*MySQLConfig) UsernameEnv() string {
	return StringEnv(`MYSQL_USERNAME`)
}

func (*MySQLConfig) PasswordEnv() string {
	return StringEnv(`MYSQL_PASSWORD`)
}

func (*MySQLConfig) DatabaseEnv() string {
	return StringEnv(`MYSQL_DATABASE`)
}

func (c *MySQLConfig) LoadUriEnv() {
	uri := StringEnv(`DATABASE_MYSQL_URI`)
	v := ParseConnectionString(uri)
	c.Username = v.User.Username()
	c.Password, _ = v.User.Password()
	c.IP = v.Host
	c.DB = v.DatabaseName
	c.MaxIdleConn = IntEnvF("DB_MAX_IDLE_CONNECTIONS", 10)
	c.MaxOpenConn = IntEnvF("DB_MAX_CONNECTIONS", 100)
	c.MaxLifeTime = time.Duration(IntEnvF("DB_MAX_LIFETIME_CONNECTIONS", 1)) * time.Hour
}

func (c *MySQLConfig) LoadEnvs() {
	c.IP = c.IPEnv()
	c.Username = c.UsernameEnv()
	c.Password = c.PasswordEnv()
	c.DB = c.DatabaseEnv()
	c.MaxIdleConn = IntEnvF("DB_MAX_IDLE_CONNECTIONS", 10)
	c.MaxOpenConn = IntEnvF("DB_MAX_CONNECTIONS", 100)
	c.MaxLifeTime = time.Duration(IntEnvF("DB_MAX_LIFETIME_CONNECTIONS", 1)) * time.Hour
}

func (c *MySQLConfig) BuildConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.IP, c.DB)
}
