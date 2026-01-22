package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr := "10.22.227.63:33546"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Koblet til", serverAddr)

	reader := bufio.NewReader(conn)

	// Goroutine for å motta meldinger
	go func() {
		for {
			msg, err := reader.ReadString(0) // les til '\0'
			if err != nil {
				fmt.Println("Connection closed")
				os.Exit(0)
			}
			fmt.Println("Server:", msg[:len(msg)-1]) // fjern '\0'
		}
	}()

	// Goroutine for å lytte etter innkommende tilkoblinger
	go func() {
		listenPort := ":40000"

		ln, err := net.Listen("tcp", listenPort)

		if err != nil {
			fmt.Println("Listen error:", err)
			return
		}
		defer ln.Close()

		fmt.Println("Lytter på port", listenPort)

		for {
			conn, err := ln.Accept() //Venter og aksepterer innkommende tilkoblinger
			if err != nil {
				fmt.Println("Accept error:", err)
				continue
			}
			conn.Close()
		}
	}()

	// Send meldinger fra terminal
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		if text == "exit" {
			break
		}
		conn.Write(append([]byte(text), 0)) // legg til '\0'
	}
}
