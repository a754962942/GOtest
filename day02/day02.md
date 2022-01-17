#switch简化大量的判断
	
	var n = 2
	switch n{ 
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	}

	//当分支使用表达式时，switch语句后就不需要跟判断变量，例如
	func switchAge(){
	age:=30
	switch {
	case age<25:
		fmt.Println("好好学习")
	case age>25&&age<35:
		fmt.Println("工作")
	case age>60:
		fmt.Println("享受")
	default:
		fmt.Println("好")
	}
	}
	//如果switch后有判断变量，那么分支就不用使用表达式
	func switchCase(n int){
		switch n {
		case 1,3,5,7,9:
			fmt.Println("奇数")
		case 2,4,6,8:
			fmt.Println("偶数")
		default:
			fmt.Println("ERROR")
		}
	}
	//fallthrough语法可以执行满足条件的下一个case，是为了兼容C语言中的case设计的
	func switchDemo(){
		s:="a"
		switch  {
		case s == "a":
			fmt.Println("a")
			fallthrough
		case s == "b":
			fmt.Println("b")
		case s =="c":
			fmt.Println("c")
		}
	}
#goto语句（跳转到指定标签）
**goto语句通过标签进行代码见的无条件转换。**
goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用goto语句能简化一些代码的实现过程。例如：双层嵌套的for循环要退出时
	
	func gotoDemo1(){
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if j ==4{
					goto Label //跳到Label标签
				}
				fmt.Printf("%d - %d \n",i,j)
			}
		}
		Label: //Label标签
			fmt.Println("This is Label")
	}
#运算符
GO语言内置的运算符有：

1. 算术运算符
2. 关系运算符
3. 逻辑运算符
4. 位运算符
5. 赋值运算符

##算数运算符
**+**	相加

**-**	相减

***	相乘

/	除

%	求余

**注意 ++（自增）和--（自减）在Go语言中时单独的语句，并不是运算符。**
##关系运算符
**==**	检查两个值是否相等，如果相等返回True否则返回False

**！=**	检查两个值是否不等，如果不等返回True，否则返回False

**>**

**<**

**<=**

**>=**
##逻辑运算符
**&&**	逻辑运算符AND，如果两边的操作数都是True，则为True，否则为False

**||**	逻辑运算符OR，如果两边的操作数有一个为True，则为True，否则为False

**！**	逻辑运算符NOT，如果条件为True，则为False，否则为True

##位运算符
**&**	参与运算的两数各对应的二进制位相与。(两数均位1才算1）

**|**	参与运算的两数各对应的二进制相或。（两位有一个为1就为1）

**^**	参与运算的两数各对应位二进位相异或，当两对应的二进位相异时，结果为1.（两位不一样则为1）

**<<**	左移n位就是乘以2的n次方（a<<b就是把a的各二进位全部左移b位，高位丢弃，低位补零）

**>>**	右移n位就是除以2的n次方（a>>b就是把a的各二进位全部右移b位）

##赋值运算符
	=
	+=
	-=
	\=
	*=
	%=
	<<=
	|=
	&=
	>>=
	^=
#复合数据类型
##Array（数组）
数组是同一种数据类型元素的集合，在Go语言中，数组从声明是就确定，使用时可以修改数组成员，但是数组大小不可变化，基本语法：
	
	//定义一个长度为3元素类型为int的数组a
	var a [3]int

**数组定义：**

	var 数组变量名 [元素数量]Type
比如var a[5]int ,数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。[5]int 和[10]int是不同的类型
	
	var a [3]int 
	var b [4]int 
	a = b //不能这么做，因为此时a和b是不同的类型
数组可以通过下标进行访问，下标是从0开始的，最后一个下标是len-1,访问越界会触发panic

**数组的初始化**

数组初始化有很多方式

	var a1 [3]bool 
	var a2 [4]bool 
	//如果不初始化，默认元素都是零值（布尔值：false，整数、浮点数：0，字符串："")
	fmt.Println(a1,a2)
	//方式1:初始化数字组时可以使用初始化列表来设置数组元素的值
	a1 = [3]bool{true,true,true}
	//方式2:根据初始值自动推断数组长度
	a3 := [...]int {0,1,2,3,4,5,6,7,8}
	fmt.Println(len(a3))
	//方法3：根据索引来初始化
	a4:=[5]int{0:1,4:2}
	fmt.Println(a4)
**数组的遍历**

	//数组遍历
	//方法1：按照index取
	citys:=[...]string{"北京","上海","广州","成都"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2.for range遍历
	for i,v :=range citys{
		fmt.Printf("%d:%s\n",i,v)
	}

**多维数组**
	
	//a[3][2]int 
	//[[1,2][3,4][5,6]]
	多维数组只有第一层能使用...来让编译器推导数组长度。
	a := [...][2]int{
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6},
	}
	//fmt.Println(a)
	//二维数组的遍历
	for _,v:=range a{
		fmt.Printf("%v\n",v)
	}
**数组是值类型**

**值类型与引用类型详解**

> https://studygolang.com/articles/33388

数组支持“==”、“!=”操作符，因为内存总是被初始化过的

[n]*T表示指针数组， *[n]T表示数组指针
#切片（slice）
引言

因为数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性。例如：
	
	func arrauSum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
	}
