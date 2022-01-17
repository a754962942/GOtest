package main

import "fmt"

func main() {
	var m1 map[string]int         //还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10) //估算好容量
	m1["abc"] = 18
	m1["def"] = 20
	fmt.Println(m1)
	fmt.Println(m1["abc"])

	//约定俗成用OK接受返回的布尔值
	v, ok := m1["ghi"]
	if !ok {
		fmt.Println("DON'T HAVE THIS MESSAGE")
	} else {
		fmt.Println(v)
	}
	//遍历map
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//删除
	delete(m1, "def")
	fmt.Println(m1)

}
