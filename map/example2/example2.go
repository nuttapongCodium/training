package main

import "fmt"

type user struct {
	name string
}

func main() {
	users := map[string]user{
		"1150": user{
			name: "Sander",
		},
	}

	users["1112"] = user{
		name: "MacDonald",
	}

	users["191"] = user{
		name: "จ่าเฉย",
	}

	keys := []string{}
	for key, user := range users {
		fmt.Printf("map key: %s users %+v\n", key, user)
		keys = append(keys, key)
	}

	fmt.Printf("keys: %+v\n", keys)
}
