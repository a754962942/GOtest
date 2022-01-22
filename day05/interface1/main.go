package main

import "fmt"

//接口实例2
//不管什么接口的车，都能跑

//不管什么结构体，只要有run方法都能视car类型
type car interface {
	run()
}
type BMW struct {
	brand string
}
type Benz struct {
	brand string
}

func (b BMW) run() {
	fmt.Printf("%s 速度100KM/H\n", b.brand)
}
func (b Benz) run() {
	fmt.Printf("%s 速度120KM/H\n", b.brand)
}

//drive函数接受一个car类型的变量

func drive(c car) {
	c.run()
}
func main() {
	var b1 = BMW{
		brand: "宝马",
	}
	b2 := Benz{
		brand: "奔驰",
	}
	drive(b1)
	drive(b2)
}
