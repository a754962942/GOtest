package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取用户输入时有空格

func useScan() {
	var s string
	fmt.Print("请输入：")
	fmt.Scan(&s)
	fmt.Printf("您输入的是：%v\n", s)
}
func useBuf() {
	var s string
	fmt.Print("请输入：")
	reader := bufio.NewReader(os.Stdin) //获取输入端的文本
	s, _ = reader.ReadString('\n')
	fmt.Printf("您输入的是:%v\n", s)
}
func main() {
	//useScan()
	useBuf()
}
