package main

import (
	"fmt"
	"slices"
)

func main() {

	// var sliceName []ElementType

	// Different ways of slice declaration
	var numbers []int
	var numbers1 = []int{1, 2, 3}
	numbers2 := []int{1, 2, 3}
	numbers3 := make([]int, 5)
	fmt.Println(numbers, numbers1, numbers2, numbers3)

	// Convert array to slice
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]
	fmt.Println("Converted slice from array:", slice)

	// Append slice
	slice = append(slice, 6, 7)
	fmt.Println("Appended slice:", slice)

	// Copy slice
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	fmt.Println("Copied slice:", copiedSlice)

	// Nil slice
	var nilSlice []int
	fmt.Println(nilSlice)

	// Slice - Range based iteration
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %v", i, v)
	}

	// Value change - index of slice 3
	fmt.Println("Element at index 3 of slice:", slice[3])
	slice[3] = 60
	fmt.Println("Element at index 3 of slice:", slice[3])

	// Equal slice
	equalSlice1 := []int{2,4,6}
	equalSlice2 := []int{2,4,6}
	fmt.Println(slices.Equal(equalSlice1, equalSlice2))
	// change equalSlice2 value
	equalSlice2[1] = 100
	fmt.Println(slices.Equal(equalSlice1, equalSlice2))

	// Dynamic two 2D Slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}
	fmt.Println("Dynamic 2D Array:", twoD)

	// slice[low:high]
	sliceOfSlice := slice[1:4]
	fmt.Println("Slice of slice:", sliceOfSlice)
	fmt.Println("The capacity of sliceOfSlice:", cap(sliceOfSlice))
	fmt.Println("The length of sliceOfSlice:", len(sliceOfSlice))
}

/*
----------------------------------------Output---------------------------
[] [1 2 3] [1 2 3] [0 0 0 0 0]
Converted slice from array: [2 3 4]
Appended slice: [2 3 4 6 7]
Copied slice: [2 3 4 6 7]
[]
Index: 0, Value: 2Index: 1, Value: 3Index: 2, Value: 4Index: 3, Value: 6Index: 4, Value: 7Element at index 3 of slice: 6
Element at index 3 of slice: 60
true
false
Adding value 0 in outer slice at index 0, and in inner slice index of 0
Adding value 1 in outer slice at index 1, and in inner slice index of 0
Adding value 2 in outer slice at index 1, and in inner slice index of 1
Adding value 2 in outer slice at index 2, and in inner slice index of 0
Adding value 3 in outer slice at index 2, and in inner slice index of 1
Adding value 4 in outer slice at index 2, and in inner slice index of 2
Dynamic 2D Array: [[0] [1 2] [2 3 4]]
Slice of slice: [3 4 60]
The capacity of sliceOfSlice: 7
The length of sliceOfSlice: 3
----------------------------------------Output---------------------------
*/