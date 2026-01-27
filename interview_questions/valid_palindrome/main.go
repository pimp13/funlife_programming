package main

import "unicode"
import "fmt"

func isValidPailndrome(s string) bool {
	si, ei := 0, len(s)-1
	for si < ei {
		for si < ei && !unicode.IsLetter(rune(s[si])) && !unicode.IsDigit(rune(s[si])) {
			si++
		}
		for si < ei && !unicode.IsLetter(rune(s[ei])) && !unicode.IsDigit(rune(s[ei])) {
			ei--
		}

		if si < ei && unicode.ToLower(rune(s[si])) != unicode.ToLower(rune(s[ei])) {
			return false
		}

		si++
		ei--
	}

	return true
}

func main() {
	// s := "A man, a plan, a canal: Panama"

	strs := []string{
		"A man, a plan, a canal: Panama",
		" ",
		"0P",
		"Grg",
		",,,,",
	}

	for _, str := range strs {
		fmt.Printf("Input: \"%s\" -> isValidPailndrome: %v\n", str, isValidPailndrome(str))
	}

}
