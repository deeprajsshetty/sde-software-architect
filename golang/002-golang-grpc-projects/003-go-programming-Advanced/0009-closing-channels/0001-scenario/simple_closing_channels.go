package main

import "fmt"

// Unbuffered - simple closing channels
func main() {
	ch := make(chan int)

	go func(){
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Received:", value)
	}
}

/*
---Output---
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
---Output---
*/