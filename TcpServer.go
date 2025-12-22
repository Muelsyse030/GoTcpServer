package main

import(
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main(){
	listener , err := net.Listen("tcp","127.0.0.1:8080")
	if err != nil{
		fmt.Println("Error starting TCP server:", err)
	}
	defer listener.Close()
	fmt.Println("Server started on 127.0.0.1:8080")
	for{
		conn , err := listener.Accept()
		if err != nil{
			fmt.Println("Error accepting connection:", err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close() 
	reader := bufio.NewReader(conn)
	message , _ := reader.ReadString('\n')
	fmt.Println("Received message:", strings.TrimSpace(message))

	newmessage := strings.ToUpper(message)
	conn.Write([]byte(newmessage + "\n"))
}