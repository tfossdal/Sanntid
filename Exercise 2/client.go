package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a UDP address to bind to
	addr, err := net.ResolveUDPAddr("udp", ":30000")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	fmt.Printf("Adress: %s\n", addr)

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		return
	}
	defer conn.Close()
}
