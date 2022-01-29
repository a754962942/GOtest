package main

import (
	"fmt"
	"runtime"
	"sync"
)

//GOMAXPROCS

var wg sync.WaitGroup

func A() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}
func B() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
}
