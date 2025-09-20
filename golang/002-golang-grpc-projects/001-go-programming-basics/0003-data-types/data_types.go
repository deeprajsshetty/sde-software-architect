package main

import "fmt"

func main() {
	fmt.Println("Hello, " + "World!")
	fmt.Println("9 * 10 = ", 9*10)
	fmt.Println("180.18/2.0 = ", 180.18/2.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

/*
------Output----------
Hello, World!
9 * 10 =  90
180.18/2.0 =  90.09
false
true
false
------Output----------
*/