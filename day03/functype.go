package main

import "fmt"

//函数类型
func f1() {
	fmt.Println("Hello")
}

func f2() int {
	return 10
}

//函数可以作为参数的类型
func f3(x func() int) int {
	ret := x()
	return ret
}
func f4(x, y int) int {
	return x + y
}
func ff(x, y int) int {
	return x + y
}
func f5(x func() int) func(int, int) int {

	return ff
}
func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	fmt.Println(f3(f2))
	fmt.Println(f3(b))

	//f3(f4)
}
