package main

import (
	"os"
	"unix-tools-go/internal/echo"
)

func main() {

	if len(os.Args) > 1 {
		echo.Run(os.Args[1:])
	} else {
		os.Stdout.WriteString("\n")
	}

}
