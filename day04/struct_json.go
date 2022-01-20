package main

import (
	"encoding/json"
	"fmt"
)

//结构体与JSON

//1.(序列化)把Go语言中的结构体变量 ———>json格式的字符串
//2，(反序列化)json格式的字符串 ——>Go语言中能识别的结构体变量
type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "sh",
		Age:  100,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("Marshal failed,err:%v\n", err)
		return
	}
	fmt.Printf("%v\n", string(b))
	//反序列化
	str := `{"name":"abc","age":18}`
	var p2 person
	//函数传参修改的是值的副本，修改值本身传递指针
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%v\n", p2)

}
