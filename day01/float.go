package main

import "fmt"

//浮点型
func main(){
	//math.MaxFloat32//float32最大值
	f1:=1.23456
	fmt.Printf("%T\n",f1)//默认go语言中的小数都是float64
	f2:=float32(1.24)
	fmt.Printf("%T\n",f2)//显示声明float32类型
	
}