package main

import (
	"fmt"
	"log"
)

func is_valid(pass string) bool {
	if pass == "funlife" {
		return true
	}
	return false
}

func main() {

	fmt.Print("What is the password? ")
	var input string
	if _, err := fmt.Scan(&input); err != nil {
		log.Fatal(err)
	}

	if is_valid(input) {
		fmt.Println("Correct! Go inside!")
	} else {
		fmt.Println("Wrong.. you shall not enter!")
	}

}
