package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. 连接服务器 (Connect)
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()

	// 2. 从标准输入读取要发送的话
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')

	// 3. 发送给服务器
	fmt.Fprintf(conn, text+"\n")

	// 4. 接收服务器的回应
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)
}