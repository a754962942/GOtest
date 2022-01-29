package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	level       LogLevel
	fileName    string
	filePath    string
	fileObj     *os.File
	errfileObj  *os.File
	maxFileSize int64
}


//构造函数
func NewFileLogger(levelstr, fn, fp string, maxFileSize int64) *FileLogger {
	level := ParseLogLevel(levelstr)
	f1 := &FileLogger{
		level:       level,
		fileName:    fn,
		filePath:    fp,
		maxFileSize: maxFileSize,
	}
	err := f1.initFile() //按照文件路径和文件名打开
	if err != nil {
		panic("发生错误")
	}
	return f1
}

//根据文件大小分割
//接受参数为fileobj、errfileobj
func (f *FileLogger) checkSize(file *os.File) bool {
	FileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file Info failed,err:%v \n", err)
		return false
	}
	// fmt.Printf("文件大小是：%d B\n", FileInfo.Size())
	//如果文件大小大于等于日志文件的最大值就返回true
	return FileInfo.Size() > f.maxFileSize
}

func (f *FileLogger) initFile() (err error) {
	loc := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(loc, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("打开文件失败，err:\n", err)
		return err
	}
	errfileObj, err := os.OpenFile(loc+".err", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("打开ERR文件失败，err:\n", err)
		return err
	}
	//日志文件打开成功，开始写入
	f.fileObj = fileObj
	f.errfileObj = errfileObj
	return nil
}
func (f FileLogger) enable(level LogLevel) bool {
	return level >= f.level
}
func (f *FileLogger) split(file *os.File) (*os.File, error) {
	//需要切割日志文件
	nowstr := time.Now().Format("2016010215040500")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file Info failed,err:%v \n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowstr)
	//1.关闭当前的日志文件
	file.Close()
	//2.备份一下 rename
	os.Rename(logName, newLogName)
	//3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return nil, err
	}
	return fileObj, nil

}
func (f *FileLogger) log(l LogLevel, format string, a ...interface{}) {
	if f.enable(l) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		fileName, filepath, line := getInfo(2)
		level := logParse(l)
		if f.checkSize(f.fileObj) == true {
			newfile, err := f.split(f.fileObj)
			if err != nil {
				fmt.Printf("新建文件出错，err:%v\n", err)
				return
			}
			f.fileObj = newfile
		}
		fmt.Fprintf(f.fileObj, "[%s]\t[%s][%s]%s:%d\t%s \n", now.Format("2006-01-02 15:04:05"), level, fileName, filepath, line, msg)
		if l >= ERROR {
			if f.checkSize(f.errfileObj) == true {
				newfile, err := f.split(f.errfileObj)
				if err != nil {
					fmt.Printf("新建文件出错，err:%v\n", err)
					return
				}
				f.errfileObj = newfile
			}
			fmt.Fprintf(f.errfileObj, "[%s]\t[%s][%s]%s:%d\t%s \n", now.Format("2006-01-02 15:04:05"), level, fileName, filepath, line, msg)
		}
	}
}

//支持格式化输出
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DUBUG, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
func (f *FileLogger) toTheEnd() {
	f.fileObj.Close()
	f.errfileObj.Close()
}
