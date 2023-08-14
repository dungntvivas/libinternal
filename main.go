package main

import (
	"github.com/DungntVccorp/libinternal/logger"
	"time"
)

func fnlog() {
	_logger, _ := logger.New2(logger.Info, []logger.Destination{logger.DestinationFile, logger.DestinationUdplog}, "test.log", "10.84.86.34:44953", "server_abc")
	var i int = 0
	for i < 1000 {
		_logger.Log(logger.Info, "%v", i)
		i++
	}
	_logger.Log(logger.Info, "ABC ")
}

func main() {
	fnlog()
	time.Sleep(time.Second * 5)

}
