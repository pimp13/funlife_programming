package main

import "fmt"

/// Faghat roye slice int kar mikone
// func removeElementFromSlice(index int, slice *[]int) {
// 	if index >= len(*slice) {
// 		return
// 	}
// 	*slice = append((*slice)[:index], (*slice)[index+1:]...)
// }

// / Generic hast va roye hameye datatype haye slice karmikone
func removeElementFromSlice[T comparable](index int, val *[]T) {
	if index >= len(*val) {
		return
	}
	*val = append((*val)[:index], (*val)[index+1:]...)
}

func main() {
	myslice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(myslice)
	removeElementFromSlice(3, &myslice)
	fmt.Println(myslice)

	mystr := []string{"a", "b", "c"}
	removeElementFromSlice(1, &mystr)
	fmt.Println(mystr)
}
