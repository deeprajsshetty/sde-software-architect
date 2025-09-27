package main

import "fmt"

func main() {
	ch := make(chan int)

	go func ()  {
		ch <- 9	// Blocking until the value is received
		fmt.Println("Sent value")
	}()

	value := <-ch	// Blocking until the value is sent
	fmt.Println(value)
}

/*
---Output---
Sent value
9
---Output---
*/