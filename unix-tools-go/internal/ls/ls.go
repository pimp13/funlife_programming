package ls

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"syscall"

	"golang.org/x/term"
)

func Run(args []string) {
	fs := flag.NewFlagSet("ls", flag.ExitOnError)
	longFormat := fs.Bool("l", false, "use a long listing formant")
	fs.Parse(args)

	dir := "."
	if fs.NArg() > 0 {
		dir = fs.Arg(0)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ls: cannot access '%s': %v\n", dir, err)
		os.Exit(1)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	if *longFormat {
		printTotalBlocks(entries)

	} else {
		printMultiColumn(entries)
	}

}

func printMultiColumn(entries []os.DirEntry) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 60
	}

	var names []string

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("\033[1m\033[36m%s\033[0m  ", entry.Name())
		} else {
			fmt.Printf("%s  ", entry.Name())
		}
	}
	println()
}

func printTotalBlocks(entries []os.DirEntry) {
	var totalBlocks int64

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		stat := info.Sys().(*syscall.Stat_t)
		totalBlocks += stat.Blocks
	}

	fmt.Printf("total %dK\n", totalBlocks/2)
}
