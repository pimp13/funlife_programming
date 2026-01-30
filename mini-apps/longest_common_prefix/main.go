package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	// strs := []string{"b", "a", "c"}

	sort.Strings(strs)
	var ans string
	for i := 1; i < len(strs[0])-1; i++ {
		lookingFor := strs[0][0:i]
		for _, s := range strs[1:] {
			if len(s) < i || s[0:i] != lookingFor {
				return
			}
		}
		ans = lookingFor
	}
	fmt.Println(ans)
}
