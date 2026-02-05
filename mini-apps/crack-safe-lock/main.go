package main

import (
	"fmt"
)

func removeByTarget[T comparable](slice []T, target T) []T {
	for i, v := range slice {
		if v == target {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func sit(a, b string) (int, int) {
	var correct int
	var found int
	var ar []rune
	var br []rune
	// midonam me len(a) == len(b)
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			correct++
		} else {
			ar = append(ar, rune(a[i]))
			br = append(br, rune(b[i]))
		}
	}

	lookup := make(map[rune]struct{})
	for _, ch := range br {
		lookup[ch] = struct{}{}
	}

	for _, ch := range ar {
		if _, ok := lookup[ch]; ok {
			found++
			br = removeByTarget(br, ch)
		}
	}

	return correct, found
}

func main() {
	for i := 0; i < 1000; i++ {
		strI := fmt.Sprintf("%03d", i)
		if fix, found := sit(strI, "682"); fix != 1 || found != 0 {
			continue
		}
		if fix, found := sit(strI, "614"); fix != 0 || found != 1 {
			continue
		}
		if fix, found := sit(strI, "206"); fix != 0 || found != 2 {
			continue
		}
		if fix, found := sit(strI, "738"); fix != 0 || found != 0 {
			continue
		}
		if fix, found := sit(strI, "380"); fix != 0 || found != 1 {
			continue
		}

		fmt.Println(strI)

	}
}
