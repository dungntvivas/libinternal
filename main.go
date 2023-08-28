package main

import (
	"github.com/DungntVccorp/libinternal/logger"
	redisDB "github.com/DungntVccorp/libinternal/redis_db"
	"github.com/redis/go-redis/v9"
	"time"
)

func fnlog() {
	_logger, _ := logger.New2(logger.Info, []logger.Destination{logger.DestinationFile, logger.DestinationUdplog,logger.DestinationStdout}, "test.log", "127.0.0.1:44953", "ServerA")

	//_db := redisDB.New(_logger,&redis.ClusterOptions{
	//	Addrs: []string{"10.3.3.33:6379"},
	//	NewClient: func(opt *redis.Options) *redis.Client {
	//		opt.Username = "default"
	//		opt.Password = "123456a@"
	//		return redis.NewClient(opt)
	//	},
	//
	//})
	_db := redisDB.NewRedisClient(_logger,&redis.UniversalOptions{
		Addrs: []string{"10.3.3.33:6379"},
		Username: "default",
		Password: "123456a@",
		DB: 10,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second*10,
		DialTimeout:time.Second*5,
		PoolSize: 10,

	})
	_logger.Log(logger.Info,"%v",_db.Run())




	//for  {
	//	var i int = 0
	//	for i < 10 {
	//		_logger.Log(logger.Info, "%v", i)
	//		i++
	//	}
	//	_logger.Log(logger.Info, "ABC ")
	//	time.Sleep(time.Second * 20)
	//}





}

func main() {



	fnlog()
	time.Sleep(time.Second * 5)

}
