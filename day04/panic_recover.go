package main

import "fmt"

//panic发生时，执行defer 不输出panic
func f1() {
	defer func() {
		err := recover() //收集当前panic
		//此处可以对捕获的panic进行判断
		fmt.Println("捕获panic")
		fmt.Println(err)
	}()
	panic("错误")
	fmt.Println("this is f1")
}
func f2() {
	fmt.Println("this is f2")
}
func main() {
	f1()
	f2()
}
