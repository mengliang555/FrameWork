package common

import (
	"FrameWork/util/time_util"
	"log"
	"runtime"
)

const (
	Log_Info  = "Info"
	Log_Warn  = "Warn"
	Log_Error = "Error"
)


func printLog(level string, msg string, skip int) {
	_, file, line, _ := runtime.Caller(skip)
	log.Printf("[%s]:%s [%s:%d] %s", level, time_util.GetCurrentTimeToSecond(), file, line, msg)
}

func Info(msg string) {
	printLog(Log_Info, msg, 1)
}


func Warn(msg string) {
	printLog(Log_Warn, msg, 1)
}

func Error(msg string) {
	printLog(Log_Error, msg, 1)
}
