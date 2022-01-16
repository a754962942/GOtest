#变量

----------
GO语言中的变量必须先声明再使用
## 变量的定义
var 变量名 变量类型

例如：var a int //声明一个int类型的变量叫a
 	
var b string //声明一个string类型的变量叫b
>声明的变量必须使用，否则编译报错

**推荐使用小驼峰命名**

例：studentName
## 声明同时赋值
var a string ="abc"
## 类型推导
var a ="abc"//根据值来判断变量是什么类型
## 简短变量声明
a:=123

b:="aala"
## 匿名变量  _
匿名变量不占用命名空间，不分配内存，所以匿名变量之间不存在重复声明。

>函数外的每个语句都必须以关键字开始（var、const、func等）

>:=不能用于函数外

>_多用于占位，表示忽略值

#常量和iota
使用const声明

	例如：const pi = 3.14
	const(
		a=1
		b=2	
	)
##iota
iota是go语言的常量计数器，只能在常量的表达式中使用。

>iota在const关键字出现时被重置为0，const中每增加一行变量声明iota计数一次，使用iota可以简化定义，在定义枚举时很有效

	例const(
		a=iota
		a1=5
		a2=iota
		a3
	)
	print：	a: 0
			a1: 5
			a2: 2
			a3: 3
	经典数量级
		const(
			_=iota
			kb=1<<(10*iota)
			mb
			gb
			tb
			pb
		）