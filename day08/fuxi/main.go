package main

import "fmt"

//channel
func main() {
	ch := make(chan int, 1)
	ch <- 100
	<-ch
	x, ok := <-ch
	fmt.Println(x, ok)
}
