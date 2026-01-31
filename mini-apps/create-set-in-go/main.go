package main

import "fmt"

type Set = map[string]struct{}

func Add(s Set, v string) {
	s[v] = struct{}{}
}

func Contains(s Set, v string) bool {
	_, ok := s[v]
	return ok
}

func main() {
	newSet := Set{}
	Add(newSet, "Pouya")
	Add(newSet, "Ali")
	Add(newSet, "Ali")

	fmt.Println("my set", newSet)
	fmt.Println("contains", Contains(newSet, "Ali"))

	for s := range newSet {
		fmt.Println("My set is", s)
	}
}
