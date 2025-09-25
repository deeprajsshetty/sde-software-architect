package main

import "fmt"

func init()  {
	fmt.Println("Initializing package1...")
}

func init()  {
	fmt.Println("Initializing package2...")
}

func init()  {
	fmt.Println("Initializing package3...")
}

func main() {
	fmt.Println("Inside the main function")
}

/*
---------Output----------
Initializing package1...
Initializing package2...
Initializing package3...
Inside the main function
---------Output----------
*/