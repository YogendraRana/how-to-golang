package main

import "fmt"

func main() {
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"

	var numbers = [3]int{1, 2, 3}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	var i int
	for i = 0; i < len(numbers); i++ {
		fmt.Println(i)
	}
}
