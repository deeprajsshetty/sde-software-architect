package main

import (
	"fmt"
	"time"
)

// SYNCHRONIZING DATA EXCHANGE
func main() {
	data := make(chan string)

	go func ()  {
		for i := range 5 {
			data <- fmt.Sprintf("hello 0%d\n", i)
		}
		close(data)
	}()
	// close(data)	// Channel closed before Goroutine could send a value to the channel


	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	}	// Loops over only on active channel, creates receiver each time and stops creating reciver (looping) once the channel is closed
}

/*
-----------------------Output-------------------------
Received value: hello 00
 : 2025-09-28 00:18:26.440191 +0530 IST m=+0.000171793
Received value: hello 01
 : 2025-09-28 00:18:26.440671 +0530 IST m=+0.000651376
Received value: hello 02
 : 2025-09-28 00:18:26.440686 +0530 IST m=+0.000666501
Received value: hello 03
 : 2025-09-28 00:18:26.440689 +0530 IST m=+0.000669959
Received value: hello 04
 : 2025-09-28 00:18:26.440747 +0530 IST m=+0.000728293
 -----------------------Output-------------------------
*/