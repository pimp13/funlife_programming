package main

import (
	"os"
	"unix-tools-go/internal/ls"
)

func main() {
	ls.Run(os.Args[1:])
}
