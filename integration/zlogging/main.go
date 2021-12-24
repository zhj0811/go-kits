package main

import "github.com/zhj0811/go-kits/zlogging"

func main() {
	logger := zlogging.MustGetLogger("test")
	logger.Info("Test info log")
	zlogging.SetLogLevel("warning")
	logger.Info("Test info log level")
	zlogging.SetLogLevel("debug")
	logger.Info("Test info log level")
}