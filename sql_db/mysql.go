package sqldb

import (
	"fmt"
	"time"

	"github.com/dungntvivas/libinternal/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql_db struct {
	MYSQL_DB_USERNAME string
	MYSQL_DB_PASSWORD string
	MYSQL_DB_NAME     string
	MYSQL_DB_HOST     string
	logger            *logger.Logger
	cf                *gorm.Config
	db                *gorm.DB
}

func (p *Mysql_db) LogInfo(format string, args ...interface{}) {
	p.logger.Log(logger.Info, "[mysql] "+format, args...)
}
func (p *Mysql_db) LogDebug(format string, args ...interface{}) {
	p.logger.Log(logger.Debug, "[mysql] "+format, args...)
}
func (p *Mysql_db) LogError(format string, args ...interface{}) {
	p.logger.Log(logger.Error, "[mysql] "+format, args...)
}

func mysql_open_connection(p *Mysql_db) bool {
	var err error

	createDBDsn := fmt.Sprintf("%s:%s@tcp(%s)/", p.MYSQL_DB_USERNAME, p.MYSQL_DB_PASSWORD, p.MYSQL_DB_HOST)
	var database *gorm.DB
	database, err = gorm.Open(mysql.Open(createDBDsn), p.cf)
	if err == nil {
		_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + p.MYSQL_DB_NAME + " CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;")

	} else {
		return false
	}

	defer func() {
		dbInstance, _ := database.DB()
		_ = dbInstance.Close()
	}()

	dsn := p.MYSQL_DB_USERNAME + ":" + p.MYSQL_DB_PASSWORD + "@tcp" + "(" + p.MYSQL_DB_HOST + ")/" + p.MYSQL_DB_NAME + "?" + "charset=utf8mb4&parseTime=true&loc=Local"

	p.db, err = gorm.Open(mysql.Open(dsn), p.cf)
	if err != nil {
		p.LogInfo("Open Connection Error %v", err.Error())
		return false
	}
	sqlDB, err := p.db.DB()
	if err != nil {
		// control error
		return false
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(1_000)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	return true
}

func NEW_MYSQL_DB(user string, pass string, db_name string, host string, _logger *logger.Logger, _cf *gorm.Config) (*Mysql_db, bool) {
	p := &Mysql_db{
		MYSQL_DB_USERNAME: user,
		MYSQL_DB_PASSWORD: pass,
		MYSQL_DB_NAME:     db_name,
		MYSQL_DB_HOST:     host,
		logger:            _logger,
		cf:                _cf,
	}
	return p, true
}

func (p *Mysql_db) START_CONNECT_MYSQL_DB(didConnected func(db *gorm.DB)) bool {
	ok := mysql_open_connection(p)
	if ok {
		didConnected(p.db)
	}
	return ok

}
func (p *Mysql_db) GetInstance() *gorm.DB {
	return p.db
}
