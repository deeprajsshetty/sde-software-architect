package main

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the background
// and come back to join the main thread once the functions are finished/ready to return
// any value
// Goroutine do not stop the program flow and are non blocking
func main() {
	var err error
	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After say hello function.")

	go func() {
		err = doWork()
	}()
	// err = go doWork() // This is not accepted

	go printNumbers()
	go printLetters()
	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}
}

func sayHello()  {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers()  {
	for i := 0; i < 5; i++ {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters()  {
	for _, letter := range "abcdef" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)

	return fmt.Errorf("an error occured in doWork.")
}

/*
-----------------------------Output----------------------------
Beginning program.
After say hello function.
a 2025-09-27 21:11:10.420124 +0530 IST m=+0.000268584
Number: 0 2025-09-27 21:11:10.420169 +0530 IST m=+0.000313501
Number: 1 2025-09-27 21:11:10.521561 +0530 IST m=+0.101705709
b 2025-09-27 21:11:10.621586 +0530 IST m=+0.201731751
Number: 2 2025-09-27 21:11:10.621727 +0530 IST m=+0.201872459
Number: 3 2025-09-27 21:11:10.722827 +0530 IST m=+0.302972667
c 2025-09-27 21:11:10.822774 +0530 IST m=+0.402920167
Number: 4 2025-09-27 21:11:10.823017 +0530 IST m=+0.403163001
d 2025-09-27 21:11:11.023888 +0530 IST m=+0.604035251
e 2025-09-27 21:11:11.224346 +0530 IST m=+0.804493667
Hello from Goroutine
f 2025-09-27 21:11:11.424934 +0530 IST m=+1.005082501
Error: an error occured in doWork.
-----------------------------Output----------------------------
*/