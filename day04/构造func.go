package main

import "fmt"

//构造函数
type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

//构造函数：约定俗称用new开头
//返回的是结构还是结构体指针
//当结构体比较大时，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func newDog(name string) dog {
	return dog{
		name: name,
	}
}
func main() {
	p1 := newPerson("123", 10)
	p2 := newPerson("abc", 20)
	fmt.Println(p1, p2)
	dog1 := newDog("baobao")
	fmt.Println(dog1)

}
