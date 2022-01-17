package main

import "fmt"

//append()为切片追加元素
func main() {
	// s1 := []string{"北京", "上海", "深圳"}
	// fmt.Printf("s1:%v,len(s1):%v,cap(s1):%v\n", s1,len(s1),cap(s1))
	// s1 = append(s1, "日本")
	// fmt.Printf("s1:%v,len(s1):%v,cap(s1):%v\n",s1,len(s1),cap(s1))
	// s2:=[]string{"武汉","西安","杭州"}
	// s1=append(s1, s2...)//使用...可以将slice拆分开进行添加
	// fmt.Printf("s1:%v,len(s1):%v,cap(s1):%v\n",s1,len(s1),cap(s1))

	//关于append删除切片中的某个元素
	a1 := [...]int{1, 2, 3, 4, 5, 6, 8, 9, 9, 05, 3}
	s1 := a1[:]

	//避开索引为1的哪个2
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1)
}
