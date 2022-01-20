package main

import "fmt"

//闭包
//闭包原理
//需要将f2传入f1进行调用
func f1(f func()) { //f1函数只能接受一个没有参数没有返回值的函数
	f()
}
func f2(x, y int) { //f1函数有两个int型的参数
	fmt.Println("This is F2")
	fmt.Println(x + y)
}

//定义一个函数对f2进行包装
func f3(f func(int, int), x, y int) func() {
	ret := func() {
		f(x, y)
	}
	return ret
}
func main() {
	ret := f3(f2, 100, 200)
	//查看ret类型
	//fmt.Printf("%T",ret)
	//实现闭包调用
	f1(ret)
}