这个求和函数只能接受[3]int类型，其他的都不支持。再比如，
	
	arr := [3]int{1, 2, 3}
数组arr中已经有三个元素了，我们不能在继续往数组arr中添加新元素了。

**切片**
切片(slice)是一个拥有相同类型元素的可变长度的序列。

它是基于数组类型做的一层封装。**支持自动扩容**

**切片是一个引用类型，它内部包括地址、长度、容量。**切片一般用于快速地操作一块数据合计

**切片的定义**

声明切片类型基本语法为：

	var name []T
其中，

- name:表示变量名
- T：表示切片中的元素类型

##make

**切片的本质**

切片就是一个框，框住了一块连续的内存。属于引用类型，真正的数据都是保存在底层数组里的。

**切片不能直接比较**
切片之间不能比较，不能用==操作符来判断两个切片是否含有全部相等元素。

切片唯一合法的比较操作是和nil比较。

**一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是零**

**但不能说一个长度和容量都为零的切片一定是nil**，例：
	
	//len(s1)=0;cap(s1)=0;s1==nil
	var s1[]int
	//len(s2)=0;cap(s2)=0;s2!=nil
	s2 := []int{}
	//len(s3)=0;cap(s3)=0;s3!=nil
	s3 := make([]int,0)
所以要判断一个切片是否是空的，要用len(s)=0来判断。
**Slice遍历**

	//切片的遍历
	s:=[]int{1,2,3}
	//1.索引遍历
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	//for range遍历
	for i,v:=range s{
		fmt.Printf("index:%v\tvalue:%v\n",i,v)
	}

#append
**append使用格式**

	s1 := []string{"北京", "上海", "深圳"}
	s1 = append(s1, "日本")
	fmt.Println("s1:",s1) //s1: [北京 上海 深圳 日本]
**append(目标slice，追加内容)**

*调用append函数必须用原来的切片变量接受返回值*

**append导致slice.cap超出，那么系统会重新分配一块连续的内存，cap值为原来的2倍**例如：

	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1:%v,len(s1):%v,cap(s1):%v\n", s1,len(s1),cap(s1))
	s1 = append(s1, "日本")
	fmt.Printf("s1:%v,len(s1):%v,cap(s1):%v\n",s1,len(s1),cap(s1))

	//s1:[北京 上海 深圳],len(s1):3,cap(s1):3
	//s1:[北京 上海 深圳 日本],len(s1):4,cap(s1):6

	s1=append(s1, s2...)//使用...可以将slice拆分开进行添加
	fmt.Printf("s1:%v,len(s1):%v,cap(s1):%v\n",s1,len(s1),cap(s1))
	//s1:[北京 上海 深圳 日本 武汉 西安 杭州],len(s1):7,cap(s1):12

**关于append删除切片中某个元素**

	//关于append删除切片中的某个元素
	a1 := [...]int{1, 2, 3, 4, 5, 6, 8, 9, 9, 05, 3}
	s1 := a1[:]

	//避开索引为1的哪个2
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1)
	//[1 3 4 5 6 8 9 9 5 3]
	//[1 3 4 5 6 8 9 9 5 3 3]
#copy

	a:=[]int{1,3,5}
	b:=a1
	c:=make([]int, 3,3)
	copy(c,a1)
	fmt.Println(a,b,c)//[1 3 5] [1 3 5] [1 3 5]	
	b[0]=100
	fmt.Println(a,b,c)//[100 3 5] [100 3 5] [1 3 5]
由于切片是引用类型，所以a和b都指向了同一个内存。修改b的同时a的值也会发生变化。

Go语言内建的copy()函数可以迅速将一个切片复制到另外一个切片空间中，copy()函数的使用格式如下：

	copy(destSlice,srcSlice)

- srcSlice:数据来源切片
- destSlice:目标切片

Go语言中没有删除切片的方法
#指针
Go语言中不存在指针操作，只需要记住两个符号：

1. & 	取地址
2. *** 	根据地址取值

 #make和new的区别
**1.make和new都是用来申请内存的**

**2.new很少用，一般用来给基本数据类型申请内存，如int/string...，返回的是对应类型的指针（*int、*string）**

**3.make是用来给Slice、map、chan申请内存的，make函数返回的是这三个类型本身**
#map
Go语言中提供映射关系容器为map，其内部使用散列表（hash）实现。

map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
#函数

