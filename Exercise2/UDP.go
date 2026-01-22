package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:20000")
	if err != nil {
		log.Fatal("Couldn't resolve address:", err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer conn.Close()

	// Send messages
	go func() {
		for {
			message := []byte("Hello, UDP!")
			_, err := conn.Write(message)
			if err != nil {
				log.Printf("Send failed: %v", err)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Receive messages
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Receive error: %v", err)
			return
		}
		fmt.Printf("Server says: %s\n", string(buffer[:n]))
	}
}
