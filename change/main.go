package main

import "fmt"

func calculateChange(cost, amount int, cashProfile []int) []int {
	changeCash := []int{}
	change := amount - cost
	for change > 0 {
		for _, c := range cashProfile {
			if change >= c {
				changeCash = append(changeCash, c)
				change -= c
				break
			}
		}
	}
	return changeCash
}

func main() {
	cashProfile := []int{1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

	fmt.Println(calculateChange(215, 300, cashProfile))
	fmt.Println(calculateChange(486, 600, cashProfile))
	fmt.Println(calculateChange(12, 400, cashProfile))
}
