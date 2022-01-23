package main

import (
	"fmt"
	"os"
)

//在文件中插入内容
/*

func (*File) Seek
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误。
*/

func f() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("打开文件失败，错误是:%v\n", err)
		return
	}
	defer fileObj.Close()
	fileObj.Seek(3, 0) //光标移动
	var ret [1]byte
	n, err := fileObj.Read(ret[:]) //读到了ret的切片里
	if err != nil {
		fmt.Printf("读取文件失败，错误是：%v\n", err)
		return
	}
	fmt.Println(string(ret[:n])) //return的结果存在了slice里

}
func main() {
	f()
}
