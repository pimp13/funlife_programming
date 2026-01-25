package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

var word_list = []string{
	"golang",
	"python",
	"typescript",
}

func getRandomWord() string {
	return word_list[rand.Intn(len(word_list))]
}

func initializeProverb(proverb string) string {
	var display string
	for _, charRune := range proverb {
		char := string(charRune)
		if char == " " {
			display += " "
		} else {
			display += "_"
		}
	}
	return display
}

func getUserInput() (string, error) {
	var input string
	if _, err := fmt.Scan(&input); err != nil {
		log.Fatal("failed to get input from stdin:", err)
	}
	trimedInput := strings.TrimSpace(input)
	strings.ToLower(trimedInput)
	if len(trimedInput) < 1 {
		return "", fmt.Errorf("kalame bayad bishtar az 1 harf bashe")
	}
	// if trimedInput == usedLetters {
	// 	return "", fmt.Errorf("shoma ghablan in kalame ro haads zadied")
	// }
	return trimedInput, nil
}

/*
def update_proverb(proverb, display, letter):
    new_display = ""
    for p_char, d_char in zip(proverb, display):
        if p_char == letter:  # اگر حرف درست باشد، در جای خود قرار می‌گیرد
            new_display += letter
        else:
            new_display += d_char  # اگر نادرست باشد، حالت قبلی حفظ می‌شود
    return new_display
*/
func updateProverb(proverb, display, letter string) string {
	// for 
	fmt.Println("proverb", proverb)
	fmt.Println("display", display)
	fmt.Println("letter", letter)
	return ""
}

func main() {

	fmt.Println("Please enter your guess word:")
	randWord := getRandomWord()
	strings.ToLower(randWord)
	hiddenWordDisplay := initializeProverb(randWord)

	chance := 10

	for chance > 0 {
		fmt.Println("the character word:")
		fmt.Println(hiddenWordDisplay)
		fmt.Printf("\nYour Guess (enter q for quit): ")

		guessLetter, err := getUserInput()
		if err != nil {
			log.Fatal(err)
			break
		}
		_ = updateProverb(randWord, hiddenWordDisplay, guessLetter)

		if guessLetter == "q" || guessLetter == "exit" {
			break
		}

		chance--
	}

}
