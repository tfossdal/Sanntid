package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

func sendMessage(s string, conn *net.UDPConn) {
	_, err := conn.Write([]byte(s))
	if err != nil {
		fmt.Println("Failed to send")
	} else {
		fmt.Println("test")
	}
}

func main() {
	cmd := exec.Command("gnome-terminal", "--", "go", "run", "A.go")
	cmd.Run()

	addr, err := net.ResolveUDPAddr("udp4", "localhost:20095")
	if err != nil {
		fmt.Println("Failed to resolve")
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Println("Failed to dial")
	}
	defer conn.Close()
	time.Sleep(1 * time.Second)
	for {
		sendMessage("hei hei", conn)
		time.Sleep(2 * time.Second)
	}
}
