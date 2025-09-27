package main

import (
	"fmt"
	"time"
)

// BLOCKING ON SEND ONLY IF THE BUFFER IS FULL
func main() {
	// make(chan Type, capacity)
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	go func ()  {
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
	}()
	fmt.Println("Blocking starts")
	ch <- 3			// Blocks because the buffer is full
	fmt.Println("Blocking ends")
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	
	fmt.Println("Buffered channels")
}

/*
-----Output-----
Blocking starts
Blocking ends
Received: 2
Received: 3
Buffered channels
-----Output-----
*/