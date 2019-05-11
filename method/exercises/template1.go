// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

// user defines a user in the program.

// notify implements a method with a value receiver.

// changeEmail implements a method with a pointer receiver.

func main() {

	// Values of type user can be used to call methods
	// declared with both value and pointer receivers.

	// Pointers of type user can also be used to call methods
	// declared with both value and pointer receiver.

	// Create a slice of user values with two users.

	// Iterate over the slice of users switching
	// semantics. Not Good!

	// Exception example: Using pointer semantics
	// for a collectoin of strings.

}
