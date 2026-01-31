package main

import (
	"fmt"
)

/*
Set hast shabihe set dar Python
ya shabihe be HashMap in Rust
*/

type Set[T comparable] = map[T]struct{}

func Add[T comparable](s Set[T], v T) {
	s[v] = struct{}{}
}

func Contains[T comparable](s Set[T], v T) bool {
	_, ok := s[v]
	return ok
}

func main() {

	newSet := Set[string]{}
	Add(newSet, "Pouya")
	Add(newSet, "Ali")
	Add(newSet, "Ali")

	fmt.Println("my set", newSet)
	fmt.Println("contains", Contains(newSet, "Ali"))

	for s := range newSet {
		fmt.Println("My set is", s)
	}
}
