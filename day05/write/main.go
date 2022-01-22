package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

func read1() {
	filObj, err := os.Open("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\day05.md")
	if err != nil {
		fmt.Printf("打开文件失败,错误是：%v\n", err)
		return
	}
	//定义读取的字符存储位置
	var arr [128]byte
	for {
		n, err := filObj.Read(arr[:])
		if err == io.EOF {
			fmt.Print("文件读取结束")
			return
		}
		if err != nil {
			fmt.Printf("读取文件失败,错误是：%v\n", err)
			return
		}
		//读取文件
		fmt.Print(string(arr[:n]))
		if n < 128 {
			return
		}
	}
}

//利用Bufio读取文件
func read2() {

	filObj, err := os.Open("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\day05.md")

	if err != nil {
		fmt.Printf("打开文件失败，错误是:%v\n", err)
		return
	}
	reader := bufio.NewReader(filObj)
	for {
		//按行读取文件
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print("文件读取结束")
			return

		}
		if err != nil {
			fmt.Printf("读取文件失败，错误是:%v\n", err)
			return
		}
		fmt.Print(str)

	}
}

//利用IOutil读取文件
func read3() {

	s, err := ioutil.ReadFile("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\day05.md")
	if err != nil {
		fmt.Printf("读取文件失败,\n", err)
		return
	}
	fmt.Print(string(s))
}

//打开文件写内容
//os.openfile
func write1() {
	filObj, err := os.OpenFile("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\write\\demo1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败，问题是：%v\n", err)
		return
	}
	//write
	filObj.Write([]byte("heixiaoyubumingbai\n"))
	//writestring
	filObj.WriteString("黑小雨理解不了\n")
	defer filObj.Close()
}

//bufio
func write2() {
	filObj, err := os.OpenFile("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\write\\demo1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败，err:\n", err)
		return
	}
	wr := bufio.NewWriter(filObj)
	//写入字符串
	wr.WriteString("hello 灵宝\n") //将数据先写入缓存
	wr.Flush()                   //将缓存中的内容写入文件
	defer filObj.Close()
}

//ioutl
func write3() {
	str := "helloworld"
	err := ioutil.WriteFile("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\write\\demo1.txt", []byte(str), fs.FileMode(os.O_APPEND))
	if err != nil {
		fmt.Printf("打开文件错误%v\n", err)
		return
	}
}
func main() {
	//write1()
	//write2()
	write3()
	//read1()
	//read2()
	//read3()
}
