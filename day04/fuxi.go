package main

import "fmt"

// func f(name string) {
// 	fmt.Println("Hello", name)
// }
// //函数作为参数
// func f1(f func(string), name string) {
// 	f(name)
// }

// //函数作为返回值
// func f2() func(int, int) int {
// 	return func(x, y int) int {
// 		return x + y
// 	}
// }

func funcdemo(f func()) {
	f()
}
func f1(name string) {
	fmt.Println("Hello", name)
}
func f2(f func(string), name string) func() {
	return func() {
		f(name)
	}
}
func main() {
	// f1(f, "123") //传入两个参数 第一个为func，第二个为string
	// ret := f2()
	// fmt.Printf("%T\n", ret)
	// sum := ret(1, 2)
	// fmt.Println(sum)

	//闭包
	funcdemo(f2(f1, "123"))
}
