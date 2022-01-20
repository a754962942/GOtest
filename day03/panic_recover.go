package main

import "fmt"

//panic和recover
func funcA() {
	fmt.Println("This is A")
}
func funcB() {
	defer func(){
		err:=recover()
		
		fmt.Println("In defer")
	}()
	panic("出错了")//程序崩溃推出
	fmt.Println("This is B")
	
}
func funcC() {
	fmt.Println("This is C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
