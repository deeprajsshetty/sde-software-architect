package main

import "fmt"

func main() {

	// ... Ellipsis
	// func functionName(param1 type1, param2 ...type2) returnType {
	// 	function body
	// }

	statement, total := sum("The sum of 20, 99 and 123 is", 20, 99, 123)
	fmt.Println(statement, total)

	sequenceStatement1, sequenceTotal1 := sumBasedOnSequence(1, 23, 50, 42, 55, 90)
	fmt.Printf("Sequence: %d. Total: %d\n", sequenceStatement1, sequenceTotal1)

	sequenceStatement2, sequenceTotal2 := sumBasedOnSequence(2, 231, 150, 412, 515, 190)
	fmt.Printf("Sequence: %d. Total: %d\n", sequenceStatement2, sequenceTotal2)

	numbers := []int{23, 76, 90, 666}
	sequenceStatement3, sequenceTotal3 := sumBasedOnSequence(3, numbers...)
	fmt.Printf("Sequence: %d. Total: %d\n", sequenceStatement3, sequenceTotal3)
}

// Regular parameter and variadic parameter
func sum(returnString string, nums ...int) (string, int) {
	total := 0

	for _, v := range nums {
		total += v
	}

	return returnString, total
}

// Sum of sequence using sequence parameter and variadic parameter
func sumBasedOnSequence(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}

/*
-----------Output---------------
The sum of 20, 99 and 123 is 242
Sequence: 1. Total: 260
Sequence: 2. Total: 1498
Sequence: 3. Total: 855
-----------Output---------------
*/