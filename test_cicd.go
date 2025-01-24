package main

import (
	"fmt"
	"net"
	"os"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	message := "hello world! \n"
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("error: can not write to connection, err:", err.Error())
		return
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./test_cicd <PORT_NUMBER>")
		os.Exit(1)
	}

	address := fmt.Sprintf(":%s", os.Args[1])

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("error: failed to create tcp server, err:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("listening on address:", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error: failed to accept connection, err:", err.Error())
			continue
		}

		go handleConn(conn)
	}
}
