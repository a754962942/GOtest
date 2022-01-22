package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readfile() {
	filObj, err := os.Open("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\day05.md")
	if err != nil {
		fmt.Printf("打开文件失败，问题是：%v\n", err)
		return
	}

	//定义需要每次读多少
	//var s = make([]byte, 128)
	var tmp [128]byte
	for {
		n, err := filObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("读取文件失败，错误是:%v\n", err)
			return
		}

		//以字符串格式输出读取到的字符
		fmt.Println(string(tmp[:n]))
		if n <= 0 {
			return
		}

	}
}

//利用Bufio包读取文件
func readfromfilebyBufio() {
	filobj, err := os.Open("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\day05.md")
	if err != nil {
		fmt.Printf("打开文件失败，错误是：%v\n", err)
	}
	//关闭文件
	defer filobj.Close()
	//创建一个用来从文件中读取内容的对象
	reader := bufio.NewReader(filobj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("读取文件失败，错误是:%v\n", err)
			return
		}
		fmt.Print(line)

	}
}

func readFromFileByIouyil() {
	ret, err := ioutil.ReadFile("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\day05.md")
	if err != nil {
		fmt.Printf("读文件失败，错误是：%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

//读文件
func main() {
	//readFromBufio
	//readfromfilebyBufio()
	readFromFileByIouyil()
}
