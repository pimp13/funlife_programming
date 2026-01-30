package main

import "fmt"

func main() {
	instructions := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	position := 50
	zeroCount := 0

	for _, instruction := range instructions {
		direction := string(instruction[0])   // L , R
		distance := int(instruction[1] - '0') // فاصله

		if direction == "R" {
			position = (position + distance) % 100
		}

		if direction == "L" {
			position = (position - distance) % 100
		}

		if position == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}
