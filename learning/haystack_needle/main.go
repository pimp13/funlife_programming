package main

import "fmt"

func main() {
	haystack := "hello world my name is pouya"
	needle := "name"

	for _, v := range haystack {
		fmt.Println(string(v))
	}

	fmt.Println(len(haystack), needle)
}
