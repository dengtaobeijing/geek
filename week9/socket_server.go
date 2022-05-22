package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"geek/week9/until"
)

func main() {
	// 本地端口启动服务
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Start server fail :", err)
		return
	}
	fmt.Println("Start server...")
	// 等待连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connect fail ", err)
			break
		}
		fmt.Println("Connect ok ")
		go Process(conn)
	}
	// 通信
}

func Process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := until.Decode(reader)

		fmt.Println("Receive msg：", msg)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Decode fail，err:", err)
		}
	}
}
