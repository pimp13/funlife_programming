package main

import "fmt"

type number interface {
	~int | ~int64 | ~int32 | ~float32 | ~float64 |
		~uint | ~uint32 | ~uint64
}

func sum[T number](a T, b T) T {
	return a + b
}

func process[T any](value T) {
	/*
		mitavan meghdar value ke az type T yani generic hast ro check konim key
		type oon key hast
	*/
	switch v := any(value).(type) {
	case int:
		fmt.Println("value T is int", v)
	case string:
		fmt.Println("value T is string", v)
	default:
		fmt.Println("Unknown type (only string or int)")
	}
}

func main() {
	process(23)
	fmt.Println(sum(23, 23.2))
}
