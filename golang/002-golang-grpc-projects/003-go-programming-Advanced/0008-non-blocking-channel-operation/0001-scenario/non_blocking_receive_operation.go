package main

import "fmt"

func main() {
	ch := make(chan int)

	// === NON BLOCKING RECEIVE OPERATION
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No messages available.")
	}
}

/*
-------Output---------
No messages available.
-------Output---------
*/