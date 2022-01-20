package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	a := 1
	b := 2
	//defer调用语句内部值为函数，则先运算内部函数。剩余一层函数在最后执行
	//函数调用顺序进行传参，此时a=1
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
//1.calc("1",a,calc("10",1,2))
//2.defer calc("1",1,3)//"10",1,2,3
//3.a=0
//4.calc("2",0,clac("20",0,2)
//5.defer calc("2",0,2)//"20",0,2,2
//6.b=1
//7.defer calc("2",0,2)//"2",0,2,2
//8.defer calc("1",1,3)//"1",1,3,4

//"10",1,2,3
//"20",0,2,2
//"2",0,2,2
//"1",1,3,4