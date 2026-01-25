package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return 1
	}
	// prev := nums[0]
	// iw := 1
	// k := 1
	for ir := 1; ir < len(nums); ir++ {

	}
	return 0
}


func main() {
	nums := []int{1, 2, 2, 3, 4, 4}

	fmt.Println(removeDuplicates(nums))
}