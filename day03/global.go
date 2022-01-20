package main

import "fmt"

//变量的作用域
//1.全局作用域
var x = 100 //定义一个全局变量
//定义一个函数

func f1() {
	x := 10
	name := "123"
	//函数中查找变量的顺序
	//1.先在函数内部查找。
	//2.找不到就往函数外部查找，一直找到全局
	//如果全局没有找到，报错

	fmt.Println(x, name)

}
func main() {
	f1()
	//2.函数作用域
	//函数内部定义的变量只能在该函数内部使用
	//fmt.Println(name)

	//3.语句块作用域
	if i := 10; i < 18 {
		fmt.Println("乖乖上学")
	}
	//fmt.Println(i)
	for j := 0; j < count; j++ {
		fmt.Println("123")
	}
	//fmt.Println(j)
}
