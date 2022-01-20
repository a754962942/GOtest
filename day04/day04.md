#递归

	//递归:函数自己调用自己
	//递归适合处理那种问题相同或者问题规模越来越小的问题
	//递归一定要有一个明确的退出条件
	//3！ =3*2*1  =3*2！
	//4！ = 4*3*2*1=4*3!
	//5！ =5*4*3*2*1=5*4!
	
	//计算n的阶乘
	
	func digui(n int) int {
		if n <= 1 {
			return 1
		}
		return n * digui(n-1)
	}	

	
	/*
	一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级台阶。
	求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
	
	
	
	解题思路：
	如果有一个台阶，那么只有一个跳法，跳一次
	如果有两个台阶，那么有两种跳法，跳一次一阶或跳一次两阶
	那么n阶台阶，跳一次一阶，剩下n-1阶
	如果开始跳一次两阶剩下n-2阶。
	把jump（n）记为跳台阶的函数
	跳一次一阶，剩下n-1阶，跳法为jump(n-1)
	跳一次两阶,剩下n-2阶，跳法为jump(n-2)
		*/
	
		func f(n int) int {
			if n == 1 {
				//如果只有1个台阶就一种走法
				return 1
			}
			if n == 2 {
				return 2
			}
		return f(n-1) + f(n-2)
	}
	func main() {
		sum := digui(5)
		fmt.Println(sum)
		fmt.Println(f(6))
	}	

#type
**自定义类型和类型别名**

	//自定义类型和类型别名
	//type后边跟的是类型
	type myInt int      //自定义类型
	type youerInt = int //类型别名

	func main() {
		var n myInt
		n = 100
		fmt.Println(n)
		fmt.Printf("%T %d\n ", n, n)
		var m youerInt
		fmt.Println(m)
		fmt.Printf("%T %d\n", m, m)
	
		var c rune
		c = '中'
		fmt.Println(c)
		fmt.Printf("%T %c\n", c, c)
	}
#struct结构体

Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称：struct。

我们可以通过struct来定义自己的类型。

Go语言中通过struct来实现面向对象。

结构体的定义

使用type和struct关键字来定义结构体，如下：

	type 类型名 struct{
		字段名 字段类型
		字段名 字段类型
		……………… ………………
	}
>类型名：表示自定义结构体的名称，在同一个包内不能重复。

>字段名：表示结构体字段名，结构体的字段名必须唯一。

>字段类型：表示结构体字段的具体类型。

**基本实例化**

	type person struct {
		name string
		city string
		age  int8
	}
	
	func main() {
		var p1 person
		p1.name = "沙河娜扎"
		p1.city = "北京"
		p1.age = 18
		fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
		fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}
	}
**通过.来访问结构体的字段(成员变量)，例如p1.name和p1.age**

**匿名结构体**

在定义一些临时数据结构还可以使用匿名结构体

	func main() {
	    var user struct{
			Name string
			Age int
		}
	    user.Name = "小王子"
	    user.Age = 18
	    fmt.Printf("%#v\n", user)
	}
**创建指针类型的结构体**

可以使用new关键字来进行实例化，得到的是结构体的地址。
	
	var p2 = new(person)
	fmt.printf("%T\n",p2)	  //*main.person
	
	fmt.printf("p2=%#v\n",p2) //p2=&main.person{name:"", city:"", age:0}

从打印结果中可以看出p2是一个结构体指针。

**取结构体的地址实例化**

使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作

#构造函数：返回一个结构体变量的函数

**结构体是值类型，赋值的时候都是拷贝**

#方法method

定义：
	
	func（接收者变量 接收者类型）方法名（参数列表）（返回参数）{
	}	
>接收者变量：接收者中的参数变量名在命名之前，建议用接收者类型名的第一个小写字母，例如Person类型应该命名为p

>接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。

>方法名、参数列表、返回参数、具体格式与函数定义相同

		//method
	
	type dog struct {
		name string
	}

	func newDog(name string) dog {
		return dog{
			name: name,
		}
	}

	//方法是用作于特定类型的函数
	//接收者表示的是调用该方法的具体的类型变量，多用类型名首字母小写表示
	func (d dog) wang() {
		fmt.Printf("%s~汪汪汪~", d.name)
	}
	func main() {
		d1 := newDog("团子")
		d1.wang()
	}

#结构体与JSON

