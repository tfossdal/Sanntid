package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		// TODO: push real value to buffer
		ch <- i
	}

}

func consumer(ch chan int) {

	time.Sleep(1 * time.Second)
	for {
		select {
		case i := <-ch: //TODO: get real value from buffer
			fmt.Printf("[consumer]: %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func main() {

	// TODO: make a bounded buffer
	bufferChan := make(chan int, 4)

	go consumer(bufferChan)
	go producer(bufferChan)

	select {}
}
