package main

import "fmt"

func main() {

	// Factorial
	fmt.Println("Factorial of 3 is:", factorial(3))
	fmt.Println("Factorial of 10 is:", factorial(10))

	// Sum of digits
	fmt.Println("Sum of digit 9:", sumOfDigits(9))
	fmt.Println("Sum of digit 34:", sumOfDigits(34))
	fmt.Println("Sum of digit 234:", sumOfDigits(2346789))
}

func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}

	// Recursive case: factorial of n is n * factorial(n - 1)
	return n * factorial(n-1)
	// n * (n - 1) * (n - 2) * factorial(n - 3).... factorial(0)
}

func sumOfDigits(n int) int {
	// Base case
	if 	n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)
}

/*
----------Output------------
Factorial of 3 is: 6
Factorial of 10 is: 3628800
Sum of digit 9: 9
Sum of digit 34: 7
Sum of digit 234: 39
----------Output------------
*/
