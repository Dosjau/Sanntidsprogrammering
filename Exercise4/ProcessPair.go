package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func check(err error) {
	//Error handling funksjon
	if err != nil {
		panic(err)
	}
}

func create_copy() {
	// Finn egen executable
	exe, err := os.Executable()
	check(err)

	// Start en ny kopi av oss selv
	cmd := exec.Command(exe)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	check(err)
}

func listen_UDP() int {
	// Sett opp UDP multicast lytter
	addr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:9999")
	check(err)

	conn, err := net.ListenUDP("udp4", addr)
	check(err)

	buf := make([]byte, 1024)
	last_known_value := 0

	for {
		//Leser etter UDP meldinger med 4 sekunders timeout og returnerer siste kjente verdi hvis timeout
		_ = conn.SetReadDeadline(time.Now().Add(4 * time.Second))

		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				fmt.Printf("PID %d: Ingen melding mottatt på 4 sekunder. Overtar som master\n", os.Getpid())
				_ = conn.Close()
				create_copy() //Lager en process pair, ny lytter som skal overvåke den nye masteren
				return last_known_value
			}
			panic(err)
		}

		last_known_value, err = strconv.Atoi(string(buf[:n]))
		if err != nil {
			continue
		}
	}
}

func main() {
	//Starter å lytte for å sjekke om det finnes en master
	fmt.Printf("PID %d: Prosess er startet\n", os.Getpid())
	value := listen_UDP()

	//Setter opp UDP sender som master
	addr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:9999")
	check(err)

	conn, err := net.DialUDP("udp4", nil, addr)
	check(err)

	amount_of_prints := 0
	// Tell opp hvert sekund
	for amount_of_prints < 30 {
		fmt.Printf("PID %d: %d\n", os.Getpid(), value)   //Printer verdi til terminal
		_, err = conn.Write([]byte(strconv.Itoa(value))) //Skriver verdi til UDP
		check(err)

		value++
		amount_of_prints++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("PID %d: Avslutter som master\n", os.Getpid())
	_ = conn.Close()
	os.Exit(0)
}
