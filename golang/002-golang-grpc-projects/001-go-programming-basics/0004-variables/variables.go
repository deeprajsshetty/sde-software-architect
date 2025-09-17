package main

import "fmt"

var middleName = "Cane"

func main() {
	var age int
	var name string = "John"
	var name1 = "Jane"

	count := 10
	lastName := "Smith"

	fmt.Println("Age: ", age, "Name: ", name, "Name1: ", name1, "Count: ", count, " LastName: ", lastName)

	fmt.Println(middleName)
	middleName = "Cane Updated"
	fmt.Println(middleName)


	// Default values
	// Numeric Types: 0
	// Boolean Types: false
	// String Type: ""
	// Pointers, slices, maps, functions, and structs: nil

	// ----Scope
	printName()
}

func printName() {
	firstName := "Michel"
	fmt.Println(firstName)
}