package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func generateRandomNumber(ta ...int) int {
	var rndTa int
	if len(ta) > 0 {
		rndTa = ta[0]
	} else {
		rndTa = 100
	}
	return rand.Intn(rndTa)
}

func getUserInput() (int, error) {
	var strinput string
	if _, err := fmt.Scan(&strinput); err != nil {
		return -1, fmt.Errorf("lotfan number vared kon.")
	}

	toint, err := strconv.Atoi(strinput)
	if err != nil {
		return -1, err
	}
	return toint, nil
}

func main() {
	fmt.Println("Welcome to haadse number!!")

	this_org_num := generateRandomNumber()
	chanes := 8
	var att int
	var isok bool
	for {
		fmt.Printf("\nLotfan ye number az 0 ta 100 vared konid (baraye khoroj -1): ")

		userInput, err := getUserInput()
		if err != nil {
			log.Fatal(err)
		}
		if userInput == -1 {
			fmt.Println("goodbye!!!")
			break
		}
		if chanes == 0 {
			fmt.Printf("Shoma %d bar talash kardi vali natonesti :(\n", att)
			break
		}
		chanes--

		if userInput > this_org_num {
			att++
			fmt.Println("number bozorg hast")
		}
		if userInput < this_org_num {
			att++
			fmt.Println("number kochik hast")
		}
		if userInput == this_org_num {
			att++
			isok = true
			fmt.Println("eval dorost haads zadi.")
			break
		}
	}
	if isok {
		log.Printf("\nShoma dar %d talash tonesti number ro haads bezani :)", att)
	}
}
