package main

import "fmt"

func main() {
	// Declare and make a map of integer type values.
	points := map[string]int{}
	points["Seven-Eleven"] = 100
	points["TRUE"] = 1

	// Display each value in the slice.
	fmt.Println(points)

	// Initialize some data into the map.
	hps := map[string]int{
		"Cabigon": 500,
		"Pikachu": 1,
	}

	// Display each key/value pair
	for key, value := range hps {
		fmt.Printf("%s %d\n", key, value)
	}
}
