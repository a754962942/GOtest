package main

import "fmt"

//接口的实现

type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int
}
func (c cat) move() {
	fmt.Println("猫动")
}
func (c cat) eat(s string) {
	fmt.Printf("猫吃%s\n", s)
}
type chichen struct {
	feet int
}

func (c chichen) move() {
	fmt.Println("鸡 MOVE")
}
func (c chichen) eat(s string) {
	fmt.Printf("鸡吃%s\n", s)
}


func main() {
	var a1 animal //定义一个接口类型的变量
	fmt.Printf("%T\n", a1)
	c := cat{ //定义一个cat类型的变量bc
		name: "白",
		feet: 4,
	}
	a1 = c
	a1.eat("鱼")
	fmt.Printf("%T\n", a1)
	fmt.Println(a1)

}
