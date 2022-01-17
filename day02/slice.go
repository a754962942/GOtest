package main

import "fmt"
//切片Slice
func main() {
	//切片的定义
	// var s1[]int//定义一个存放int类型元素的切片
	// var s2[]string//定义一个存放string类型元素的切片
	// fmt.Println(s1,s2)
	// fmt.Println(s1 == nil)
	// //初始化
	// s1 = []int {1,2,3}
	// s2 = []string{"a","b","c"}
	// fmt.Println(s1,s2)
	// //make()函数声明slice
	// a:=make([]int ,0,0)
	// fmt.Printf("a = %v,len(a) = %v,cap(a) = %v\n",a,len(a),cap(a))
	// fmt.Println(a==nil)
	// a1:=make([]int ,1,5)
	// fmt.Printf("a1 = %v,len(a1) = %v,cap(a1) = %v\n",a1,len(a1),cap(a1))
	
	//切片的赋值
	// s:=[]int{1,3,5}
	// s1:=s //s和s1都指向同一个底层数组
	// fmt.Println(s,s1)
	// s[0] = 1000
	// fmt.Println(s,s1)

	//切片的遍历
	s:=[]int{1,2,3}
	//1.索引遍历
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	//for range遍历
	for i,v:=range s{
		fmt.Printf("index:%v\tvalue:%v\n",i,v)
	}


}
