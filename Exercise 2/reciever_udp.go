package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a UDP address to bind to
	addr, err := net.ResolveUDPAddr("udp", ":20012")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		return
	}

	for {

		//IP Adress: 10.100.23.129

		buf := make([]byte, 1024)

		conn.ReadFromUDP(buf)

		fmt.Printf("IP Adress: %s\n", buf[:])
	}
}
