package main

import "fmt"

func main() {

	// Printing Functions
	fmt.Print("Hello")
	fmt.Print("World!")
	fmt.Print(12, 456)

	fmt.Println("Hello")
	fmt.Println("World!")
	fmt.Println(12, 45)

	name := "John"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %x\n", age, age)

	// Formatting Functions
	s := fmt.Sprint("Hello", "World!", 123, 456)
	fmt.Print(s)

	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age %d", name, age)
	fmt.Println(sf)
	fmt.Println(sf)

	// Scanning functions
	var scanName string
	var scanAge int

	fmt.Print("Enter your name and age:")
	fmt.Scan(&scanName, &scanAge)
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	fmt.Print("Enter your name and age:")
	fmt.Scanln(&scanName, &scanAge)
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	fmt.Print("Enter your name and age:")
	fmt.Scanf("%s %d", &scanName, &scanAge)
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Error formatting functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}

/*
---------------Output--------------
HelloWorld!12 456Hello
World!
12 45
Name: John, Age: 25
Binary: 11001, Hex: 19
HelloWorld!123 456Hello World! 123 456
Hello World! 123 456
Name: John, Age 25
Name: John, Age 25
Enter your name and age:Alan

34
Name: John, Age: 25
Enter your name and age:Name: John, Age: 25
Enter your name and age:Dev 60
Name: John, Age: 25
Error: Age 15 is too young to drive.
---------------Output--------------
*/