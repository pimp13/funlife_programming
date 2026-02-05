package main

func removeByTarget[T comparable](slice []T, target T) []T {
	for i, v := range slice {
		if v == target {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func sit(a, b string) (int, int) {
	var correct int
	var found int
	var ar []rune
	var br []rune
	// midonam me len(a) == len(b)
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			correct++
		} else {
			ar = append(ar, rune(a[i]))
			br = append(br, rune(b[i]))
		}
	}

	lookup := make(map[rune]struct{})
	for _, ch := range br {
		lookup[ch] = struct{}{}
	}

	for _, ch := range ar {
		if _, ok := lookup[ch]; ok {
			found++
			br = removeByTarget(br, ch)
		}
	}

	return correct, found
}

func main() {
}
