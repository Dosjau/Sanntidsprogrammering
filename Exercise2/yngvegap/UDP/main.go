package main

import (
	"fmt"
	"net"
	"time"
)

// From home version (WFH).
func main() {

	// InternetAddress addr = (0.0.0.0:20001) accept packets from any interface.
	// Describe wich local UDP address to listen to.
	recvAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 20001}

	// Creates a UDP socket and binds it to the UDP address ready to recive UDP datagrams.
	// ListenUDP means "create a UDP socket".
	recvSock, err := net.ListenUDP("udp4", recvAddr)
	//Fail safe; aborts program if errors.
	if err != nil {
		panic(err)
	}
	defer recvSock.Close()

	// Create space to store up to 1024 bytes of data received from the network.
	// UDP receive functions do not allocate memory for you.
	buffer := make([]byte, 1024)

	// Runs the code inside this function in parallel with main(), (Concurrency).
	go func() {

		// addr = (127.0.0.1:20000) since we running server and client on same machine.
		// Describes the remote UDP address to send messages to.
		// ResolveUDPAddr means "turn a human-readable address into a machine-usable UDP".
		remoteAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:20000")
		if err != nil {
			panic(err)
		}

		// Create a UDP socket for sending messages.
		// Socket used for sending so we let OS decide local port (ephemeral port), nobody needs to find the sender.
		sendSock, err := net.ListenUDP("udp4", nil)
		if err != nil {
			panic(err)
		}
		defer sendSock.Close()

		// UDP properties; no connection, no acknowledgement and no retries.
		for { // Loop
			message := []byte("Hello")
			_, _ = sendSock.WriteToUDP(message, remoteAddr) // We can ignore byte count and err since UDP is unreliable by design.
			time.Sleep(500 * time.Millisecond)
		}
	}() // Calls the function.

	for {
		// ReadFromUDP blocks until UDP packet arrives, copies data into buffer and tells how many bytes were received and who sent it.
		n, fromWho, err := recvSock.ReadFromUDP(buffer)
		// Ignore failed read attempt and continue.
		if err != nil {
			continue
		}

		// Converts received bytes into a string, prints who sent it, prints what was sent.
		fmt.Printf("from %s: %s\n", fromWho.String(), string(buffer[:n]))
	}
}
