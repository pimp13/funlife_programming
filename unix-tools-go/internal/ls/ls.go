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

type Entry struct {
	IsDir bool
	Name  string
}

func printMultiColumn(entries []os.DirEntry) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 60
	}

	entryT := make([]*Entry, len(entries))
	var maxLen int
	for i, entry := range entries {
		name := entry.Name()
		isDir := entry.IsDir()
		entryT[i] = &Entry{IsDir: isDir, Name: name}
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	colWidth := maxLen + 2
	cols := width / colWidth
	if cols == 0 {
		cols = 1
	}

	for i, e := range entryT {
		if e.IsDir {
			fmt.Printf("\033[1m\033[36m%-*s\033[0m", colWidth, e.Name)
		} else {
			fmt.Printf("%-*s", colWidth, e.Name)
		}
		if (i+1)%cols == 0 || i == len(entryT)-1 {
			println()
		}
	}

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
