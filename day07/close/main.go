package main

import "fmt"

//close

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	//通道关闭了还可以取值
	a := <-ch1
	b := <-ch1
	fmt.Println(a, b)
	//通道值关闭数据取完了也可以取值，取的是对应类型的零值
	x, ok := <-ch1
	fmt.Println(x, ok)
}
