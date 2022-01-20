package main

import "fmt"

//defer执行顺序
//Go语言中函数的return不是原子操作，在底层分两步执行
//第一步：返回值赋值
//defer
//第二步：真正的return返回
//函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1() int {//返回值未命名，相当于第一步在底层开辟了一个单独的空间存放返回值=5
	x := 5
	defer func() {
		x++ //修改的是X 不是返回值
	}()
	return x // 第一步 返回值= 5 第二步 defer x=6 第三步 RET=5
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //第一步 返回值=x=5    defer（修改的是返回值指向的x）：  x++=6   RET x=6
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值 = y = x =5 defer x++ RET = y = 5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //函数传参修改更改的是参数的拷贝，不影响原始值 RET = 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
