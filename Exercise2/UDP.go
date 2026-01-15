package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	go func() {

		addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:20000")
		if err != nil {
			log.Fatal("Couldn’t resolve address:", err)
		}

		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			log.Fatal("Connection failed:", err)
		}
		defer conn.Close()

		// Send message
		for {
			message := []byte("Hello, UDP!")
			_, err = conn.Write(message)
			if err != nil {
				log.Printf("Send failed: %v", err)
				return
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	buffer := make([]byte, 1024)
	for {
		n, _, err := recvConn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Receive error: %v", err)
			return
		}
		fmt.Printf("Server says: %s\n", string(buffer[:n]))
	}
	//time.Sleep(100 * time.Microsecond)
}
