package main

import (
	"fmt"
	"runtime"
	"strings"
)

func f2() {
	//caller(skip)根据层数算
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("runtime.Caller failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(line)
}
func f1() {
	f2()
}

//runtime.Caller
func main() {
	f1()

}
