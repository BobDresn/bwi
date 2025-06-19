package shared

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)
	
var (
	currentLevel = LevelDebug
	logger 		 = log.New(os.Stdout, "", 0)
)

func SetLogLevel(level LogLevel) {
	currentLevel = level
}

func logWithLevel(level LogLevel, label string, format string, args ...interface{}) {
	if level < currentLevel {
		return
	}

	_, file, line, _ := runtime.Caller(2)
	shortFile := file
	if len(file) > 30 {
		shortFile = "..." + file[len(file)-27:]
	}

	timestamp := time.Now().Format("15:04:05:000")
	msg := fmt.Sprintf(format, args...)
	logger.Printf("[%s] %-5s %s:%d: %s", timestamp, label, shortFile, line, msg)
}

func Debug(msg string, args ...interface{}) {
	logWithLevel(LevelDebug, "DEBUG", msg, args...)
}

func Info(msg string, args ...interface{}) {
	logWithLevel(LevelInfo, "INFO", msg, args...)
}

func Warn(msg string, args ...interface{}) {
	logWithLevel(LevelWarn, "WARN", msg, args...)
}

func Error(msg string, args ...interface{}) {

	logWithLevel(LevelError, "ERROR", msg, args...)
}
