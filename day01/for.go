package main

import "fmt"

//for循环
func main() {

	//基本格式
	//for i := 0; i <= 9; i++ {
	//	fmt.Println(i)
	//}
	//变种
	var i = 5
	for ; i < 10; i++ {
		fmt.Println("i=", i)
	}
	//死循环
	//for{
	//
	//}

	//for range(键值循环)
	s := "Hello"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
