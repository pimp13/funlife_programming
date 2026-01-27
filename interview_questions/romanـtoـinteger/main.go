package main

import "fmt"

func romanToInt(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && values[s[i]] < values[s[i+1]] {
			result -= values[s[i]]
		} else {
			result += values[s[i]]
		}
	}

	return result
}

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("IV"))      // 4
	fmt.Println(romanToInt("IX"))      // 9
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}
