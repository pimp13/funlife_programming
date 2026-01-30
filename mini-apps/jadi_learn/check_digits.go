package main

import (
	"fmt"
	"slices"
)

// set check mikone ke dar yek array ya slice value tekrari nabashe agar bod pak mikone!
// faghat ba int slice kar mikone
func set[T comparable](val []T) []T {
	var unique []T
	for _, item := range val {
		if !slices.Contains(unique, item) {
			unique = append(unique, item)
		}
	}
	return unique
}

func hasConditions(i int) bool {
	var digits []int
	for l := 0; l < 3; l++ {
		digit := i % 10
		digits = set(append(digits, digit))
		i /= 10
		// fmt.Printf("digit => %d \n", digit)
	}
	// if slices.Contains(digits, 0) || len(digits) != 2 {
	// 	return false
	// }
	// return true
	return !(slices.Contains(digits, 0) || len(digits) != 2)
}

func main() {
	var result []int
	for i := 100; i < 999; i++ {
		if hasConditions(i) {
			result = append(result, i)
		}
	}
	fmt.Println(result)
	fmt.Printf("result of => %d \n", len(result))

	// u := []string{"pouya", "pouya", "ali", "ali"}
	// var final []string
	// for _, v := range u {
	// 	final = set(append(final, v))
	// }
	// fmt.Println(final)
}
