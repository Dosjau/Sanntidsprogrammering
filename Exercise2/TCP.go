package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:34933")
	if err != nil {
		fmt.Printf("Connect error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Send a message: ")
		message, _ := reader.ReadString('\n')

		// SEND 1024 bytes
		sendBuf := make([]byte, 1024)
		copy(sendBuf, []byte(message))
		conn.Write(sendBuf)

		// RECEIVE 1024 bytes
		recvBuf := make([]byte, 1024)
		_, err := conn.Read(recvBuf)
		if err != nil {
			fmt.Printf("Server error: %v\n", err)
			return
		}

		fmt.Printf("Server says: %s\n", string(recvBuf))
	}
}
