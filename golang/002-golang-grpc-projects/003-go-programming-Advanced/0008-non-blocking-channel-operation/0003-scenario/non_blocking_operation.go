package main

import (
	"fmt"
	"time"
)

// === NON BLOCKING OPERATION IN REAL TIME SYSTEMS
func main() {
	data := make(chan int)
	quit := make(chan bool)

	// Consumer
	go func ()  {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received:", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}	
	}()
	
	// Producer
	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	// quit single
	quit <- true
}

/*
------Output------
Data received: 0
Waiting for data...
Waiting for data...
Data received: 1
Waiting for data...
Waiting for data...
Data received: 2
Waiting for data...
Waiting for data...
Data received: 3
Waiting for data...
Waiting for data...
Data received: 4
Waiting for data...
Waiting for data...
Stopping...
------Output------
*/