package main

import "fmt"

func main() {
	process(10)
}

// While using defer arg passing should be careful
func process(i int)  {
	defer fmt.Println("Deferred value i:", i)
	defer fmt.Println("Deferred first time")
	defer fmt.Println("Deferred second time")
	defer fmt.Println("Deferred third time")
	i++
	fmt.Println("Normal execution")
	fmt.Println("Value of i:", i)
}

/*
------Output-------
Normal execution
Value of i: 11
Deferred third time
Deferred second time
Deferred first time
Deferred value i: 10
------Output-------
*/