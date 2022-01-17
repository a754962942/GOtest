package main

import "fmt"

//函数
// 函数存在的意义
// 函数是一段代码的引擎
// 把一段逻辑抽象出来封装到一个函数中，给它起个名字，每次用到它的时候使用函数名调用即可
// 使用函数能够使代码结构更清晰、更简洁。

//函数的定义
func sum(x, y int) int {
	return x + y
}

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数、没有返回值
func f2() {
	fmt.Println("分")
}

//没有参数但有返回值
func f3() int {
	return 3
}

//参数可以命名也可以不命名
// 命名的返回值就相当于在函数中声明一个局部变量
func f4(x, y int) (ret int) {
	ret = x + y
	return
}

//多返回值
func f5() (int, string) {

}

//参数的类型简写：
//当参数中连续多个参数的类型一致时，我们可以将非最后一个参数的类型省略
func f6(x, y, z int, m, n string, i, j bool) {

}

//可变长参数
//可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y的类型是切片 []int
}
func main() {
	f1(2, 2)
	f2()
	s := f3()
	fmt.Println(s)
	s1 := f4(4, 6)
	fmt.Println(s1)
}
