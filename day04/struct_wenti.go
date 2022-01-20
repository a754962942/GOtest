package main

import "fmt"

//结构体遇到的问题

//1.myint（100）
type myInt int

func (m myInt) hello() {
	fmt.Println("MYINT")
}
func main() {
	//var m myInt = 10
	m:=myInt(100) //类型转换
	m.hello()
}
