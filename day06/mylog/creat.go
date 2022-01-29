package mylog

type Logger struct {
	level LogLevel
}

func NewLogger(level string) Logger {
	//解析输入日志级别返回日志结构体
	l := ParseLogLevel(level)
	return Logger{
		level: l,
	}
}

//控制日志开关
func (l Logger) enable(level LogLevel) bool {
	return level >= l.level
}

//支持格式化输出
func (l Logger) Debug(format string, a ...interface{}) {
	if l.enable(DUBUG) {
		log(DUBUG, format, a...)
	}
}
func (l Logger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format, a...)
	}
}
func (l Logger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, a...)
	}
}
func (l Logger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format, a...)
	}
}
func (l Logger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL, format, a...)
	}
}
