package main

import "fmt"

//同一个结构体可以实现多个接口

//接口还可以嵌套
type animal interface{
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat(string)
}
type cat struct {
	name string
	feet int
}

func (c *cat) move() {
	fmt.Println("猫走路")
}
func (c *cat) eat(s string) {
	fmt.Printf("猫吃%s",s)
}
func main() {
	var 
	bc :=cat{
		name: "蓝猫",
		feet: 4,
	}

}
