package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp client
func main() {
	//1.与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("connect 127.0.0.1:20000 failed,err:", err)
		return
	}
	//2.发送数据
	for {
		fmt.Print("请输入：")
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		if str == "exit" {
			break
		}
		conn.Write([]byte(str))
	}
	conn.Close()
}
