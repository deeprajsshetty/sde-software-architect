package main

import "fmt"

// Receiving from the closed channel.
func main() {
	ch := make(chan int)
	close(ch)

	value, ok := <-ch
	if !ok {
		fmt.Println("Channel is closed.")
		return
	}
	fmt.Println("Received value:", value)
}

/*
-----Output------
Channel is closed.
-----Output------
*/