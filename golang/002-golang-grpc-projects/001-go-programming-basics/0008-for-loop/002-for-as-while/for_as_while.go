package main

import "fmt"

func main() {

	// for as while with break
	sum := 0
	for {
		sum += 10
		fmt.Println("Sum:", sum)
		if sum >= 50 {
			break
		}
	}

	// for as while with continue
	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Number:", num)
		num++	// ++ increment operator increase value by 1 and -- decrement operator, decrease value by 1
	}
}

/*
---------Output----------
Sum: 10
Sum: 20
Sum: 30
Sum: 40
Sum: 50
Odd Number: 1
Odd Number: 3
Odd Number: 5
Odd Number: 7
Odd Number: 9
---------Output----------
*/