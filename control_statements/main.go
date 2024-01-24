package main

import (
	"fmt"
)

func main() {
	var name string = "John"
	var age int = 17

	// check if the person can vote
	if age >= 18 {
		fmt.Printf("%v can vote.\n", name)
	} else {
		fmt.Printf("%v cannot vote.\n", name)
	}

	// for loop
	fmt.Println("For loop:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}
}
