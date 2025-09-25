package main

import "fmt"

func main() {

	// panic(interface{})

	// Example of a valid input
	process(10)

	// Example of an invalid input
	process(-3)

}

func process(input int)  {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		// fmt.Println("After Panic")
		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}

/*
------------------------------------------------Output------------------------------------------
Processing input: 10
Deferred 2
Deferred 1
Before Panic
Deferred 2
Deferred 1
panic: input must be a non-negative number

goroutine 1 [running]:
main.process(0x1010242f8?)
	/Users/deeprajsshetty/go/src/github.com/deeprajsshetty/sde-software-architect/golang/002-golang-grpc-projects/001-go-programming-basics/0020-panic/panic.go:23 +0x138
main.main()
	/Users/deeprajsshetty/go/src/github.com/deeprajsshetty/sde-software-architect/golang/002-golang-grpc-projects/001-go-programming-basics/0020-panic/panic.go:13 +0x28
exit status 2
------------------------------------------------Output------------------------------------------
*/