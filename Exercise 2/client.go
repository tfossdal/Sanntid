package main

import (
	"fmt"
	"net"
)

func main() {
	listen_addr, err := net.ResolveUDPAddr("udp4", "10.100.23.129:20012")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp4", nil, listen_addr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hei på deg"))
	if err != nil {
		fmt.Println("Error writing to UDP:", err)
		return
	}

}
