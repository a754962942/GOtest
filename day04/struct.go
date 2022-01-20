package main

import "fmt"

//结构体
type person struct {
	name  string
	age   int
	sex   string
	hobby []string
}

func main() {
	//声明一个person类型的变量people
	var people person
	//通过字段赋值
	people.name = "abc"
	people.age = 100
	people.sex = "male"
	people.hobby = []string{"旅游"}
	fmt.Println(people)
	//访问变量peopel的字段
	fmt.Println(people.name)
	var p1 person
	p1.name = "bbb"
	p1.age = 20
	fmt.Println("age:", p1.age)
	//匿名结构体
	var s struct {
		name string
		age  int
		city string
	}
	s.name = "hhh"
	s.age = 18
	s.city = "beijing"
	fmt.Printf("type :%T value: %v\n", s, s)
}
