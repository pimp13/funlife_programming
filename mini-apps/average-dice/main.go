package main

import (
	"fmt"
	"math/rand/v2"
)

func dice() int {
	return rand.IntN(6) + 1
}

func avg(nums []int) float64 {
	var sum float64
	for _, n := range nums {
		sum += float64(n)
	}
	return sum / float64(len(nums))
}

func main() {
	var awnser []int
	for i := 0; i < 100000; i++ {
		n := 1
		for dice() != 6 {
			n++
		}
		awnser = append(awnser, n)
	}
	fmt.Printf("%.1f\n", avg(awnser)) // 6.0
}
