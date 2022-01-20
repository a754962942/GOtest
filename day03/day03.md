#defer
		
	//defer
	//多用于函数结束之前释放资源(文件句柄、数据库连接、socket链接等)
	//多个defer同时调用时，执行顺序为后定义的defer先执行，顺序为后进先出。
	
	func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("aaaa")//defer 把它后边的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("bbb")
	defer fmt.Println("ccc")
	fmt.Println("end")
	}
	//defer执行时机，如果函数中有
	func main() {
	deferDemo()
	}

#defer和return配合使用
		
	//defer执行顺序
	//Go语言中函数的return不是原子操作，在底层分两步执行
	//第一步：返回值赋值
	//defer
	//第二步：真正的return返回
	//函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
	
	func f1() int {
		x := 5
		defer func() {
			x++ //修改的是X 不是返回值
		}()
		return x // 第一步 返回值= 5 第二步 defer x=6 第三步 RET=5
	}
	func f2() (x int) {
		defer func() {
			x++
		}()
		return 5 //分为两步， 返回值=x    defer  x++   RET x=6
		}
	func f3() (y int) {	
		x := 5
		defer func() {
			x++
		}()
		return x //返回值 = y = x =5 defer x++ RET = 5
		}
	func f4() (x int) {
		defer func(x int) {
			x++
		}(x)
		return 5 //函数传参修改更改的是参数的拷贝，不影响原始值 RET = 5
		}
	func main() {
		fmt.Println(f1())//5
		fmt.Println(f2())//6
		fmt.Println(f3())//5
		fmt.Println(f4())//5
		}
#作用域

		//变量的作用域
		//1.全局作用域
		var x = 100 //定义一个全局变量
		//定义一个函数
		
		func f1() {
		x := 10
		name := "123"
		//函数中查找变量的顺序
		//1.先在函数内部查找。
		//2.找不到就往函数外部查找，一直找到全局
		//如果全局没有找到，报错
	
		fmt.Println(x, name)
	
		}
		func main() {
		f1()
		//2.函数作用域
		//函数内部定义的变量只能在该函数内部使用
		//fmt.Println(name)
	
		//3.语句块作用域
		if i := 10; i < 18 {
			fmt.Println("乖乖上学")
		}
		//fmt.Println(i)
		for j := 0; j < count; j++ {
			fmt.Println("123")
		}
		//fmt.Println(j)
	}

#函数
**函数的定义**
基本格式

参数的类型

有参数的函数

参数类型简写

可变参数

返回值的格式

有返回值

多返回值

命名返回值

**变量的作用域**

1. 全局作用域
2. 函数作用域
>先在函数内部找变量，找不到往外层找
>
>函数内部的变量，外部是访问不到的


3. 代码块作用域

**高阶函数**

函数也是一种类型，它可以作为参数，也可以作为返回值


#匿名函数
没有名字的函数， 定义格式如下：

	func (参数)(返回值){
		函数体
	}
匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或作为立即执行函数

**多用于实现回调函数和闭包**

#闭包
闭包是指一个函数和与其相关的引用环境组合而成的实体。

**闭包=函数+引用环境**

例：
	//闭包是什么？
	//闭包是一个函数，这个函数包含了它外部作用域的一个变量
		
	//底层原理：
	//1.函数可以作为返回值
	//2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找
	func adder(x int) func(int) int { //闭包等于函数+外部变量的引用
		return func(y int) int {
			x += y
			return x
		}
	}
	func main() {
		ret := adder(1)
		ret2 := ret(200)
		ret3 :=ret(1)
		fmt.Println(ret2)
		fmt.Println(ret3)
	}

**不懂看这里**

	//闭包
	//闭包原理
	//需要将f2传入f1进行调用
	func f1(f func()) { //f1函数只能接受一个没有参数没有返回值的函数
		f()	
	}
	func f2(x, y int) { //f1函数有两个int型的参数
		fmt.Println("This is F2")
		fmt.Println(x + y)
	}

	//定义一个函数对f2进行包装
	func f3(f func(int,int),x,y int)func() {
		ret := func ()  {
			f(x,y)
		}
		return ret
	}
	func main() {
		ret :=f3(f2,100,200)
		//查看ret类型
		//fmt.Printf("%T",ret)
		//实现闭包调用
		f1(ret)
	}
	
	
#内置函数
>close  用来关闭channel

>len 用来求长度，比如string、array、slice、map、channel

>new 用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针

>make 用来分配内存，主要用来分配引用类型，比如map、chan、slice

>append 用来追加元素到数组、slice中

>panic和recover 用来处理错误

#panic/recover

**Go语言中目前是没有异常机制，但是使用panic和recover模式来处理错误。**

**panic可以在任何地方引发，但recover只有在defer调用的函数中有效**

	//panic和recover
	func funcA() {
		fmt.Println("This is A")
		}
	func funcB() {
		defer func(){
			err:=recover()
			fmt.Println("In defer")
		}()
		panic("出错了")//程序崩溃推出
		fmt.Println("This is B")
	
		}
		func funcC() {
		fmt.Println("This is C")
		}
		func main() {
			funcA()
			funcB()
			funcC()
		}

1. **recover()必须搭配defer使用。**
2. **defer一定要在可能引发panic的语句之前定义。**