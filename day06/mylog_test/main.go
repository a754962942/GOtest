package main

import (
	"github.com/a754962942/Gotest/day06/mylog"
)

func main() {
	for {
		logger:=mylog.NewLogger("Debug")
		//logger= mylog.NewFileLogger("Info", "123", "./", 1000)
		Log.Debug("这是一条Debug")
		logger.Info("这是一条Info")
		logger.Error("这是一条Error")
		id := 100
		name := "CHAIN"
		logger.Warning("这是一条Warning,FROM:%d,name:%s", id, name)
		logger.Fatal("这是一条Fatal")
		// time.Sleep(2 * time.Second)
	}
}
