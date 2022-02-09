package main

import (
	"fmt"
	"sync"
	"time"
)

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
