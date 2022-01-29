package main

import (
	"fmt"
	"sync"
)

//channel
var (
	a  []int
	b  chan int
	wg sync.WaitGroup
)

func NobufChannel() {
	fmt.Println(b)     //nil
	b = make(chan int) //初始化一个不带缓冲区的通道
	wg.Add(1)
	go func() {
		wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到值", x)
	}()
	b <- 10
	fmt.Println("10发送到通道B中了")
	wg.Wait()

}
func BufChannel() {
	fmt.Println(b)        //nil
	b = make(chan int, 2) //初始化一个不带缓冲区的通道
	b <- 10
	fmt.Println("10发送到通道B中了")
	b <- 20
	fmt.Println("20发送到通道B中了")
	x := <-b
	fmt.Println("从通道B中取到了", x)
	y := <-b
	fmt.Println("从通道B中取到了", y)

}
func main() {
	BufChannel()
	// b = make(chan int, 16) //初始化一个带缓冲区的通道
	// fmt.Println(b)
}
