package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	port            = ":9999"
	heartbeatPeriod = 500 * time.Millisecond
	timeoutPeriod   = 2 * time.Second
	stateFile       = "state.txt"
	spawnCooldown   = 5 * time.Second
)

// -------------------- State --------------------

func readState() int {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return 0
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0
	}
	return n
}

func writeState(n int) {
	_ = os.WriteFile(stateFile, []byte(strconv.Itoa(n)), 0644)
}

// -------------------- Backup --------------------

func runBackup() {
	fmt.Println("--- Backup phase ---")

	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		// Port already taken → primary exists
		fmt.Println("... primary detected, staying passive")
		return
	}
	defer conn.Close()

	buf := make([]byte, 16)
	lastHeartbeat := time.Now()

	for {
		_ = conn.SetReadDeadline(time.Now().Add(heartbeatPeriod))
		_, _, err := conn.ReadFromUDP(buf)
		if err == nil {
			lastHeartbeat = time.Now()
		}

		if time.Since(lastHeartbeat) > timeoutPeriod {
			fmt.Println("... timed out")
			return
		}
	}
}

// -------------------- Primary --------------------

func spawnBackup() {
	fmt.Println("... creating new backup")
	cmd := exec.Command(
		"cmd",
		"/C",
		"start",
		"",
		".\\ProcessPair.exe",
	)
	_ = cmd.Start()
}

func backupExists() bool {
	addr, _ := net.ResolveUDPAddr("udp", port)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return true // Port in use → backup exists
	}
	conn.Close()
	return false
}

func runPrimary() {
	fmt.Println("--- Primary phase ---")

	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1"+port)
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()

	lastSpawn := time.Time{}

	spawnBackup()
	lastSpawn = time.Now()

	counter := readState()
	fmt.Println("... resuming from", counter)

	for {
		counter++
		fmt.Println(counter)
		writeState(counter)

		_, _ = conn.Write([]byte("alive"))

		if !backupExists() && time.Since(lastSpawn) > spawnCooldown {
			fmt.Println("... backup missing, recreating")
			spawnBackup()
			lastSpawn = time.Now()
		}

		time.Sleep(1 * time.Second)
	}
}

// -------------------- Main --------------------

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--reset" { // Type: ".\ProcessPair.exe --reset" to resest count
		fmt.Println("... reset requested, clearing state")
		os.Remove(stateFile)
	}

	runBackup()
	runPrimary()
}
