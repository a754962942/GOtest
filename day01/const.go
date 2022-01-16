package main

import "fmt"

const(
	a=iota
	a1=5
	a2=iota
	a3
)
const(
	_=iota
	kb=1<<(10*iota)
	mb
	gb
	tb
	pb
)
func main() {
	fmt.Println("a:",a)
	fmt.Println("a1:",a1)
	fmt.Println("a2:",a2)
	fmt.Println("a3:",a3)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("kb:",kb)
	fmt.Println("mb:",mb)
	fmt.Println("gb:",gb)
	fmt.Println("tb:",tb)
	fmt.Println("pb:",pb)
}