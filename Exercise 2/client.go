package main

import (
	"fmt"
	"net"
)

// func getLabAdress() {
// 	// Create a UDP address to bind to
// 	addr, err := net.ResolveUDPAddr("udp", ":30000")
// 	if err != nil {
// 		fmt.Println("Error resolving UDP address:", err)
// 		return
// 	}

// 	fmt.Printf("Adress: %s\n", addr)

// 	// Create a UDP connection
// 	conn, err := net.ListenUDP("udp", addr)
// 	if err != nil {
// 		fmt.Println("Error listening on UDP:", err)
// 		return
// 	}

// 	//IP Adress: 10.100.23.129

// 	buf := make([]byte, 1024)

// 	conn.ReadFromUDP(buf)

// 	fmt.Printf("IP Adress: %s\n", buf[:])
// }

// func main() {
// 	send_addr, err := net.ResolveUDPAddr("udp", ":20021")
// 	if err != nil {
// 		fmt.Println("Error resolving UDP address:", err)
// 		return
// 	}

// 	conn_echo, err := net.ListenUDP("udp", send_addr)

// 	if err != nil {
// 		fmt.Println("Error listening on UDP:", err)
// 		return
// 	}
// 	_, err1 := conn_echo.WriteTo([]byte("Hei på deg"), send_addr)
// 	if err1 != nil {
// 		fmt.Println("Error writing to UDP:", err1)
// 		return
// 	}

// 	buf_echo := make([]byte, 1024)

// 	conn_echo.ReadFromUDP(buf_echo)

// 	fmt.Printf("Echo: %s\n", buf_echo[:])
// }

func main() {
	listenAddr, err := net.ResolveUDPAddr("udp4", ":30000")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp4", listenAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	addr, err := net.ResolveUDPAddr("udp4", "10.100.23.129:20021")
	if err != nil {
		panic(err)
	}

	_, err = conn.WriteTo([]byte("Halla"), addr)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)

	conn.ReadFromUDP(buf)

	fmt.Printf("Echo: %s\n", buf[:])

}
