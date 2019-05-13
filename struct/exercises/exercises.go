package main

// Add imports
import (
	"encoding/json"
	"fmt"
)

type struct1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	// Declare variable of type user and init using a struct literal.
	person1 := struct1{
		Name: "Nuttapong",
		Age:  23,
	}

	// Display the field values.
	fmt.Printf("%+v \n", person1)

	// Declare a variable using an anonymous struct.
	person2 := struct {
		Name string
		Age  int
	}{
		Name: "KK",
		Age:  12,
	}

	// Display the field values.
	fmt.Printf("%+v\n", person2)

	// Declare variable of type product public and init using a struct literal.
	// set struct tag
	person3 := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "KK",
		Age:  12,
	}

	// Call json.Mashal package encode/json
	person3Json, _ := json.Marshal(person3)

	// Display the field values.
	fmt.Printf("%+v\n", string(person3Json))

}
