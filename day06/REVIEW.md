# 内容回顾

# package 
包的定义--->package关键字

- 一个文件夹就是一个包
- 文件夹里边放的都是.go文件


包的导入--->import

- 导入路径从*$GOPATH/src*后边的路径开始写起
- 单行导入
- 多行导入
- 给导入的包起别名
- 匿名导入

包中的标识符(变量名/函数名/结构体名/接口名/常量...)、

可见性(标识符首字母大写：对外可见)

Go语言不支持循环导入

**初始化函数init()**

- 包导入时自动执行
- 一个包里边只有一个init()
- init()没有参数也没有返回值也不能调用
- 多个包init的执行顺序
- 一般用于做一些初始化操作..

# interface

接口是一种类型，一种抽象的类型

接口就是你要实现方法的一个清单。

接口的定义

	type mover interface{
		方法的签名(参数)(返回值)
	}
接口的实现
>实现了接口的所有方法就实现了这个接口
>
>实现了接口就可以当成这个接口类型的变量

接口变量
>实现了一个万能的变量，可以保存所有实现了我这个接口的类型的值。
>
>通常是作为函数的参数出现的


空接口
>接口中没有定义任何方法，也可以说任何类型都实现了空接口

>任何类型都可以存到空接口变量中作为函数参数


**接口底层**

- 动态类型
- 动态值

类型断言

做类型断言的前提是一定要是一个接口类型的变量。

x.(T)
使用switch来做类型判断

	//类型断言
	func main() {
		var a interface{} //定义一个空接口变量a。
	
		//如何判断a保存的值的具体类型是什么？
		//类型断言
		//1.x.(T)
		a = 100
		// v1, ok := a.(int8)
		// if ok {
		// 	fmt.Println("v1是int8")
		// }
		//2.switch
		switch v2 := a.(type) {
		case int8:
			fmt.Println("int8", v2)
		case int32:
			fmt.Println("int32", v2)
		case string:
			fmt.Println("string", v2)
		case int:
			fmt.Println("int64", v2)
		default:
			fmt.Println("I don't know")
		}
	
	}

# 文件操作
**打开文件和关闭文件**

	fileObj,err:=os.Open() //返回一个*File和err
	defer fileObj.Close()
**读文件**

	fileObj.read()
	bufio
	ioutil

**写文件**
	
	os.OpenFile(name string,flag int,perm FileMode)(*File,err){
	}

- name:要打开的文件名
- flag:打开文件的模式。
	- 以下几种：
	- os.O_WRONLY   只写
	- os.O_CREATE   创建文件
	- os.O_RDONLY   只读
	- os.O_RDWR     读写
	- os.O_TRUNC	清空
	- os.O_APPEND	追加
	
- perm:文件权限，一个八进制数。R(读)04，W(写)02，X(执行)01.

