package main

import (
	"iter"
	"slices"
)

type Iterator[V any] struct {
	seq iter.Seq[V]
}

func From[V any](src []V) *Iterator[V] {
	return &Iterator[V]{seq: slices.Values(src)}
}

func (it *Iterator[V]) Map(f func(V) V) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for v := range it.seq {
				if !yield(f(v)) {
					break
				}
			}
		},
	}
}

func main() {

}
