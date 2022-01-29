package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//waitGroup

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) //0<=x<10
		fmt.Println(r1, r2)
	}
}
func f1(i int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
	defer wg.Done()
}
//同步goroutine
var wg sync.WaitGroup

func main() {
	// f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait() //等待wg计数器等于0
}
