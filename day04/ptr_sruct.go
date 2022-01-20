package main

import "fmt"

//结构体是值类型
//灵宝市防疫政策电话8822689
type person struct {
	name, gender string
}

//函数传参传的是参数的拷贝
//修改的是拷贝，不影响原始值
func f(p person) {
	p.gender = "女"
}

//想要通过函数修改参数，传入该类型的指针
func f1(p *person) {
	//(*p).gender = "女" //根据内存地址找到那个原变量，修改的就是原来的变量
	p.gender = "女" //语法糖，自动根据指针找对应的变量
}
func main() {
	var p person
	p.name = "abv"
	p.gender = "男"
	f(p)
	fmt.Println(p) //男
	f1(&p)
	fmt.Println(p) //女
	//结构体指针1
	var p2 = new(person)
	p2.gender = "女"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  //p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) //求b2的内存地址

	//2.结构体指针

	//2.1key-value初始化
	var p3 = &person{
		name:   "yuanananna",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)
	//2.2使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p4 := &person{
		"小王",
		"男",
	}
	fmt.Println(p4)
}
