package main

import "fmt"

func main() {
	myNum := 23
	var ptr = &myNum

	// Actual Value by being referenced inside that pointer
	fmt.Println("Value of ptr is ", *ptr)
	// Value of Pointer
	fmt.Println("Value of ptr is ", ptr)
	*ptr = *ptr + 2
	fmt.Print("New Value of ptrvalue is ", ptr)
}
