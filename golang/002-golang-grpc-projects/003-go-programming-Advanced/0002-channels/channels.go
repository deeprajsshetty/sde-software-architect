package main

import (
	"fmt"
	"time"
)

func main() {

	// variable := make(chan type) '<-' operator
	greeting := make(chan string) 
	greetString := "Hello"

	go func ()  {
		greeting <- greetString // blocking because it is continuously trying to receive values, it is ready to receive continuous flow of data.
		greeting <- "World"
		for _, e := range "abcdef" {
			greeting <- fmt.Sprintf("Alphabet: %c", e)
		}
	}()

	// go func ()  {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)
	for range 6 {
		receiver = <-greeting
		fmt.Println(receiver)
	}

	time.Sleep(1 * time.Millisecond)
	fmt.Println("End of program.")
}