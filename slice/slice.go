package main

import "fmt"

func main() {
	// create a slice with make()
	slice := make([]int, 1, 8)
	fmt.Println("Slice:", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// append elements to the slice
	slice = append(slice, 1, 2, 3)
	fmt.Println("Slice:", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))
}
