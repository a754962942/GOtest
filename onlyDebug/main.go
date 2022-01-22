package main

import "fmt"

// func PointerParam(ip *int) {
// 	fmt.Printf("指针为函数参数的内存地址是：%p\n", &ip)
// 	*ip = 2
// }

//斐波那契数列（Fibonacci sequence），又称黄金分割数列，指的是这样一个数列：0、1、1、2、3、5、8、13、21、34、……。
//使用递归方式实现输出前20个斐波那契数
//解题思路  斐波那契数列：f(0)=0 f(1)=1 f(n) = f(n-1)+f(n-2)
//递归思路
// func fib(n int) int {
// 	if n <= 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 1
// 	}
// 	return fib(n-1) + fib(n-2)
// }

//非递归思路
//使用闭包返回一个返回值为int的func
func fib2() func() int {
	a, b := 0, 1
	return func() int {
		temp := a
		a, b = b, a+b
		return temp
	}
}
func main() {
	//利用闭包思路解决斐波那契数列
	f := fib2()
	s := make([]int, 0, 10)
	for i := 0; i < 20; i++ {
		s = append(s, f())
	}
	fmt.Println(s)
	//利用递归思路计算斐波那契数列
	// 	s := make([]int, 0)
	// 	for i := 0; i < 20; i++ {
	// 		s = append(s, fib(i))
	// 	}
	// 	fmt.Println("s:", s)

	// p := 1
	// ip := &p
	// fmt.Printf("原始p的内存地址是: %p\n", &p)
	// fmt.Printf("原始指针的内存地址是: %p\n", &ip)
	// PointerParam(ip)
	// fmt.Printf("int值被修改了，新值为: %d\n", p)
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("输入的是：",s)

	// var (
	// 	name string
	// 	age  int
	// 	home string
	// )
	// fmt.Scanf("%s %d %s\n", &name, &age, &home)
	// fmt.Println(name, age, home)
	// s:="s"
	// s1:="s"
	// fmt.Println(s ==s1)
}
