package main

import "fmt"

func main() {
	ch := make(chan int)

	// === NON BLOCKING SEND OPERATION
	select {
	case ch <- 1:
		fmt.Println("Sent message.")
	default:
		fmt.Println("Channel is not ready to receive.")
	}
}

/*
-------------Output-------------
Channel is not ready to receive.
-------------Output-------------
*/