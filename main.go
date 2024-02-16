package main

import (
	"time"

	"github.com/dungntvivas/libinternal/logger"
)

func fnlog() {
	//_logger, _ := logger.New2(logger.Info, []logger.Destination{logger.DestinationFile, logger.DestinationUdplog, logger.DestinationStdout}, "test.log", "10.3.3.115:44953", "ServerA")
	_logger, _ := logger.New(logger.Info, []logger.Destination{logger.DestinationFile, logger.DestinationUdplog, logger.DestinationStdout}, "test.log")

	//_db := redisDB.New(_logger,&redis.ClusterOptions{
	//	Addrs: []string{"10.3.3.33:6379"},
	//	NewClient: func(opt *redis.Options) *redis.Client {
	//		opt.Username = "default"
	//		opt.Password = "123456a@"
	//		return redis.NewClient(opt)
	//	},
	//
	//})
	// _db := redisDB.NewRedisClient(_logger, &redis.UniversalOptions{
	// 	Addrs:        []string{"10.3.3.33:6379"},
	// 	Username:     "default",
	// 	Password:     "123456a@",
	// 	DB:           10,
	// 	ReadTimeout:  time.Second * 10,
	// 	WriteTimeout: time.Second * 10,
	// 	DialTimeout:  time.Second * 5,
	// 	PoolSize:     10,
	// })
	// _logger.Log(logger.Info, "%v", _db.Run())
	var _ok bool = false
	var c int = 0
	for {
		var i int = 0
		for i < 10 {
			_logger.Log(logger.Info, "%v", i)
			i++
			c += 1
		}
		_logger.Log(logger.Info, "ABC %v", c)
		time.Sleep(time.Second * 5)
		if !_ok {
			_ok = _logger.EnableUDPLogServer("ipcam.vivas.vn:44953", "ServerA")
		}

		if _ok && c == 30 {
			_logger.DisableUDPLogServer()
		}

	}

}

func main() {

	fnlog()
	time.Sleep(time.Second * 5)

}
