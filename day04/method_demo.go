package main

import "fmt"

//给自定类型添加方法
//不能给别的包里边的类型定义方法，只能给自己的包里定义方法
type myInt int

func (m myInt) hello() {
	fmt.Println("我是Myint")
}
func main() {
	i := myInt(100) //类型声明
	i.hello()
}
