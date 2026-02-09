package echo

import (
	"os"
	"strings"
)

func Run(args []string) {
	output := strings.Join(args, " ")
	if _, err := os.Stdout.WriteString(output + "\n"); err != nil {
		os.Exit(1)
	}
}
