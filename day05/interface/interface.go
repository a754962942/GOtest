package main

import "fmt"

//引出接口的实例

type speaker interface {
	
	speak() //只要实现了speak()方法，都是speaker类型,方法签名
}
type cat struct {
}
type dog struct {
}
type person struct{}

func (p person) speak() {
	fmt.Println("嘿嘿嘿")
}
func (c cat) speak() {
	fmt.Println("喵喵喵")
}
func (d dog) speak() {
	fmt.Println("汪汪汪")
}
func da(x speaker) {
	//接受一个参数，传进来什么输出什么
	x.speak()
}
func main() {
	c1 := cat{}
	d1 := dog{}
	da(c1)
	da(d1)
	p1 := person{}
	da(p1)
	var ss speaker //定义一个接口类型：sperker 的变量:ss
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)
}
