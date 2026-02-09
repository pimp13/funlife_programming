package main

import (
	"os"
	"unix-tools-go/internal/cat"
)

func main() {
	if len(os.Args) > 1 {
		cat.Run(os.Args[1:])
	} else {
		os.Stdout.WriteString("\n")
	}
}
