package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
	set1 := []int{}

	// Append numbers to the slice.
	set1 = append(set1, 1)

	// Display each value in the slice.
	fmt.Println(set1)

	// Declare a slice of strings and populate the slice with names.
	names := []string{
		"cabigon",
		"scizor",
		"dog",
		"cat",
	}

	// Display each index position and slice value.
	fmt.Println("Display each index position and slice value.")
	for i, name := range names {
		fmt.Printf("%v %s\n", i, name)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	slicedNames := make([]string, len(names[1:3]))
	copy(slicedNames, names[1:3])

	// Display each index position and slice values for the new slice.
	fmt.Println("Display each index position and slice values for the new slice.")
	for i, name := range slicedNames {
		fmt.Printf("%v %s\n", i, name)
	}

}
