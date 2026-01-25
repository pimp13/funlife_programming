package main

import (
	"fmt"
)

func isPrime(n int) bool {
	prime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
		}
	}
	return prime
}

func main() {
	var primeCount int
	for i := 1; i <= 10000; i++ {
		if isPrime(i) {
			primeCount++
		}
	}
	fmt.Println("\nprint count", primeCount)

}
