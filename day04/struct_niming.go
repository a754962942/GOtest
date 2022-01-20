package main

import "fmt"

//匿名字段
//适用于字段比较少也比较简单的场景
//不常用
type address struct {
	city     string
	province string
}
type person struct {
	name    string
	age     int
	address //匿名嵌套结构体
}
type company struct {
	name string
	addr address
}

func main() {
	// p := person{
	// 	"sbv",
	// 	28,
	// }
	// fmt.Println(p)
	// fmt.Println(p.string)
	// fmt.Println(p.int)
	p1 := person{
		name: "abc",
		age:  10,
		address: address{
			province: "河南",
			city:     "三门峡",
		},
	}
	p2 := company{
		name: "OKEX",
		addr: address{
			province: "北京",
			city:     "北京",
		},
	}
	fmt.Println(p1.city)
	fmt.Println(p2, p2.name)
}
