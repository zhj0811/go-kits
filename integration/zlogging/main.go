package main

import (
	"github.com/zhj0811/go-kits/zlogging"
)

func main() {
	//z := &lumberjack.Logger{
	//	Filename:   "./foo.log",
	//	MaxSize:    500, // megabytes
	//	MaxBackups: 3,
	//	MaxAge:     28, //days
	//	Compress:   true, // disabled by default
	//}
	//conf := zlogging.Config{
	//	Writer: z,
	//}
	//zlogging.Init(conf)
	logger := zlogging.MustGetLogger("test")
	logger.Info("Test info log")
	zlogging.SetLogLevel("warning")
	logger.Info("Test info log level should not exist")
	zlogging.SetLogLevel("debug")
	logger.Info("Test info log level")
	logger.Error("Test failed")
	logger.Info("Test info log level")

	b := zlogging.MustGetLogger("btest")
	b.Info("Test info log level")
	b.Error("Test failed")
	zlogging.SetLogLevel("warning")
	logger.Info("Test info log level should not exist")
	b.Info("Test info log level")
	//logger.Infow()
}