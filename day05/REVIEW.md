#回顾

**自定义类型和类型别名**

	type Myint int //自定义类型
	type newInt = int //类型别名
类型别名只在代码编写过程中有效，编译完成之后就不存在，内置的byte和rune都属于类型别名
 
**结构体**

基本数据类型:表示现实中的物件有局限性。
编程是代码解决现实生活中的问题。

	var name = "万宝路"
结构体是一种数据类型，一种我们自己造的可以保存多个维度数据的类型。
	
	type person struct{
	 		name string
			age  int
			addr string
	}

	
	
匿名结构体

多用于临时场景

	var a = struct {
		x int
	}

**构造函数**



#方法和接收者
>方法是有接收者的函数，接收者指的是哪个类型的变量可以调用这个函数
	
	type person struct {
		name string
		age  int
	}
	func (p person) dream(dream string) {
		fmt.Printf("%s的梦想是：%s\n", p.name, dream)
	}
结构体是值类型
**结构体的嵌套**

	//结构体嵌套
	type address struct{
		province string
		city string
	}
	type student struct{
		name string
		address address //嵌套别的结构体
	}

**结构体的匿名字段**

	//结构体嵌套
	type address struct{
		province string
		city string
	}
	type student struct{
		name string
		 address //匿名嵌套结构体，就使用类型名做名称
	}

**JSON序列化与反序列化**
>1. 结构体内部字段首字母要大写，不大写在外部包访问不到

>2. 反序列化时要传递指针。
	
	package main
	
	import (
		"encoding/json"
		"fmt"
	)
	
	//序列化和反序列化
	func main() {
		type point struct {
			X int `json:"x"`
			Y int `json:"y"`
		}
		p1 := point{100, 200}
		b, err := json.Marshal(p1)
		if err != nil {
			fmt.Printf("Marshal failed, err:%v\n", err)
			}
		fmt.Println(string(b))
		
		//反序列化：由字符串-->Go中的结构体变量
		str1 := `{"x":100,"y":123}`
		str2 := point{} //造一个结构体变量，准备接受反序列化的值
		err = json.Unmarshal([]byte(str1), &str2)
		if err != nil {
			fmt.Printf("Ummarshal failed,err:%v\n", err)
			}
		fmt.Println(str2)
	}