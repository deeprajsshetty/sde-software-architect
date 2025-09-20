package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int 
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		// Check if the guess if correct
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}
}

/*
---------------------Output----------------------
Welcome to the Guessing Game!
I have chosen a number between 1 and 100
Can you guess what it is?
Enter your guess: 
20
Too low! Try guessing a higher number.
Enter your guess: 
40
Too high! Try guessing a lower number.
Enter your guess: 
30
Too high! Try guessing a lower number.
Enter your guess: 
25
Too low! Try guessing a higher number.
Enter your guess: 
28
Too high! Try guessing a lower number.
Enter your guess: 
27
Too high! Try guessing a lower number.
Enter your guess: 
26
Congratulations! You guessed the correct number!
---------------------Output----------------------
*/