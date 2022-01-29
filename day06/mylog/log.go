package mylog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel int16

const (
	UNKNOW LogLevel = iota
	DUBUG
	INFO
	ERROR
	WARNING
	FATAL
)

func ParseLogLevel(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DUBUG
	case "info":
		return INFO
	case "error":
		return ERROR
	case "warning":
		return WARNING
	case "fatal":
		return FATAL
	default:
		panic("输入有误")
	}
}

func logParse(l LogLevel) string {
	switch l {
	case DUBUG:
		return "Debug"
	case INFO:
		return "Info"
	case ERROR:
		return "Error"
	case WARNING:
		return "Warning"
	case FATAL:
		return "Fatal"
	}
	return "dubug"
}
type Log interface {
	Debug()
	Info()
	Error()
	Warning()
	Fatal()
}
func getInfo(skip int) (fileName, filepath string, line int) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		fmt.Println("runtime.Caller failed")
		return
	}
	fileName = runtime.FuncForPC(pc).Name()
	fileName = strings.Split(fileName, ".")[1]
	filepath = path.Base(file)
	return
}
func log(l LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	fileName, filepath, line := getInfo(2)
	level := logParse(l)
	fmt.Printf("[%s]\t[%s][%s]%s:%d\t%s \n", now.Format("2006-01-02 15:04:05"), level, fileName, filepath, line, msg)
}
