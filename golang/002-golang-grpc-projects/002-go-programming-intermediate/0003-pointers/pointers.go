package main

import "fmt"

func main() {
	var ptr	*int
	var a int = 10
	ptr = &a			// referencing

	fmt.Println(a)
	fmt.Println(ptr)	
	
	// dereferencing the pointer
	fmt.Println(*ptr)
	//if ptr == nil {
	//	fmt.Println("Pointer is nil")
	//}

	modifyValue(ptr)
	fmt.Println("Pointer ptr modified:", a)

}

func modifyValue(ptr *int)  {
	*ptr++
}

/*
----------Output---------
10
0xc00000a0c8
10
Pointer ptr modified: 11
----------Output---------
*/