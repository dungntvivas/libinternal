package sqldb

import (
	"gitlab.vivas.vn/go/libinternal/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql_db struct {
	MYSQL_DB_USERNAME string
	MYSQL_DB_PASSWORD string
	MYSQL_DB_NAME     string
	MYSQL_DB_HOST     string
	logger            *logger.Logger

	db *gorm.DB
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

func mysql_connectionDB(p *Mysql_db) bool {
	var err error
	dsn := p.MYSQL_DB_USERNAME + ":" + p.MYSQL_DB_PASSWORD + "@tcp" + "(" + p.MYSQL_DB_HOST + ")/" + p.MYSQL_DB_NAME + "?" + "parseTime=true&loc=Local"

	p.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: glog.Default.LogMode(glog.Silent),
	})
	if err != nil {
		p.LogInfo("Open Connection Error %v", err.Error())
		return false
	}
	return true
}

func NEW_MYSQL_DB(user string, pass string, db_name string, host string, _logger *logger.Logger) (*Mysql_db, bool) {
	p := &Mysql_db{
		MYSQL_DB_USERNAME: user,
		MYSQL_DB_PASSWORD: pass,
		MYSQL_DB_NAME:     db_name,
		MYSQL_DB_HOST:     host,
		logger:            _logger,
	}
	return p, true
}

func (p *Mysql_db) START_CONNECT_MYSQL_DB(didConnected func(db *gorm.DB)) bool {
	ok := mysql_connectionDB(p)
	if ok {
		didConnected(p.db)
	}
	return ok

}
func (p *Mysql_db) GetInstance() *gorm.DB {
	return p.db
}