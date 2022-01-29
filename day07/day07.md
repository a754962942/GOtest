# day07笔记

## 内容回顾

### time

`2006-01-02 15:04:05.000`

**时间类型**

- `time.Time`:`time.Now`

- 时间戳
  - `time.Now( ).Unix( )`:1970.1.1到现在的秒数
  - `time.Now( ).UnixNano( )`:1970.1.1到现在的纳秒数

**时间间隔类型**

- `time.Duration`:基于int64的一个自订类型

  - `time.Second`

  ```go
  const (
  	Nanosecond  Duration = 1
  	Microsecond          = 1000 * Nanosecond
  	Millisecond          = 1000 * Microsecond
  	Second               = 1000 * Millisecond
  	Minute               = 60 * Second
  	Hour                 = 60 * Minute
  )
  ```

**时间操作**

时间对象+ - 一个时间间隔对象

```go
	// now + 24小时
	fmt.Println(now.Add(24 * time.Hour))
	// Sub 两个时间相减
	nextYear, err := time.Parse("2006-01-02 15:04:05", "2019-08-04 12:25:00")
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	now = now.UTC()
	d := nextYear.Sub(now)
	fmt.Println(d)
```

**时间格式化**

`2006-01-02 15:04:05.000`

**定时器**

```go
// 定时器
	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t) // 1秒钟执行一次
	}
```

**解析字符串时间/时区**

```go
//按照指定格式去解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2022-01-27 20:56:00")
	//按照东八区的时区和格式解析一个字符串格式的时间
	//根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("loadlocation failed,err:%v\n", err)
		return
	}
	//按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-01-27 20:56:00", loc)
	if err != nil {
		fmt.Printf("Parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
```

**反射**

接口类型的变量底层分为两部分：动态类型和动态值。

反射的应用：`json`,`xml`等数据解析、ORM等工具

**反射的两个方法**

`reflect.TypeOf( )`

`reflect.ValueOf( )`

**ini解析 **

## 今日内容

**strconv标准库**

```go
package main

import (
	"fmt"
	"strconv"
)

//strconv
func main() {
	str := "10000"
	//strconv.PaeseInt(转换的字符，进制，类型)
	//返回值为int64
	ret1, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("parseint failed ,err:", err)
		return
	}
	fmt.Printf("%#v,%T\n", ret1, ret1) //10000,int64
	//strconv.Atoi( )返回值类型为Int
	retint, err1 := strconv.Atoi(str)
	if err != nil {
		fmt.Println("AtoI failed,err:", err1)
		return
	}
	fmt.Println(retint)
	//从字符串中解析布尔值
	boolstr := "TRUE"
	boolvalue, err := strconv.ParseBool(boolstr)
	if err != nil {
		fmt.Println("parse Bool failed,err:", err)
		return
	}
	fmt.Println(boolvalue)
	floatstr := "3.1415921"
	fstr, err := strconv.ParseFloat(floatstr, 32)
	if err != nil {
		fmt.Println("parseFloat failed,err:", err)
		return
	}
	fmt.Printf("%#v,%T\n", fstr, fstr)
```



### 并发与并行

**并发：同一时间段执行多个任务**

**并行：同一时刻执行多个任务**

Go语言的并发通过`goroutine`实现，`goroutine`类似于线程，属于用户态的线程，我们可以根据需要，创建成千上万个`goroutine`并发工作。

`goroutine`是由Go 语言的运行时`runtime`调度完成，而线程是由操作系统调度完成。

Go语言还提供`channel`在多个`goroutine`间进行通信。`goroutine`和`channel`是Go语言秉承的CSP并发模式的重要实现基础。

### goroutine(用户态的线程)

```go
//goroutine
func hello(i int) {
	fmt.Println("hello", i)
}

//程序启动之后会创建一个主gorontine去执行
func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i) //用的是函数参数的i，不是外部的i
		}(i)
		//go hello(i)开启一个单独的goroutine去执行hello函数（任务）
	}
	fmt.Println("main")
	time.Sleep(1 * time.Second)
	//main函数结束了，由mian函数启动的goroutine也结束了

}
```

**goroutine什么时候结束**

goroutine对应的函数执行结束了，goroutine就结束了。

`main`函数执行完了，由`main`函数创建的`goroutine`也结束了。

**goroutine与线程**

*可增长的栈*

os线程（操作系统线程）一般有固定的栈内存(通常为2MB)，一个`goroutine`的栈在其生命周期开始时只有很小的栈(典型情况下2KB)

`goroutine`的栈不是固定的，可以按需增加和缩小，`goroutine`的栈的大小限制可以达到1GB.所以在GO语言中一次创建数以万计的`goroutine`是可以的。

### goroutine调度模型

`GMP`是Go语言运行时(runtime)层面的实现，是Go语言自己实现的一套调度系统。区别于操作系统调度OS线程。

- G很好理解，就是这个`goroutine`，里边除了存放本goroutine信息外，还有与所在P的绑定等信息。
- M(Machine)是GO运行时(runtime)对操作系统内核线程的虚拟，M与内核线程一般是一一映射的关系，一个gorontine最终是要凡在M上执行的。
- P管理着一组goroutine队列，P里边会存储当前goroutine运行的上下文环境(函数指针，堆栈地址及地址边界)，P会对自己管理的goroutine队列做一些调度(比如把占用GPU时间较长的goroutine暂停、运行后续的goroutine等)当自己的队列消费完了就去全局队列里取，如果全局队列里的也消费完了会去其他P的队列里拿任务。

P与M一般是一一对应的。他们的关系是：P管理者一组G挂在在M上运行。当一个G长久阻塞在一个M上是，runtime会新建一个M，阻塞G所在的P会把其他的G挂载在新建的M上/当旧的G阻塞完成或者认为其已经死掉时回收旧的M。



**M：N**：把M个goroutine分配给N个操作系统线程去执行。

goroutine初始栈大小大约是2K

**math/rand**

```go
func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) //0<=x<10
		fmt.Println(r1, r2)
	}
}
```



### channel(协调goroutine工作)

```go
var b chan int  //需要指定通道中的元素类型
```

通道必须使用make函数初始化才能使用

#### 通道的操作

`<-`

1. 发送：`chan<-1`
2. 接受：`<-chan`
3. 关闭：`Close(chan)`

### sync

