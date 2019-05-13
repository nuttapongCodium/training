package main

import "fmt"

type user struct {
	name string
}

func main() {
	// users := make(map[string]user)
	users := map[string]user{}

	users["1150"] = user{
		name: "Sander",
	}

	users["1112"] = user{
		name: "MacDonald",
	}

	fmt.Printf("map users %+v\n", users)

	delete(users, "1150")

	user, ok := users["1150"]
	fmt.Printf("map users %+v %+v\n", user, ok)

}
