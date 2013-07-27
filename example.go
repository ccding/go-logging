package main

import (
	"github.com/ccding/go-logging/logging"
	"time"
)

func main() {
	logger := logging.SimpleLogger("main")
	logger.SetLevel(logging.NOTSET)
	logger.Error("this is a test from error")
	logger.Debug("this is a test from debug")
	logger.Log(time.Now().UnixNano())
	time.Sleep(time.Second)
}
