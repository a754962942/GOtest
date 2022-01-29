package main

import (
	"fmt"
	"os"
)

//1.文件对象的类型
//2. 获取对象文件的信息
func main() {
	fileObj, err := os.Open("github.com\\a754962942\\Gotest\\day06\\file\\main.go")
	if err != nil {
		fmt.Printf("open file failed ,err:%v\n", err)
		return
	}
	//1.文件对象的类型
	fmt.Printf("%T\n", fileObj)
	//2. 获取对象文件的信息
	FileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file Info failed,err:%v \n", err)
		return
	}
	fmt.Printf("文件大小是：%d B\n", FileInfo.Size())
}
