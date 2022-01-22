package main

import "fmt"

//复习结构体
//命名结构体
type person struct {
	name string
	age  int
}

//person的构造函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//person的方法
func (p *person) dream(dream string) {
	fmt.Printf("%s的梦想是：%s\n", p.name, dream)
}

//指针接收者
//1.需要修改结构体变量的值时要使用指针接收者
//2.结构体本身比较大，拷贝的内存开销比较大时也要用指针接收者
//3.保持一致性:如果有一个方法使用了指针接收者，其他的方法为了统一也要使用指针接收者
func (p *person) guonian() {
	p.age++ //此时调用的是P的副本
}
func main() {

	//匿名结构体
	var a = struct {
		x int
		y int
	}{10, 20}

	fmt.Println(a)
	//使用构造方法newperson
	p := newPerson("dhw", 18)
	//fmt.Println(p.age, p.name)
	// p.dream("123")
	// p.dream("567")
	fmt.Println(p.age)
	p.guonian()
	fmt.Println(p.age)

	//结构体嵌套
	type address struct {
		province string
		city     string
	}
	type student struct {
		name    string
		address //匿名嵌套结构体，就使用类型名做名称
	}
}
