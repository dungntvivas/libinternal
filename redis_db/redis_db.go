package redisDB

import (
	"context"
	"github.com/DungntVccorp/libinternal/logger"
	"github.com/redis/go-redis/v9"
)



type RedisDB struct {
	logger             *logger.Logger
	rdb				   redis.UniversalClient
}
func (p *RedisDB) LogInfo(format string, args ...interface{}) {
	p.logger.Log(logger.Info, "[redis] "+format, args...)
}
func (p *RedisDB) LogDebug(format string, args ...interface{}) {
	p.logger.Log(logger.Debug, "[redis] "+format, args...)
}
func (p *RedisDB) LogError(format string, args ...interface{}) {
	p.logger.Log(logger.Error, "[redis] "+format, args...)
}

func NewRedisClient(_logger *logger.Logger,option *redis.UniversalOptions) *RedisDB{
	p := RedisDB{
		logger: _logger,
		rdb: redis.NewUniversalClient(option),
	}
	return &p
}

func (p *RedisDB)Run() bool{
	status := p.rdb.Ping(context.Background())
	if status.Err() != nil {
		return false
	}
	return  true
}
func (p *RedisDB)Close(){
	p.rdb.Close()
}
