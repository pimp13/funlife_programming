package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./persian_dict_19k.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 0 {
			continue
		}

		word := strings.TrimSpace(parts[0])
		words = append(words, word)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	wordSet := make(map[string]struct{}, len(words))
	for _, w := range words {
		wordSet[w] = struct{}{}
	}
	// value, exists := wordSet["گیرایی"]
	// fmt.Println(value, exists)
	// os.Exit(1)

	for _, word := range words {
		manWord := fmt.Sprintf("من%s", word)
		if _, exists := wordSet[manWord]; exists {
			fmt.Printf("Word: %s , ManWord: %s\n", word, manWord)
		}
	}

}

// buf, err := io.ReadAll(file)
// if err != nil {
// 	log.Fatal(err)
// }

// content := string(buf)
// word := strings.Split(content, ":")[0]
// fmt.Println(word)

// reader := csv.NewReader(file)
// var words []string
//** baraye file haye bozoorg
// for {
// 	reacord, err := reader.Read()
// 	if err != nil {
// 		break
// 	}
// 	for _, line := range reacord {
// 		word := strings.Split(line, ":")[0]
// 		words = append(words, word)
// 	}
// }
// for _, word := range words {
// 	fmt.Println(word)
// }

//** baraye file haye kochak
// records, err := reader.ReadAll()
// if err != nil {
// 	log.Fatal(err)
// }
// for _, row := range records {
// 	for _, line := range row {
// 		fmt.Println(line)
// 	}
// }
