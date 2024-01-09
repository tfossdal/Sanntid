// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0
var p = 5

func incrementing(message, status chan int) {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000001; j++ {
		message <- 1
	}
	status <- 1
}

func decrementing(message, status chan int) {
	//TODO: decrement i 1000000 times
	for k := 0; k < 1000000; k++ {
		message <- 2
	}
	status <- 1
}

func server(message, value chan int) {
	for {
		select {
		case cmd := <-message:
			switch cmd {
			case 1: //increment
				i++
			case 2: //decrement
				i--
			case 3: //read
				value <- i
			}

		}
	}
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	// Channels
	value_commands := make(chan int)
	status_channel := make(chan int)
	value_channel := make(chan int)

	// TODO: Spawn both functions as goroutines
	go incrementing(value_commands, status_channel)
	go decrementing(value_commands, status_channel)
	go server(value_commands, value_channel)

	<-status_channel
	<-status_channel
	value_commands <- 3
	p := <-value_channel
	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", p)
}
