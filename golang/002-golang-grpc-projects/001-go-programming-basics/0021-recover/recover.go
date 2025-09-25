package main

import "fmt"

func main() {

	process()
	fmt.Println("Returned from process")
}

func process()  {
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")
	fmt.Println("End Process")
}

/*
------------Output-------------
Start Process
Recovered: Something went wrong
Returned from process
------------Output-------------
*/