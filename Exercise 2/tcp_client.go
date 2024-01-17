package main

import (
	"fmt"
	"net"
)

const local_IPAdress = "10.100.23.22:33546"
const server_IPAdress = "10.100.23.129:33546"

func serverRead(conn *net.TCPConn) {
	for {
		buf2 := make([]byte, 1024)
		_, err := conn.Read(buf2)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s \n", buf2)
	}
}

func serverAccept(conn *net.TCPConn) {
	addr, err := net.ResolveTCPAddr("tcp", local_IPAdress)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	_, err = conn.Write(append([]byte("Connect to: "+local_IPAdress), 0))
	s2c_conn, err := listener.AcceptTCP()
	_, err = s2c_conn.Write(append([]byte("Accepted from group 12"), 0))
}

func serverConnect() *net.TCPConn {
	addr, err := net.ResolveTCPAddr("tcp", server_IPAdress)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	return conn
}

func main() {
	// Connect to labserver
	conn := serverConnect()
	defer conn.Close()

	go serverRead(conn)
	serverAccept(conn)

	select {}
}
