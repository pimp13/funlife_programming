package main

import (
	"fmt"
	"iter"
	"slices"
)

type Iterator[V any] struct {
	seq iter.Seq[V]
}

func From[V any](src []V) *Iterator[V] {
	return &Iterator[V]{seq: slices.Values(src)}
}

// Static return type return type Faghat mishe type iter
func (it *Iterator[V]) Map(f func(V) V) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for v := range it.seq {
				if !yield(f(v)) {
					return
				}
			}
		},
	}
}

// Diynamic return type
func Map[T, U any](it *Iterator[T], f func(T) U) *Iterator[U] {
	return &Iterator[U]{
		seq: func(yield func(U) bool) {
			for v := range it.seq {
				if !yield(f(v)) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Filter(pred func(V) bool) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for v := range it.seq {
				if !pred(v) {
					continue
				}
				if !yield(v) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Collect() []V {
	return slices.Collect(it.seq)
}

func (it *Iterator[V]) ForEach(f func(V)) {
	for v := range it.seq {
		f(v)
	}
}

type User struct {
	ID   uint
	Name string
}

func main() {
	myUser := []User{
		{Name: "Pouya", ID: 1},
		{Name: "Ali", ID: 2},
		{Name: "Hassan", ID: 3},
	}
	result := From(myUser).
		Map(func(u User) User { fmt.Printf("user: %s\n", u.Name); return u })

	// result := From(myUser)

	// names := Map(result, func(u User) string { return fmt.Sprintf("%s-", u.Name) }).
	fmt.Println(result)
}
