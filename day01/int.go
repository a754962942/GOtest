package main

import "fmt"

func main() {
	i1:=1024
	//十进制转二进制
	fmt.Printf("%b\n",i1)
	//十进制输出
	fmt.Printf("%d\n",i1)
	//十进制转八进制
	fmt.Printf("%o\n",i1)
	//十进制转十六进制
	fmt.Printf("%x\n",i1)
	fmt.Println("------------------------")
	i2:=077
	//八进制转二进制
	fmt.Printf("%b\n",i2)
	//八进制转十进制
	fmt.Printf("%d\n",i2)
	//八进制转十六进制
	fmt.Printf("%x\n",i2)
	//声明int8类型的变量c
	var c int8 = 9
	fmt.Printf("%T:%d\n",c,c)
}
