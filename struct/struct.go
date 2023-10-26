package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var Alex Person = Person{
		Name: "Alex",
		Age:  20,
	}

	John := Person{
		Name: "John",
		Age:  21,
	}

	// print the struct
	fmt.Println(Alex)
	fmt.Println(John.Name, John.Age)
}
