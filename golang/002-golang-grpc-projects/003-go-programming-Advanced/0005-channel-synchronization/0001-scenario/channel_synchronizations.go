package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func ()  {
		fmt.Println("Working...")
		time.Sleep(3 * time.Second)		// Do some process
		done <- struct{}{}	
	}()

	<- done
	fmt.Println("Finished.")
}

/*
---Output---
Working...
Finished.
---Output---
*/