package main

import "fmt"

func main() {
	//1.&取地址

	// n := 18
	// p := &n
	// fmt.Println(p)
	// fmt.Printf("type(p):%T\n", p) //int类型的指针
	// //2.*根据地址取值
	// m:=*p
	// fmt.Println(m)
	// fmt.Printf("%T\n",m)

	//new函数申请一个内存地址
	var a *int
	fmt.Println(a) //未初始化得到零值 nil
	var a2 = new(int)
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
}
