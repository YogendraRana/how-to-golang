package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// getFunction()
	// postFunction()
	putFunction()
	// deleteFunction()
}

// get function
func getFunction() {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making HTTP request: %v", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response body: %v", err)
	}

	fmt.Printf(string(body))
}

// post function
func postFunction() {
	var url string = "https://jsonplaceholder.typicode.com/posts"
	payload := []byte(`{"title":"Post Title","body":"Post Body","userId":1}`)
	res, err := http.Post(url, "application/json; charset=UTF-8", bytes.NewBuffer(payload))

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer res.Body.Close()

	// read the response body
	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Printf(string(body))
}

// put function
func putFunction() {
	var url string = "https://jsonplaceholder.typicode.com/posts/1"
	payload := []byte(`{"title":"Put Title","body":"Put Body","userId":1}`)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer res.Body.Close()

	// read the response body
	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Printf(string(body))
}

// delete function
func deleteFunction() {
	var url string = "https://jsonplaceholder.typicode.com/posts/1"
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		fmt.Printf(err.Error())
		panic(err.Error())
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf(err.Error())
		panic(err.Error())
	}

	defer res.Body.Close()

	// read the response body
	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Printf(string(body))
}
