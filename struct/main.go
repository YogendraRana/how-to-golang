package main

import "fmt"

type Person struct {
	Name    string
	Age     uint
	Address string
}

// Define the GetName method for the Person type
// this method has a receiver of type Person
// this method can be called with any instance of Person
func (p Person) GetName() string {
	return p.Name
}

func main() {
	yogendra := Person{
		Name:    "Yogendra",
		Age:     25,
		Address: "Pokhara",
	}

	fmt.Println(yogendra.Name, yogendra.Age, yogendra.Address)

	// Correct method call
	fmt.Println(yogendra.GetName())
}
