package main

import "fmt"

func main() {
	// var name string
	// name = "abc"
	// fmt.Println(name)
	// var ages [30]int //声明一个变量ages，他是[30]int类型
	// //初始化方法一
	// ages = [30]int{1, 2, 3, 4, 5, 6}
	// fmt.Println(ages)
	// //初始化方法二
	// var ages1 = [...]int{11, 2, 3, 4, 5, 6, 7}
	// fmt.Println(ages1)
	// //初始化方法三
	// var ages3 = [...]int{1: 100, 99: 200}
	// fmt.Println(ages3)

	// var a1 [3][2]int
	// a1 = [...][2]int{
	// 	[2]int{1, 2},
	// 	[2]int{3, 4},
	// 	[2]int{5, 6},
	// }
	// fmt.Println(a1)
	//多维数组只有最外层可以使用

	//数组是值类型

	// x := [3]int{1, 2, 3}
	// y := x
	// y[1] = 200
	// fmt.Println(y)

	// f1(x)
	// fmt.Println(x)
	// f2(&x)
	// fmt.Println(x)

	//切片
	// var s1 []int //不初始化没有分配内存， ==nil
	// //初始化方法1
	// s1 =[]int{1,2,3}
	// fmt.Println(s1)
	// //make 初始化 make 分配内存
	// s2:=make([]int,5,10)
	// fmt.Println(s2)
	// s3:=make([]int,0,2) //使用make给s3分配了内存，不管s3内是否长度为0
	// fmt.Println(s3==nil)
	// arr := [5]int{1, 2, 3, 4, 5}
	// sli := arr[:4] //声明一个sli切片，实质是一个指向arr的指针
	// fmt.Println(sli)
	// sli[0] = 100 //对sli【0】进行修改，实质是修改arr【0】
	// fmt.Println(arr)

	// addr := "abc"
	// addrp := &addr      //将内存地址赋值给addrp
	// fmt.Println(addrp)  //输出addr的地址
	// fmt.Println(*addrp) //输出addr地址对应的值

	//map
	m1 := make(map[string]int) //使用前一定要用make初始化，否则为nil
	m1["abc"] = 123

	fmt.Println(m1)
	fmt.Println(m1["bbb"]) //如果key不存在，返回的是value类型对应的零值
	v, ok := m1["c"]
	if !ok {
		fmt.Println("没有C")
	} else {
		fmt.Println("c的值为：", v)
	}
	delete(m1, "d") //删除的key不存在，什么都不做
	delete(m1, "abc")
	fmt.Println(m1)
}

// func f1(a [3]int) {
// 	//Go语言中传递的都是值
// 	a[1] = 100 //此处修改的是副本的值
// }
// func f2(a *[3]int) {
// 	a[1] = 100
// }
