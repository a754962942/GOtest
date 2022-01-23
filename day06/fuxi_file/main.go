package main

import (
	"fmt"
	"io"
	"os"
)

//文件操作
func f1() {
	fileObj, err := os.Open("F:\\go\\src\\github.com\\a754962942\\GOtest\\day06\\fuxi_file\\main.go")
	if err != nil {
		fmt.Printf("打开文件失败，错误是：%v\n", err)
		return
	}
	defer fileObj.Close() //当Err!=nil时，fileObj是nil，空指针不能调用Close()方法
	//定义读取的字符长度及储存位置
	for {
		//
		s := make([]byte, 128)
		n, err := fileObj.Read(s)
		if err == io.EOF {
			fmt.Print("文件读取完毕\n")
			return
		}
		if err != nil {
			fmt.Printf("读取文件失败，错误是：%v\n", err)
			return
		}
		fmt.Printf("读取了%v个字符\n", n)
		fmt.Println(string(s))
		if n <= 1 {
			return
		}
	}
}
func main() {
	f1()
}
