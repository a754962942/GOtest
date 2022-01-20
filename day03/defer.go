package main

import "fmt"

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
