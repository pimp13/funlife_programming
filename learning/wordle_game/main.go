package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	freq := "esiarntolcdugpmhbyfvkwzxjq"

	freqScore := make(map[byte]int)
	for i, ch := range freq {
		freqScore[byte(ch)] = len(freq) - i
	}

	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var bestWord string
	var maxScore int
	var wordScores []struct {
		word  string
		score int
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if len(word) == 5 {
			indexTowWord := string(word[2])
			if !strings.ContainsAny(word, "arnse") {
				if indexTowWord == "i" && string(word[0]) == "c" && string(word[1]) == "l" {
					fmt.Println("Word by index 2 i =>", word)
				}
			}
		}

		if word == "" || len(word) != 5 {
			continue
		}
		wordSlice := make([]string, len(word))
		for i := 0; i < len(word); i++ {
			wordSlice[i] = string(word[i])
		}
		uniqueWordSlice := lo.Uniq(wordSlice)
		if len(uniqueWordSlice) != 5 {
			continue
		}

		var score int
		for i := 0; i < len(word); i++ {
			if s, exists := freqScore[word[i]]; exists {
				score += s
			}
		}

		wordScores = append(wordScores, struct {
			word  string
			score int
		}{word, score})

		if score > maxScore {
			maxScore = score
			bestWord = word
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(wordScores, func(i, j int) bool {
		return wordScores[i].score > wordScores[j].score
	})

	fmt.Println("Kalame ba bishtarin harfe tekrari:")
	fmt.Println("====================================")
	for i := 0; i < 20 && i < len(wordScores); i++ {
		fmt.Printf("%2d. %-15s | Score: %d\n", i+1, wordScores[i].word, wordScores[i].score)
	}

	fmt.Printf("\bBestWord : %s (Score: %d)\n", bestWord, maxScore)

}
