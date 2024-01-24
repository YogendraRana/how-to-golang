package main

import "fmt"

type NumType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func min[T NumType](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(min(1, 2))
	fmt.Println(min(2.1, 2.2))
}
