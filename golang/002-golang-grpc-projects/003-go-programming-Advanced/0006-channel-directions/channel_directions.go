package main

import "fmt"

func main() {
	ch := make(chan int)

	producer(ch)
	consumer(ch)
}

// Send only channel
func producer(ch chan <- int)  {
	go func ()  {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only channel
func consumer(ch <- chan int)  {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}

/*
----Output----
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
----Output----
*/