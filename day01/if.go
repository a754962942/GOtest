package main

import "fmt"

//if 条件判断
func main(){
	age:=18
	if age>20{
		fmt.Println("CHAIN")
	}else if age>10{
		fmt.Println("USA")
	}else {
		fmt.Println("SB")
	}
}