package main

import (
	"fmt"
	"io"
	"net"
)

//tcp server端
func connect(conn net.Conn) {
	//3.与客户端通信
	var tmp [128]byte
	for {
		n, err := conn.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("输入完毕")
		}
		if err != nil {
			fmt.Println("read from conn failed,err", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}
func main() {
	//1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server failed ,err:", err)
		return
	}
	//2.等待建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept failed,err ", err)
			return
		}
		go connect(conn)
	}
}
