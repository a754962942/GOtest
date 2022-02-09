package main

import (
	"fmt"
	"sync/atomic"
)

//原子操作

var x int64

// var wg sync.WaitGroup
// var lock sync.Mutex

func add() {
	// lock.Lock()
	// x++
	// lock.Unlock()
	//原子操作，在原有变量上增加值
	// atomic.AddInt64(&x, 1)

	// wg.Done()
}
func main() {
	// for i := 0; i < 1000000; i++ {
	// 	wg.Add(1)
	// 	go add()
	// }
	// wg.Wait()
	x = 200
	//比较并交换
	//拿old的值和变量值进行比较，如果相等，就返回true，并替换new的值，如果不相等，则返回false
	ok := atomic.CompareAndSwapInt64(&x, 200, 100)
	fmt.Println(x, ok)
}
