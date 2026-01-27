package main

import "fmt"

func canConstruct(ransomNote, magazine string) bool {
	charCount := make(map[rune]int)

	for _, ch := range magazine {
		charCount[ch]++
	}

	for _, ch := range ransomNote {
		if charCount[ch] == 0 {
			return false
		}
		charCount[ch]--
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))     // false
	fmt.Println(canConstruct("aa", "ab"))   // false
	fmt.Println(canConstruct("aa", "aab"))  // true
	fmt.Println(canConstruct("hello", "helloworld")) // true
}