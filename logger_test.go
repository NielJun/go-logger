package logger

import (
	"testing"
)

//TestFileLogger：12
//Debug：79
//writeLog：68
//GetLineInfo:10
//runtime.Caller(3)
func TestFileLogger(t *testing.T) {
	config := make(map[string]string)
	config["log_level"] = "debug"
	config["log_path"] = "c:/logs/"
	config["log_name"] = "test"
	logger,_ := NewFileLogger(config)
	logger.Debug("user id[%d] is come from china", 324234)
	logger.Warn("test warn log")
	logger.Fatal("test fatal log")
	logger.Close()
}

//TestFileLogger：12
//Debug：79
//writeLog：68
//GetLineInfo:10
//runtime.Caller(3)
func TestConsoleLogger(t *testing.T) {
	config := make(map[string]string)
	config["log_level"] = "debug"
	logger,_ := NewConsoleLogger(config)
	logger.Debug("user id[%d] is come from china", 324234)
	logger.Warn("test warn log")
	logger.Fatal("test fatal log")
	logger.Close()
}
