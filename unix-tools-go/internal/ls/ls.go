package ls

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"syscall"
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
