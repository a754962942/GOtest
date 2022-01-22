package main

import (
	"encoding/json"
	"fmt"
)

//序列化和反序列化
func main() {
	type point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
	p1 := point{100, 200}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("Marshal failed, err:%v\n", err)
	}
	fmt.Println(string(b))

	//反序列化：由字符串-->Go中的结构体变量
	str1 := `{"x":100,"y":123}`
	str2 := point{} //造一个结构体变量，准备接受反序列化的值
	err = json.Unmarshal([]byte(str1), &str2)
	if err != nil {
		fmt.Printf("Ummarshal failed,err:%v\n", err)
	}
	fmt.Println(str2)
}
