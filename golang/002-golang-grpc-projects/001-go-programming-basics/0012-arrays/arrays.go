package main

import "fmt"

func main() {

	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple" ,"Banana", "Orange", "Grapes"}
	fmt.Println("Fruites array:", fruits)

	fmt.Println("Third element:", fruits[2])

	// Arrays are value types
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copiedArray)

	// Iterate array using for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":",numbers[i])
	}

	// Array - Range based iteration
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %v\n", i, v)
	}

	// underscore is blank identifier, used to store unused values
	a, _ := someFunction()
	fmt.Println(a)	

	fmt.Println("The length of numbers array is", len(numbers))

	// Comparing Arrays
	array1 := [3]int{2, 4, 8}
	array2 := [3]int{2, 4, 8}
	fmt.Println("Array1 is equal to Array2:", array1 == array2)

	// Two dimentional array
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	// Array using pointers - reference type
	originalArrayPointer := [3]int{1, 2, 3}
	var copiedArrayPointer *[3]int
	copiedArrayPointer = &originalArrayPointer

	copiedArrayPointer[0] = 100

	fmt.Println("Original Array using pointer:", originalArrayPointer)
	fmt.Println("Copied Array using pointer:", copiedArrayPointer)
}

func someFunction() (int, int) {
	return 1, 2
}

/*
--------------------Output--------------------
[0 0 0 0 0]
[0 0 0 0 20]
[9 0 0 0 20]
Fruites array: [Apple Banana Orange Grapes]
Third element: Orange
Original Array: [1 2 3]
Copied Array: [100 2 3]
Element at index, 0 : 9
Element at index, 1 : 0
Element at index, 2 : 0
Element at index, 3 : 0
Element at index, 4 : 20
Index: 0, Value: 9
Index: 1, Value: 0
Index: 2, Value: 0
Index: 3, Value: 0
Index: 4, Value: 20
1
The length of numbers array is 5
Array1 is equal to Array2: true
[[1 2 3] [4 5 6] [7 8 9]]
Original Array using pointer: [100 2 3]
Copied Array using pointer: &[100 2 3]
--------------------Output--------------------
*/