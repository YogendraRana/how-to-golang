package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 100
	var denominator int = 0

	// call the division function
	var quotient, remainder, err = division(numerator, denominator)

	// check if an error was returned
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The quotient is %v", quotient)
	default:
		fmt.Printf("The quotient is %v and the remainder is %v", quotient, remainder)
	}

}

func division(a int, b int) (int, int, error) {
	var err error

	if b == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	var quotient int = a / b
	var remainder int = a % b

	return quotient, remainder, err
}
