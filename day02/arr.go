package main

import (
	"fmt"
)

//Array
//存放元素的容器
//必须指定存放的元素的类型和容量（长度）
//数组的长度是数组类型的一部分
func main() {
	// var a1 [3]bool //[true false true]
	// var a2 [4]bool //[true false false false]

	// fmt.Printf("A1:%T\tA2:%T\n",a1,a2)

	// //数组的初始化
	// //如果不初始化，默认元素都是零值（布尔值：false，整数、浮点数：0，字符串："")
	// fmt.Println(a1,a2)
	// //方式1:初始化数字组时可以使用初始化列表来设置数组元素的值
	// a1 = [3]bool{true,true,true}
	// //方式2:根据初始值自动推断数组长度
	// a3 := [...]int {0,1,2,3,4,5,6,7,8}
	// fmt.Println(len(a3))
	// //方法3：根据索引来初始化
	// a4:=[5]int{0:1,4:2}
	// fmt.Println(a4)
	// //数组遍历
	// //方法1：按照index取
	// citys:=[...]string{"北京","上海","广州","成都"}
	// for i := 0; i < len(citys); i++ {
	// 	fmt.Println(citys[i])
	// }
	// //2.for range遍历
	// for i,v :=range citys{
	// 	fmt.Printf("%d:%s\n",i,v)
	// }

	//a[3][2]int
	//[[1,2][3,4][5,6]]

	// a := [...][2]int{
	// 	[2]int{1,2},
	// 	[2]int{3,4},
	// 	[2]int{5,6},
	// }
	// //fmt.Println(a)
	// //二维数组的遍历
	// for _,v:=range a{
	// 	fmt.Printf("%v\n",v)
	// }

	// //数组是值类型
	// //值类型：
	// b1 :=[3]int{1,2,3} //{1,2,3}
	// b2 :=b1 //{1,2,3}
	// b2[0] = 100//{100,2,3}
	// fmt.Println(b1,b2)//{1,2,3} {100,2,3}
	//作业1
	// a:=[...]int{
	// 	1,3,5,7,8,
	// }
	// sum:=0
	// for _,v:=range a{
	// 	sum+=v
	// }
	// fmt.Println(sum)
	//作业2
	a := [...]int{
		1, 3, 5, 7, 8,
	}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%d,%d)", i, j)
			}
		}

	}
}
