package main

import "fmt"

//使用值接收者和指针接收者的区别

type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int
}

// //使用值接收者实现了接口的所有方法
// func (c cat) move() {
// 	fmt.Println("猫动")
// }

// func (c cat) eat(s string) {
// 	fmt.Printf("猫吃%s\n", s)
// }
//使用指针接收者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("猫动")
}

func (c *cat) eat(s string) {
	fmt.Printf("猫吃%s\n", s)
}
func main() {
	var a1 animal
	c1 := cat{name: "TOM", feet: 4}
	c2 := &cat{name: "MIKE", feet: 4}
	a1 = &c1   //实现animal这个接口使用的是指针接收者
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)

}
