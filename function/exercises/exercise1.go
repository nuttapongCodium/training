// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

// height rooms status_eia
// Add imports.
import "fmt"

type condo struct {
	Name    string
	Floors  int
	Rooms   int
	PassEIA bool
}

func (condo_ *condo) validateEIA() {
	condo_.PassEIA = true
}

func (condo_ *condo) roomsPerFloor() float32 {
	return float32(condo_.Rooms) / float32(condo_.Floors)
}

func main() {
	condo1 := condo{
		Name:    "Condo1",
		Floors:  10,
		Rooms:   200,
		PassEIA: false,
	}

	fmt.Println(condo1)

	condo1.validateEIA()
	fmt.Println(condo1)

	fmt.Println(condo1.roomsPerFloor())
}
