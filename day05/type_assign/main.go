package main

import (
	"fmt"
)

//类型断言目的是想知道空姐口中接受的值的类型

//类型断言1

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串", str)
	}
}

//类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串", t)
	case int:
		fmt.Println("是一个Int", t)
	case bool:
		fmt.Println("是一个布尔值", t)

	}
}
func main() {
	assign(100)
	assign2("哈哈哈")
	assign2(100)
	assign(true)
}
