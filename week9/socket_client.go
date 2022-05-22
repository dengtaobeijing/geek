package main

import (
	"fmt"
	"net"

	"geek/week9/until"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Connect err:", err)
		return
	}
	defer conn.Close()
	msg := "hello socket"
	for i := 0; i < 20; i++ {
		// 调用协议编码协议
		b, err := until.Encode(msg)
		if err != nil {
			fmt.Println("Encode err:", err)
		}
		conn.Write(b)
		fmt.Println("Send msg：", b)
	}
}
