#接口(interface)

接口是一种类型:是一种特殊的类型，约束了变量有哪些方法

在编程中，会遇到以下场景：
>不关心一个变量是什么类型，只关心能调用它的什么方法

**接口的定义**

	type 接口名 interface{
		方法名1(参数1，参数2 ...)(返回值1,返回值2...)
		方法名2(参数1，参数2 ...)(返回值1,返回值2...)
		...
	}
用来给变量、参数、返回值等设置类型

**接口在底层分为两个部分：第一部分是值的类型，第二个部分是值本身**

**接口的实现**

>一个变量如果实现了接口中规定的所有的方法，那么这个变量就实现了这个接口，可以称为为这个接口类型的变量


**值接收者和指针接收者**

使用值接收者实现接口与使用指针接收者实现接口的区别？

使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存。

指针接收者实现的接口只能存指针类型的变量。

**接口和类型的关系**

多个类型可以实现同一个接口，

一个类型可以实现多个接口

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
**空接口**

	type xx interface {}
如果一个变量实现了接口中规定的所有的方法，那么这个变量就实现了这个接口。

空接口没有必要起名字，通常定义为下面的格式：

	interface{}

*意味着所有类型都实现了空接口，也就是任意类型的变量都能保存到空接口中*

**类型断言**

	//类型断言目的是想知道空接口中接受的值的类型
	
	//类型断言1
	
	func assign(a interface{}) {
		fmt.Printf("%T\n", a)
		str, ok := a.(string)
		if !ok {
			fmt.Println("猜错了")
		} else {
			fmt.Println("传进来的是一个字符串", str)
		}
	}
	
	//类型断言2
	func assign2(a interface{}) {
		fmt.Printf("%T\n", a)
		switch t := a.(type) {
		case string:
			fmt.Println("是一个字符串", t)
		case int:
			fmt.Println("是一个Int", t)
		case bool:
			fmt.Println("是一个布尔值", t)
	
		}
	}
	func main() {
		assign(100)
		assign2("哈哈哈")
		assign2(100)
		assign(true)
	}
#包(package)
>包的路径从“GOPATH/src”后边的路径开始写起，路径分隔符用**/**

>想被别的包调用的标识符都要首字母大写！

>单行导入和多行导入

>导入包的时候可以指定别名

>导入包不想使用包内部的标识符，需要使用匿名导入

>每个包导入的时候会自动执行一个名为init()的函数，他没有参数也没有返回值，也不能手动调用



	import (
		"fmt"
	//从GOPATH/src路径下开始
		abc "github.com/a754962942/Gotest/day05/package"
	)
	
	func main() {
		ret := abc.Add(1, 2)
		fmt.Println(ret)
	}

**init()初始化函数**
*init()函数介绍*
在Go语言程序执行时导入包语句会自动触发包内部init()函数的调用。

init()函数没有参数也没有返回值。

init()函数在程序运行时自动被调用执行，不能再代码中主动调用它。

	包中init函数的执行时机
		全局声明
			|
		  init()
			|
		  main()

#文件操作

**打开和关闭文件**

>os.Open()函数能够打开一个文件，返回一个*File和一个err。

>对得到的文件实例调用close()方法能够关闭文件

**为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句**

	package main
	
	import (
		"fmt"
		"os"
	)
	
	//文件操作
	
	//打开文件
	func main() {

	filObj, err := os.Open("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\REVIEW.md")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//最后关闭文件
	defer filObj.Close()
	//读文件
	// var tmp = make([]byte,128)//指定读的长度
	var tmp [128]byte
	for {
		n, err := filObj.Read(tmp[:])
		if err != nil {
			fmt.Printf("read  from file failed,err :%v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
			}
		}
	}