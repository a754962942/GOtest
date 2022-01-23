package main

import (
	"fmt"
	"time"
)

//time包
func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Second())
	//时间戳
	fmt.Printf("时间戳%v\n", now.Unix())
	fmt.Printf("纳秒时间戳：%v\n", now.UnixNano())
}
