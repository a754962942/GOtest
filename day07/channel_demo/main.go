package main

import (
	"fmt"
	"sync"
)

//channel练习
/*
1.启动一个goroutine，生成100个数发送到ch1
2.启动一个goroutine，从ch1取值，计算平方放到ch2中
3.在main函数 从ch2取值
*/
var ch1 chan int
var ch2 chan int
var wg sync.WaitGroup

//传入一个只能存的单向通道
func f1(ch1 chan<- int) {
	defer wg.Done()
	defer close(ch1)
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
}

//ch1是一个只能取的单向通道
//ch2是一个只能存的单向通道
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	defer close(ch2)
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}

}
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg.Add(2)
	go f1(ch1)
	go f2(ch1, ch2)
	wg.Wait()
	for i := range ch2 {
		fmt.Println(i)
	}
}
