package main

import (
	"fmt"
)

func main() {
	word := "  Hello world  Pouya    "
	var l int

	// TODO: mishe ham trim kard ham else if l != 0 ghozasht
	// wordtrim := strings.TrimSpace(word)
	for i := len(word) - 1; i >= 0; i-- {
		currChar := string(word[i])
		if currChar != " " {
			l++
		} else if l != 0 {
			break
		}
	}

	fmt.Println(l)

}
