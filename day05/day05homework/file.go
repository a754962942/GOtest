package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//读写文件
//read1
func read1() {
	filObj, err := os.Open("F:\\go\\src\\github.com\\a754962942\\GOtest\\README.md")
	if err != nil {
		fmt.Printf("打开文件时发生错误:%v\n", err)
		return
	}
	//定义读取的数量及存放位置
	defer filObj.Close()
	var arr [128]byte
	for {
		n, err := filObj.Read(arr[:])
		if err == io.EOF {
			fmt.Print("读完了")
			return
		}
		if err != nil {
			fmt.Printf("读取文件时发生错误:%v\n", err)
			return
		}
		fmt.Printf("读取了%v个字符\n", n)
		fmt.Println(string(arr[:n]))

	}
}

//利用ioutil
func read2() {
	s, err := ioutil.ReadFile("F:\\go\\src\\github.com\\a754962942\\GOtest\\README.md")

	if err != nil {
		fmt.Printf("打开文件时发生错误:%v\n", err)
		return
	}
	fmt.Print(string(s))
}

//利用Bufio
func read3() {
	filObj, err := os.Open("F:\\go\\src\\github.com\\a754962942\\GOtest\\README.md")
	if err != nil {
		fmt.Printf("打开文件时发生错误:%v\n", err)
		return
	}
	reader := bufio.NewReader(filObj)
	defer filObj.Close()
	//控制读取结束的字符
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print("读完了")
			return
		}
		if err != nil {
			fmt.Printf("读取文件时发生错误:%v\n", err)
			return
		}
		fmt.Println(s)

	}
}

//写文件
//write1
func write1() {
	filObj, err := os.OpenFile("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\day05homework\\test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("打开文件时发生错误:%v\n", err)
		return
	}
	//写入文件
	filObj.WriteString("sjjsjsjjsjjss\n")
	//关闭
	defer filObj.Close()
}

//使用ioutil写入文件
func write2() {
	str := "123hello"
	err := ioutil.WriteFile("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\day05homework\\test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("打开文件时发生错误:%v\n", err)
		return
	}
}

//利用bufio写入
func write3() {
	filObj, err := os.OpenFile("F:\\go\\src\\github.com\\a754962942\\GOtest\\day05\\day05homework\\test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("打开文件时发生错误:%v\n", err)
		return
	}
	wr := bufio.NewWriter(filObj)
	wr.WriteString("456_world") //将字符串写入缓冲区
	wr.Flush()                  //将字符串从缓冲区写入文件
	defer filObj.Close()
}

//使用Scan只能捕获到单个字符结束，而不能捕获至输入结束
func useScan() {
	var s string
	fmt.Print("请输入文字：")
	fmt.Scan(&s)
	fmt.Printf("您输入的是：%v\n", s)
}
func useBufio() {
	var s string
	fmt.Print("请输入文字：")
	reader := bufio.NewReader(os.Stdin) //从终端中读取内容
	s, _ = reader.ReadString('\n')      //按照换行符结束采集
	fmt.Printf("您输入的是:%v\n", s)
}
func main() {
	//read1()
	// read2()
	//read3()
	// write1()
	//write2()
	// write3()
	// useScan()
	useBufio()
}
