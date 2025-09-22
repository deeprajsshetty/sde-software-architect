package main

import "fmt"

func main() {

	// func <name>(parameter list) returnType {
	//		code block
	//		return value
	//	}

	sum := add(6, 8)
	fmt.Println("Sum:", sum)

	// Anonymous function
	greet := func ()  {
		fmt.Println("Hello Anonymous function")
	}
	greet()

	// Functions as type
	operation := add
	result := operation(10, 12)
	fmt.Println("Function add as a type and result is:", result)

	// Passing a function as an argument
	result = applyOperation(20, 30, add)
	fmt.Println("20 + 30 =", result)

	// Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 =", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b;
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(x, y int) int) int  {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

/*
----------------Output------------------
Sum: 14
Hello Anonymous function
Function add as a type and result is: 22
20 + 30 = 50
6 * 2 = 12
----------------Output------------------
*/