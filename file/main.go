package main

import (
	"fmt"
	"os"
)

func main() {
	createAndWriteToFile()
	readingFromFile()
}

func createAndWriteToFile() {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	defer file.Close()

	// write to file
	file.WriteString("Hello World")
}

func readingFromFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	defer file.Close()

	// read file
	var data = make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	fmt.Println(string(data[:count]))
}
