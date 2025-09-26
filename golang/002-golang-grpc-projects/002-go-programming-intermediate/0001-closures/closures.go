package main

import "fmt"

func main() {
	sequence := adder()

	fmt.Println("Adder: ", sequence())
	fmt.Println("Adder: ", sequence())
	fmt.Println("Adder: ", sequence())
	fmt.Println("Adder: ", sequence())

	sequence2 := adder()
	fmt.Println("Refresh adder: ", sequence2())
	fmt.Println("Refresh adder: ", sequence2())
	fmt.Println("Refresh adder: ", sequence2())

	subtractor := func () func(int) int {
		
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	// Using the clousre subtracter
	fmt.Println("Subtractor: ", subtractor(1))
	fmt.Println("Subtractor: ", subtractor(2))
	fmt.Println("Subtractor: ", subtractor(3))

}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

/*
----------Output---------
previous value of i: 0
added 1 to i
Adder:  1
added 1 to i
Adder:  2
added 1 to i
Adder:  3
added 1 to i
Adder:  4
previous value of i: 0
added 1 to i
Refresh adder:  1
added 1 to i
Refresh adder:  2
added 1 to i
Refresh adder:  3
Subtractor:  98
Subtractor:  96
Subtractor:  93
----------Output---------
*/