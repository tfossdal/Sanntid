package main

import (
	"fmt"
	"net"
	"time"
)

var startTime = float64(time.Now().Unix())

func main() {
	addr, err := net.ResolveUDPAddr("udp4", "localhost:20095")
	if err != nil {
		fmt.Println("Unable to resolve")
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println("Unable to listen")
	}
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		//fmt.Println("Running")
		if float64(time.Now().Unix()) == startTime+60 {
			fmt.Println("STOPPED")
			return
		}
		buf := make([]byte, 1024)
		_, _, err = conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Died")
			return
		}
		fmt.Printf("Recieved Message: %s\n", buf[:])
	}
}
