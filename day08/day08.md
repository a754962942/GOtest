# day08笔记

# 今日内容

## 并发和并行的区别

并发：同一时刻

## goroutine的启动

将要并发执行的任务包装成一个函数，调用函数的时候前边加上`go`关键字，就能开启一个goroutine去执行该函数的内容。

`goroutine`对应的函数执行完，该goroutine就结束了

main函数启动的时候会自动创建一个goroutine去执行main函数

main函数结束了，那么程序就结束了，由该程序启动的所有goroutine也都结束了。

## goroutine的本质

goroutine的调度模型：`GMP`

`G`:`goroutine`及其绑定P的信息

`M`:`goroutine最终在M上运行`

`P`:`管理goroutine`

**把M个goroutine分配给N个操作系统线程**

## goroutine与操作系统线程(OS线程)的区别

goroutine是用户态的线程，操作系统线程是内核态，goroutine比内核态的线程更轻量级一点。内核态线程初始占用2MB,goroutine初始占用2KB。

可以轻松开启数以万计的goroutine

### runtime.GOMAXPROCS

默认就是操作系统的逻辑核心数，默认跑满CPU

`runtime.GOMAXPROCS(1)`:只占用一个核

### work pool模式

开启一定数量的goroutine去干活

## sync.WaitGroup

```go
var wg sync.WaitGroup
```

- wg.Add(1):计数器+1
- wg.Done()：计数器-1
- wg.Wait()：等待计数器=0

## channel

### 为什么需要channel

通过channel实现多个goroutine之间的通信。

`CSP`:通过通信来共享内存



channel是一种引用类型，使用make函数初始化。(slice、map、channel)



channel的声明：`var ch chan 元素类型`

channel的初始化：`ch=make(chan 元素类型，[缓冲区大小])`

channel的操作：

- 发送：`ch <- 100`
- 接受：`x  := <- ch`
- 关闭：`close(ch)`

### for range

for range 从通道中取值，直到通道关闭才停止取值，如果通道不关闭。死锁。

### select

同一时刻有多个通道要操作的场景下，使用select。

```go
//select
func main() {
	ch := make(chan int, 1)
	for i := 1; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
```

使用`select`语句能提高代码的可读性。

- 可处理一个或多个channel的发送/接受操作
- 如果多个case同时满足，select会随机选择一个。
- 对于没有case的select会一直等待，可用于阻塞main函数

## sync包

### 并发控制

```go
var wg sync.WaitGroup
var x = 0

func add() {
	for i := 0; i < 5000; i++ {
		x += 1
	}
	wg.Done()

}
func main() {
	wg.Add(2)

	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
```



**互斥锁**

互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个`goroutine`可以访问共享资源。Go语言中使用`sync`包中的`Mutex`类型来实现互斥锁。使用互斥锁来修复上边代码的问题：

```go
//锁
var wg sync.WaitGroup
var x = 0

//声明一个sync.Mutex互斥锁，对使用公共资源的部分进行加锁，
//保证公共资源同时只能被一个goroutine调用
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //加锁
		x += 1
		lock.Unlock() //解锁
	}
	wg.Done()

}
func main() {
	wg.Add(2)

	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
```

使用互斥锁能够保证同一时间有且只有一个`goroutine`进入临界区，其他`goroutine`则在等待锁；当互斥锁释放后，等待的`goroutine`才可以获取锁进入临界区，多个`goroutine`同时等待一个锁，唤醒的策略是随机的。

**读写互斥锁**

读写锁分为两种：读锁和写锁。当一个`goroutine`获取读锁之后，其他的`goroutine`如果是获取读锁会继续获得锁，如果是活的写锁就会等待。

当一个`goroutine`获取写锁之后，其他的`goroutine`无论是获取写锁还是读锁都会等待。

```go
/*
rwlock
读要远远大于写的时候才用读写锁
*/
var (
	x     = 0
	lock  sync.Mutex
	wg    sync.WaitGroup
	RLock sync.RWMutex
)

func read() {
	defer wg.Done()
	RLock.RLock()
	// lock.Lock()
	fmt.Println(x)
	// lock.Unlock()
	RLock.RUnlock()
	time.Sleep(time.Millisecond)
}
func write() {
	defer wg.Done()
	// lock.Lock()
	RLock.Lock()
	x += 1
	// lock.Unlock()
	RLock.Unlock()
	time.Sleep(time.Millisecond * 5)
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
```

**sync.Once**

我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

Go语言中的`sync`包中提供了一个针对只执行一次场景的解决方案-`sync.Once`

`sync.Once`只有一个方法`Do`:

```go
func (o *Once) Do(f func()) {}
```

*如果要执行Once方法，传入的参数必须是一个没有参数没有返回值的函数，可以搭配闭包来使用sync.Once.Do*

**sync.Map**

Go语言的`sync`包中提供了一个开箱即用的并发安全版map-`stnc.Map`.开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时`sync.Map`内置诸如`store`、`Load`、`LoadOrStore`、`Delete`、`Range`等方法。

```go
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```



**atomic包**

Go语言中原子操作由内置的标准库sync/atomic包提供



## 网络编程

OSI七层模型：

**应用层**：

作用：规定应用程序使用的数据格式例如TCP协议常见的Email、HTTP、FTP等协议

> 发送方的 HTTP 数据经过互联网的传输过程中会依次添加各层协议的标头信息，接收方收到数据包之后再依次根据协议解包得到数据
>
> ![](D:\GO\src\github.com\a754962942\Gotest\day08\httptcpip.png)

**表示层**：

**会话层**：

**传输层**（端口）：表示这个数据包到底供哪个程序（进程）使用。IP地址+MAC地址+端口号

**网络层**（IP协议）：网络层引出IP，网络地址和MAC地址没有任何联系。网络地址帮助我们确定计算机所在的子网络，MAC地址则将数据包发送到该子网络中的目标网卡。

**数据链路层**（以太网协议）：以太网规定，一组电信号构成一个数据包，叫做帧。每一帧分为标头和数据两部分。 如果数据很长，就分割为多个帧发送。

数据包必须从一个网卡传送到另一个网卡。网卡地址就是数据包的发送地址和接受地址。叫做MAC地址。

通过ARP协议获取接收方的MAC地址，有了MAC地址之后以太网就向本网络的所有计算机广播，每台计算机读取数据包标头，找到接收方的MAC地址，然后与自身的MAC地址相比较，如果相同就接受这个包，否则就丢弃。

**物理层**：规定了网络的一些电气特性，作用是负责传送0和1的电信号





