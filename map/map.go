package main

import "fmt"

func main() {
	var ages map[string]int = make(map[string]int)
	ages["Alice"] = 30
	ages["Alex"] = 25
	ages["Bob"] = 23

	// access a map element
	var age1 = ages["Alice"]
	fmt.Printf("Alice is %d years old\n", age1)

	// loop through a map
	for name, age := range ages {
		fmt.Printf("%v\t%v\n", name, age)
	}

	// delete a map element
	delete(ages, "Alice")

	// check if a map element exists
	var age2, ok = ages["Alice"]
	if ok {
		fmt.Printf("Alice is %d years old\n", age2)
	} else {
		fmt.Println("Alice not found")
	}
}
