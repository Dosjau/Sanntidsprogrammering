package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "10.0.0.10:20000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	//status, err := bufio.NewReader(conn).ReadString('\n')
	time.Sleep(1 * time.Second)
}
