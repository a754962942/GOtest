package main

import "fmt"

func main() {
	// switchAge(15)
	// switchCase(5)
	// switchCase(8)
	// switchCase(10)
	// switchDemo()
	//gotoDemo()
	gotoDemo1()
}

//当分支使用表达式时，switch语句后就不需要跟判断变量，例如
// func switchAge(n int){
// 	age:=n
// 	switch {
// 	case age<25:
// 		fmt.Println("好好学习")
// 	case age>25&&age<35:
// 		fmt.Println("工作")
// 	case age>60:
// 		fmt.Println("享受")
// 	default:
// 		fmt.Println("好")
// 	}
// }
// //如果switch后有判断变量，那么分支就不用使用表达式
// func switchCase(n int){
// 	switch n {
// 	case 1,3,5,7,9:
// 		fmt.Println("奇数")
// 	case 2,4,6,8:
// 		fmt.Println("偶数")
// 	default:
// 		fmt.Println("ERROR")
// 	}
// }
// //fallthrough语法可以执行满足条件的下一个case，是为了兼容C语言中的case设计的
// func switchDemo(){
// 	s:="a"
// 	switch  {
// 	case s == "a":
// 		fmt.Println("a")
// 		fallthrough
// 	case s == "b":
// 		fmt.Println("b")
// 	case s =="c":
// 		fmt.Println("c")
// 	}
// }
func gotoDemo() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		if breakFlag {
			break
		}
	}
}

//goto +label实现跳出多层for循环
func gotoDemo1() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 4 {
				goto Label //跳到Label标签
			}
			fmt.Printf("%d - %d \n", i, j)
		}
	}
Label: //Label标签
	fmt.Println("This is Label")
}
