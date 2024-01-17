package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to labserver
	conn, err := net.Dial("tcp", "10.100.23.129:20012")
	if err != nil {
		panic(err)
	}

	buf1 := make([]byte, 1024)
	_, err = conn.Read(buf1)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write([]byte("Hello from the other side'\\0"))

	buf2 := make([]byte, 1024)
	_, err = conn.Read(buf2)
	if err != nil {
		fmt.Println(err)
		return
	}

}
