package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				// clean up activities
				return
			}
			fmt.Println("Received:", msg)
		case <- time.After(3 * time.Second):
			fmt.Println("Timeout.")
		}
	}
}

/*
----Output----
Received: 1
Channel closed
----Output----
*/