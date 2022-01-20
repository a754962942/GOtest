package main

import (
	"fmt"
)

//自定义类型和类型别名
//type后边跟的是类型
type myInt int      //自定义类型
type youerInt = int //类型别名
func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T %d\n ", n, n)
	var m youerInt
	fmt.Println(m)
	fmt.Printf("%T %d\n", m, m)

	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T %c\n", c, c)
}
