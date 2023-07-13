package main

import "github.com/scagogogo/sca-base-module-logger/logger"

func main() {
	logger.Infof("my name is %s", "CC11001100")
	logger.Errorf("hello %s", "CC11001100")
	logger.Debug("this is debug")
}
