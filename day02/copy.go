package main

import "fmt"

func main() {
	a1:=[]int{1,3,5}
	a2:=a1
	a3:=make([]int, 3,3)
	copy(a3,a1)
	fmt.Println(a1,a2,a3)
	a2[0]=100
	fmt.Println(a1,a2,a3)
	//将a1中的索引为1的元素删除
	a1=append(a1[:1],a1[2:]...) 
	fmt.Println(a1)
}
