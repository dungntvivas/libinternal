package main

import (
	"github.com/DungntVccorp/libinternal/logger"
	"time"
)

func fnlog() {
	_logger, _ := logger.New(logger.Info, []logger.Destination{logger.DestinationStdout, logger.DestinationFile}, "test.log")
	_logger.Log(logger.Info, "ABC ")
}

func main() {
	fnlog()
	time.Sleep(time.Second * 5)

}
