package main

import "fmt"

//闭包是什么？
//闭包是一个函数，这个函数包含了它外部作用域的一个变量

//底层原理：
//1.函数可以作为返回值
//2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找
func adder(x int) func(int) int { //闭包等于函数+外部变量的引用
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	ret := adder(1)
	ret2 := ret(200)
	ret3 := ret(1)
	fmt.Println(ret2)
	fmt.Println(ret3)
}
