package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements.
	fruits := make([]string, 5, 10)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"
	fmt.Println(len(fruits))

	// You can't access an index of a slice beyond its length.
	// fruits[5] = "Runtime error"
	fruits = append(fruits, "watermelon")

	// Error: panic: runtime error: index out of range

	fmt.Println(fruits[:3])
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))
}
