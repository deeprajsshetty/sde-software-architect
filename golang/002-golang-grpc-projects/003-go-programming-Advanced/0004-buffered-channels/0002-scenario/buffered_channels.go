package main

import (
	"fmt"
	"time"
)

// BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
func main() {
	ch := make(chan int, 2)
	go func ()  {
		time.Sleep(2 * time.Second)
		ch <- 1	
	}()
	fmt.Println("Value:", <-ch)
	fmt.Println("End of program.")
}

/*
-----Output-----
Value: 1
End of program.
-----Output-----
*/