package main

import "fmt"

func main() {
	scores := make(map[string]uint)

	scores["maths"] = 100
	scores["science"] = 98
	scores["nepali"] = 90

	fmt.Println("Scores Map: ", scores)

	// add an element to the map
	scores["english"] = 95
	fmt.Println("Scores Map: ", scores)

	// delete an element from the map
	delete(scores, "nepali")
	fmt.Println("Scores Map: ", scores)

	// update an element in the map
	scores["maths"] = 99
	fmt.Println("Scores Map: ", scores)

	// check if an element exists in the map
	score, exists := scores["geography"]

	if !exists {
		fmt.Println("Geography score not found")
	} else {
		fmt.Println("Geography score: ", score)
	}

}
