package main

import "fmt"

// func PointerParam(ip *int) {
// 	fmt.Printf("指针为函数参数的内存地址是：%p\n", &ip)
// 	*ip = 2
// }

func main() {
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
	s:="s"
	s1:="s"
	fmt.Println(s ==s1)
}
