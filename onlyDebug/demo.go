package main

import (
	"fmt"
	"time"
)

func spinner(d time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(d)
		}
	}
}
func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {

		return 2
	}
	return fib(n-1) + fib(n-2)
}
func main() {
	go spinner(time.Millisecond * 100)
	fmt.Printf("\n%d\n", fib(3))
}
