package main

import (
	"fmt"
	"time"
)

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
