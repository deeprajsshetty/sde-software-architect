package main

import "fmt"

func main() {
	// Simple iteration over a loop
	for i := 0; i <= 5; i++ {
		fmt.Println("loop:", i)
	}
	
	// iteration over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// iteration using continue and break
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue	// continue the loop but skip the rest of the lines/statements
		}
		fmt.Println("Odd Number:", i)
		if i == 5 {
			break	// break out of the loop
		}
	}

	// ASTERISK LAYOUT
	rows := 5
	// outer loop
	for i := 1; i <= rows; i++ {
		// inner loop for spaces before stars
		for j := 1; j <= rows - i; j++ {
			fmt.Print(" ")
		}
		// inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()	// Move to the next line
	}

	// Go 1.22 update
	for i := range 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println("We have a lift off!")
}

/*
---------Output----------
loop: 0
loop: 1
loop: 2
loop: 3
loop: 4
loop: 5
Index: 0, Value: 1
Index: 1, Value: 2
Index: 2, Value: 3
Index: 3, Value: 4
Index: 4, Value: 5
Index: 5, Value: 6
Odd Number: 1
Odd Number: 3
Odd Number: 5
    *
   ***
  *****
 *******
*********
1
2
3
4
5
6
7
8
9
10
We have a lift off!
---------Output----------
*/
