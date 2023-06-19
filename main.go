package main

import (
	"fmt"
	"gitlab.vivas.vn/go/libinternal/logger"
	"time"
)


func fnlog(){
	_logger, _ := logger.New(logger.Info, logger.LogDestinations{logger.DestinationFile: {}, logger.DestinationStdout: {}}, "test.log")
	_logger.Log(logger.Info,"ABC ")
}

func fnLog2(){
	_logger, err := logger.New2(logger.Info, logger.LogDestinations{logger.DestinationFile: {}, logger.DestinationStdout: {},logger.DestinationUdplog: {}}, "test2.log","10.3.3.119",44953,"SVTest")
	if err != nil {
		fmt.Printf(err.Error())
	}else{
		_logger.Log(logger.Info,"Test Log UDP %v",time.Now().String())
		_logger.Log(logger.Info,"Log Dong 2")
		_logger.Log(logger.Info,"Log Dong 3")
		_logger.Log(logger.Info,"Log Dong 4")
		_logger.Log(logger.Info,"Log Dong 5")
		_logger.Log(logger.Info,"Log Dong 6")
		_logger.Log(logger.Info,"Log Dong 7")
		_logger.Log(logger.Info,"Log Dong 8")
		_logger.Log(logger.Info,"Log Dong 9")
		_logger.Log(logger.Info,"Log Dong 10")
		_logger.Log(logger.Info,"Log Dong 11")
		_logger.Log(logger.Info,"Log Dong 12")
		_logger.Log(logger.Info,"Log Dong 13")
		_logger.Log(logger.Info,"Log Dong 14")
		_logger.Log(logger.Info,"Log Dong 15")
		_logger.Log(logger.Info,"Log Dong 16")
		_logger.Log(logger.Info,"Log Dong 17")
		_logger.Log(logger.Info,"Log Dong 18")
		_logger.Log(logger.Info,"Log Dong 19")
		_logger.Log(logger.Info,"Log Dong 20")
		time.Sleep(time.Second*2)
		_logger.Log(logger.Info,"Test Log UDP %v",time.Now().String())
		_logger.Log(logger.Info,"Log Dong 2")
		_logger.Log(logger.Info,"Log Dong 3")
		_logger.Log(logger.Info,"Log Dong 4")
		_logger.Log(logger.Info,"Log Dong 5")
		_logger.Log(logger.Info,"Log Dong 6")
		_logger.Log(logger.Info,"Log Dong 7")
		_logger.Log(logger.Info,"Log Dong 8")
		_logger.Log(logger.Info,"Log Dong 9")
		_logger.Log(logger.Info,"Log Dong 10")
		_logger.Log(logger.Info,"Log Dong 11")
		_logger.Log(logger.Info,"Log Dong 12")
		_logger.Log(logger.Info,"Log Dong 13")
		_logger.Log(logger.Info,"Log Dong 14")
		_logger.Log(logger.Info,"Log Dong 15")
		_logger.Log(logger.Info,"Log Dong 16")
		_logger.Log(logger.Info,"Log Dong 17")
		_logger.Log(logger.Info,"Log Dong 18")
		_logger.Log(logger.Info,"Log Dong 19")
		_logger.Log(logger.Info,"Log Dong 20")


	}

}
func fnLog3(){
	_logger, err := logger.New2(logger.Info, logger.LogDestinations{logger.DestinationFile: {}, logger.DestinationStdout: {},logger.DestinationUdplog: {}}, "test3.log","10.3.3.119",44953,"SVTest2")
	if err != nil {
		fmt.Printf(err.Error())
	}else{
		_logger.Log(logger.Info,"Test Log UDP %v",time.Now().String())
		_logger.Log(logger.Info,"Log Dong 2")
		_logger.Log(logger.Info,"Log Dong 3")
		_logger.Log(logger.Info,"Log Dong 4")
		_logger.Log(logger.Info,"Log Dong 5")
		_logger.Log(logger.Info,"Log Dong 6")
		_logger.Log(logger.Info,"Log Dong 7")
		_logger.Log(logger.Info,"Log Dong 8")
		_logger.Log(logger.Info,"Log Dong 9")
		_logger.Log(logger.Info,"Log Dong 10")
		_logger.Log(logger.Info,"Log Dong 11")
		_logger.Log(logger.Info,"Log Dong 12")
		_logger.Log(logger.Info,"Log Dong 13")
		_logger.Log(logger.Info,"Log Dong 14")
		_logger.Log(logger.Info,"Log Dong 15")
		_logger.Log(logger.Info,"Log Dong 16")
		_logger.Log(logger.Info,"Log Dong 17")
		_logger.Log(logger.Info,"Log Dong 18")
		_logger.Log(logger.Info,"Log Dong 19")
		_logger.Log(logger.Info,"Log Dong 20")
		time.Sleep(time.Second*2)
		_logger.Log(logger.Info,"Test Log UDP %v",time.Now().String())
		_logger.Log(logger.Info,"Log Dong 2")
		_logger.Log(logger.Info,"Log Dong 3")
		_logger.Log(logger.Info,"Log Dong 4")
		_logger.Log(logger.Info,"Log Dong 5")
		_logger.Log(logger.Info,"Log Dong 6")
		_logger.Log(logger.Info,"Log Dong 7")
		_logger.Log(logger.Info,"Log Dong 8")
		_logger.Log(logger.Info,"Log Dong 9")
		_logger.Log(logger.Info,"Log Dong 10")
		_logger.Log(logger.Info,"Log Dong 11")
		_logger.Log(logger.Info,"Log Dong 12")
		_logger.Log(logger.Info,"Log Dong 13")
		_logger.Log(logger.Info,"Log Dong 14")
		_logger.Log(logger.Info,"Log Dong 15")
		_logger.Log(logger.Info,"Log Dong 16")
		_logger.Log(logger.Info,"Log Dong 17")
		_logger.Log(logger.Info,"Log Dong 18")
		_logger.Log(logger.Info,"Log Dong 19")
		_logger.Log(logger.Info,"Log Dong 20")


	}

}

func main()  {
	//fnlog()
	go fnLog2()
	go fnLog3()
	time.Sleep(time.Second*5)

}
