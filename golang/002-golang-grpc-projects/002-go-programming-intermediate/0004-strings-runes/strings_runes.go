package main

import "fmt"

func main() {

	message1 := "Hello, \nGo!"
	message2 := "Hello, \tGo!"
	message3 := "Hello, \rGo!"			// Go!lo
	rawMessage := `Hello, \n\t\rGo!`

	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)

	// Length of string
	fmt.Println("Length of message1 variable is", len(message1))
	fmt.Println("Length of message2 variable is", len(message2))
	fmt.Println("Length of message3 variable is", len(message3))
	fmt.Println("Length of rawMessage variable is", len(rawMessage))

	fmt.Println("The first char in message1 var is", message1[0])		// ASCII

	// String concat
	greeting := "Hello, "
	name := "John"
	fmt.Println(greeting + name)

	// String compare
	str1 := "Apple"					// A has an ASCII value of 85
	str := "apple"					// a has an ASCII value of 97
	str2 := "banana"				// b has an ASCII value of 98
	str3 := "app"						// a has an ASCII value of 97
	fmt.Println(str1 < str2)		// true
	fmt.Println(str3 < str1)    // false
	fmt.Println(str > str1)			// true
	fmt.Println(str > str3)			// true

	// Print character
	for i, char := range message1 {
		fmt.Printf("Character at index %d is %c\n", i, char)
	}

	// Print Hexadecimal value and default value i.e. ASCII
	for i, char := range message1 {
		fmt.Printf("Hexadecimal at index %d is %x\n", i, char)			// HEXADECIMAL
		fmt.Printf("Default value [ASCII] at index %d is %v\n", i, char)
	}

	// Strings are Immutable
	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	// Rune
	var ch rune = 'a'
	jch := '日'
	fmt.Println("Rune value:", ch)
	fmt.Println("Japanese character[日] rune value:", jch)

	// Print real value
	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	// Convert rune to string
	cstr := string(ch)
	fmt.Println("Convert rune a to string:", cstr)
	fmt.Printf("Type of cstr is %T\n", cstr)

	// Iterating over runes
	const ARIGATO = "ありがとう"		// Japanese Word
	fmt.Println(ARIGATO)

	for _, runeValue := range ARIGATO {
		fmt.Printf("%c\n", runeValue)
	}

	// Smily
	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}

/*
---------------------Output----------------
Hello,
Go!
Hello,  Go!
Go!lo,
Hello, \n\t\rGo!
Length of message1 variable is 11
Length of message2 variable is 11
Length of message3 variable is 11
Length of rawMessage variable is 16
The first char in message1 var is 72
Hello, John
true
false
true
true
Character at index 0 is H
Character at index 1 is e
Character at index 2 is l
Character at index 3 is l
Character at index 4 is o
Character at index 5 is ,
Character at index 6 is
Character at index 7 is

Character at index 8 is G
Character at index 9 is o
Character at index 10 is !
Hexadecimal at index 0 is 48
Default value [ASCII] at index 0 is 72
Hexadecimal at index 1 is 65
Default value [ASCII] at index 1 is 101
Hexadecimal at index 2 is 6c
Default value [ASCII] at index 2 is 108
Hexadecimal at index 3 is 6c
Default value [ASCII] at index 3 is 108
Hexadecimal at index 4 is 6f
Default value [ASCII] at index 4 is 111
Hexadecimal at index 5 is 2c
Default value [ASCII] at index 5 is 44
Hexadecimal at index 6 is 20
Default value [ASCII] at index 6 is 32
Hexadecimal at index 7 is a
Default value [ASCII] at index 7 is 10
Hexadecimal at index 8 is 47
Default value [ASCII] at index 8 is 71
Hexadecimal at index 9 is 6f
Default value [ASCII] at index 9 is 111
Hexadecimal at index 10 is 21
Default value [ASCII] at index 10 is 33
Hello, John
Rune value: 97
Japanese character[日] rune value: 26085
a
日
Convert rune a to string: a
Type of cstr is string
ありがとう
あ
り
が
と
う
128522
😊
---------------------Output----------------
*/