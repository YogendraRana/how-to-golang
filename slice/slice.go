package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Printf("Before append, length = %v and capacity = %v\n", len(fruits), cap(fruits))

	fruits = append(fruits, "Date")
	fmt.Printf("After append, length = %v and capacity = %v", len(fruits), cap(fruits))

}
