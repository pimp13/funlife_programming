package main

import "fmt"

func main() {
	fmt.Println("Hello This is one print")

	goto JumpHere

	fmt.Println("This is tow print")

JumpHere:
	fmt.Println("Juming!!")
}
