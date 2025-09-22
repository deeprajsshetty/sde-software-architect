package main

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(parameter1 type1, parameter2 type2, ...) (returnType1, returnType2, ...) {
	// 	code block
	//	return returnValue1, returnValue2, ...
	// }

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d. Remainder: %d\n", q, r)

	nq, nr := divideForNamed(12, 5)
	fmt.Printf("Quotient: %d. Remainder: %d\n", nq, nr)

	result, err := compare(2, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

// Multiple return values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Multiple named return values
func divideForNamed(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// Multiple return values with error
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}

/*
-----------Output----------
Quotient: 3. Remainder: 1
Quotient: 2. Remainder: 2
b is greater than a
-----------Output----------
*/