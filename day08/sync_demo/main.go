package main

import (
	"fmt"
	"sync"
)

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
