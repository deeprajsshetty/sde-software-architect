package main

import "fmt"

func main() {

	message := "Hello World"

	for i, v := range message {
		fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
}

/*
------Output------
0 72
Index: 0, Rune: H
1 101
Index: 1, Rune: e
2 108
Index: 2, Rune: l
3 108
Index: 3, Rune: l
4 111
Index: 4, Rune: o
5 32
Index: 5, Rune:  
6 87
Index: 6, Rune: W
7 111
Index: 7, Rune: o
8 114
Index: 8, Rune: r
9 108
Index: 9, Rune: l
10 100
Index: 10, Rune: d
------Output------
*/