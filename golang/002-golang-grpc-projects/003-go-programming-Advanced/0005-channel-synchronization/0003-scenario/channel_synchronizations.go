package main

import (
	"fmt"
)

// SYNCHRONIZING MULTIPLE GOROUTINES AND ENSURING THAT ALL GOROUTINES ARE COMPLETE
func main() {
	numGoroutines := 3
	done := make(chan int, numGoroutines)

	for i := range numGoroutines {
		//time.Sleep(1 * time.Second)
		go func (id int)  {
			fmt.Printf("Goroutine %d working...\n", id)
			done <- id	// SENDING SIGNAL OF COMPLETION
		}(i)
	}
	
	for range numGoroutines {
		<-done		// wait for each goroutine to finish, WAIT FOR ALL GOROUTINES TO SIGNAL COMPLETES
	}

	fmt.Println("All goroutines are complete")
}

/*
--------Output---------
Goroutine 2 working...
Goroutine 0 working...
Goroutine 1 working...
All goroutines are complete
--------Output---------
*/