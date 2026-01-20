package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// From home version (WFH).
func main() {

	// Describe which local TCP address to listen on.
	// We listen so the server can "Connect to: <ip>:<port>\0".
	listenAddr := ":40000"

	// Creates a TCP listening socket.
	// net.Listen means "create a TCP socket and bind it".
	listenSock, err := net.Listen("tcp", listenAddr)
	if err != nil {
		panic(err)
	}
	defer listenSock.Close()

	fmt.Println("Listening for TCP connections on", listenAddr)

	// Runs the code inside this function in parallel with main(), (Concurrency).
	go func() {

		// Describe the remote TCP server address.
		serverAddr := "127.0.0.1:33546"

		// Create a TCP connection to the server.
		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			panic(err)
		}

		// Disable TCP coalescing (TCP_NODELAY).
		if tcp, ok := conn.(*net.TCPConn); ok {
			_ = tcp.SetNoDelay(true)
		}

		fmt.Println("Connected to", serverAddr)

		// Receive welcome message and echoes.
		reader := bufio.NewReader(conn)

		// Sends a few messages.
		conn.Write([]byte("First Hello\000"))
		time.Sleep(50 * time.Millisecond)
		conn.Write([]byte("Second Hello\000"))
		time.Sleep(50 * time.Millisecond)
		conn.Write([]byte("Third Hello\000"))

		// Tell server to connect back to us (insert your machine ip).
		conn.Write([]byte("Connect to: x.x.x.x:40000\000"))

		// Read messages from server (welcome + echoes).
		for {
			msg, err := reader.ReadBytes(0)
			if err != nil {
				return
			}
			fmt.Println("from server:", string(msg[:len(msg)-1]))
		}
	}()

	// Accept blocks until the server connects back.
	conn, err := listenSock.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Accepted TCP connection from", conn.RemoteAddr())

	// Receive messages on the callback connection.
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadBytes(0)
		if err != nil {
			return
		}
		fmt.Println("from callback:", string(msg[:len(msg)-1]))
	}
}
