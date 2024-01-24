package main

import "fmt"

func main() {
	// Declare a variable
	var number int = 42

	// Declare a pointer variable
	var pointerToNumber *int

	// Assign the address of the variable to the pointer
	pointerToNumber = &number

	// Print the value and address of the variable
	fmt.Println("Value of number:", number)
	fmt.Println("Address of number:", &number)

	// Print the value and address stored in the pointer
	fmt.Println("Value pointed by pointerToNumber:", *pointerToNumber)
	fmt.Println("Address stored in pointerToNumber:", pointerToNumber)

	// Modify the value indirectly through the pointer
	*pointerToNumber = 100

	// Print the modified value of the variable
	fmt.Println("Modified value of number:", number)
}
