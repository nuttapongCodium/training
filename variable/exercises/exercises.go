package main

// Add imports
import "fmt"

func main() {
	// Declare variables that are set to their zero value.
	var zero int

	// Display the value of those variables.
	fmt.Printf("zero : %v \n", zero)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	varable1 := "somethings"
	bool1 := true

	// Display the value of those variables.
	fmt.Printf("%v %v\n", varable1, bool1)

	// Perform a type conversion.
	variable2 := float32(zero)

	// Display the value of that variable.
	fmt.Printf("variable2 : %v %T \n", variable2, variable2)
}
