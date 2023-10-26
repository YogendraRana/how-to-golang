package main

import "fmt"

func printX(x int) {
	fmt.Println(x)
}

func printY(y *int) {
	fmt.Println(*y)
}

func main() {
	var x int = 10
	var y int = 50

	printX(x)
	printY(&y)
}
