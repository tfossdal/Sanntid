package main

import (
	"fmt"
	"net"
)

func main() {
	send_addr, err := net.ResolveUDPAddr("udp", ":20012")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn_echo, err := net.ListenUDP("udp", send_addr)

	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		return
	}
	_, err = conn_echo.WriteTo([]byte("Hei på deg"), send_addr)
	if err != nil {
		fmt.Println("Error writing to UDP:", err)
		return
	}

	buf_echo := make([]byte, 1024)

	conn_echo.ReadFromUDP(buf_echo)

	fmt.Printf("Echo: %s\n", buf_echo[:])
}
