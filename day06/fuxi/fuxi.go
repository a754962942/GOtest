package main

import "fmt"

//类型断言
func main() {
	var a interface{} //定义一个空接口变量a。

	//如何判断a保存的值的具体类型是什么？
	//类型断言
	//1.x.(T)
	a = 100
	// v1, ok := a.(int8)
	// if ok {
	// 	fmt.Println("v1是int8")
	// }
	//2.switch
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int32:
		fmt.Println("int32", v2)
	case string:
		fmt.Println("string", v2)
	case int:
		fmt.Println("int64", v2)
	default:
		fmt.Println("I don't know")
	}

}
