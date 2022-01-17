# day01回顾
## GO命令
	go build 编译GO程序
	go build -o "XX.EXE":编译为XX.EXE文件
	go run main.go:像脚本一样执行main.go文件
	go install：先编译后拷贝
##GO语言文件基础语法
	存放GO源代码的文件后缀名为.go
	文件第一行：package 关键字声明包名
	// 单行注释方法
	/* 多行注释方法
	如果要编译一个可执行文件，必须声明main包，并且要有一个main函数
	func main(){}
main函数是入口函数，他没有参数也没有返回值

**GO语言函数外的语句必须以关键字开头**
函数内部定义的变量必须使用
## 变量
 	3种声明方式
	1.var name string
	2.var name2 = "123"//类型推断
	3.函数内部专属：name:="abcd"

**匿名变量**
当有些数据必须用变量接受但是又不使用时，就可以用_来接受这个值。
## 常量
	const PI = 3.1415926
	const UserNotExistErr = 10000

iota:实现枚举
两个要点：
1.iota 在const关键字出现时将被重置为0
2.const中每新增一个常量声明，iota累加1
##流程控制
### if
	var age = 19
	if age > 20 {
		fmt.Println("成年了")
	}else if age>10{
		fmt.Println("未成年人")
	}else{
		fmt.Println("小孩子")
	}
##for循环
标准for循环
	
	for i := 0; i < 10; i++ {
	fmt.Println(i)
	}

for range

	s:="hello"
	for i,v:=range s{
	fmt.Printf("index:%d,value:%c\n",i,v)
	}
##基本数据类型
整形

无符号整型：uint8,uint16,uint32,uint64

有符号整型：int8,int16,int32,int64

int:具体是32位还是64位看操作系统

uintptr:表示指针
##其他进制数
Go语言中没办法直接定义二进制数

	var a = 077
	var b = 0xff
	fmt.Println(a,b)   //63  255
	fmt.Printf("%o\n",a) //输出八进制数77
	fmt.Printf("%x\n",b) //输出十六进制数ff

##浮点型：

float32和float64

Go语言中浮点数默认是float64

##复数

	complex128和complex64
##布尔值
true 和false

不能和其他的类型做转换
##字符串
常用方法

字符串不能修改
##byte和rune类型
都属于类型别名

##字符串、字符、字节都是什么？
字符串：双引号包裹的是字符串

字符：单引号包裹的是字符，单个字母，单个字符，单个文字

字节：1byte = 8Bit

Go语言中字符串都是UTF-8编码，一个汉字一般占用3个字节。
