package cat

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

var width int

func Run(args []string) {
	var err error
	// width, _, err = term.GetSize(int(syscall.Stdin))
	width, _, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
		// fmt.Fprintf(os.Stderr, "cat: error reading input: %v\n", err)
		// os.Exit(1)
	}

	fs := flag.NewFlagSet("cat", flag.ExitOnError)
	isShowLineNumber := fs.Bool("n", false, "print by line numbers")
	fs.Parse(args)

	files := fs.Args()

	if len(files) == 0 {
		printFromReader(os.Stdin, isShowLineNumber)
	} else {
		for _, file := range files {
			if err := printFile(file, isShowLineNumber); err != nil {
				fmt.Fprintf(os.Stderr, "cat: error reading input: %v\n", err)
				os.Exit(1)
			}
		}
	}
}

func printFile(filename string, lineNumber *bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	printFromReader(file, lineNumber)
	return nil
}

func printFromReader(reader *os.File, isShowLineNumber *bool) {
	scanner := bufio.NewScanner(reader)

	/*
		hamashon ba \033[0m tamom mishan
		BOLD: \033[1m
		Muted: \033[2m
		Italic: \033[3m
		Underline: \033[4m
		Red: \033[31m
		Green: \033[32m
		Yellow: \033[33m
		Blue: \033[34m
		Arghavani: \033[35m
		Cyan: \033[36m
		Sefid: \033[37m
		BgGreen: \033[42m
		BgRed: \033[41m
		BgBlue: \033[44m
		BgCyan: \033[46m
	*/

	if reader == os.Stdin {
		fmt.Printf("> \033[32mReading from stdin\033[0m\n\n")
	} else {
		// fmt.Println(
		// 	strings.Repeat("──", width),
		// )
		fmt.Printf("\033[2m%s\033[0m\n", strings.Repeat("─", width-1))
		print("\tFile: ")
		fmt.Printf(
			"\033[32m%s\033[0m\n", reader.Name(),
		)
		fmt.Printf("\033[2m%s\033[0m\n", strings.Repeat("─", width-1))
	}

	lineCounter := 1
	for scanner.Scan() {
		if *isShowLineNumber {
			fmt.Printf("\033[2m%6d │\033[0m %s\n", lineCounter, scanner.Text())
			lineCounter++
		} else {
			fmt.Printf("\t%s\n", scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "cat: error reading input: %v\n", err)
		os.Exit(1)
	}
}
