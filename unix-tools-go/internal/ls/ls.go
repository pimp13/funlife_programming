package ls

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"sort"
	"syscall"
	"time"

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
		for _, entry := range entries {
			printLongEntry(entry)
		}

	} else {
		printMultiColumn(entries)
	}

}

type Entry struct {
	IsDir bool
	Name  string
	E     os.DirEntry
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
		entryT[i] = &Entry{IsDir: isDir, Name: name, E: entry}
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
			fmt.Printf("\033[1m\033[32m%-*s\033[0m", colWidth, e.Name)
		} else {
			fmt.Printf("%-*s", colWidth, e.Name)
		}
		// fmt.Printf("%-*s ", colWidth, getFilenameWithColor(e.E))
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

func printLongEntry(entry os.DirEntry) {
	info, err := entry.Info()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ls: error reading file info: %v\n", err)
		return
	}

	stat := info.Sys().(*syscall.Stat_t)
	perms := info.Mode().Perm().String()

	if info.IsDir() {
		perms = "d" + perms[1:]
	} else {
		perms = "-" + perms[1:]
	}

	uid := fmt.Sprint(stat.Uid)
	gid := fmt.Sprint(stat.Gid)

	usr, err := user.LookupId(uid)
	if err != nil {
		usr = &user.User{Username: uid}
	}

	grp, err := user.LookupGroupId(gid)
	if err != nil {
		grp = &user.Group{Name: gid}
	}

	// Determine which time format to use for modification time display
	var timeFormat string
	sixMonths := time.Hour * 24 * 30 * 6 // approximate 6 months as 180 days

	if time.Since(info.ModTime()) > sixMonths || info.ModTime().After(time.Now()) {
		timeFormat = "Jan _2 2006"
	} else {
		timeFormat = "Jan _2 15:04"
	}

	// Print the long entry format, matching typical 'ls -l'
	fmt.Printf("%s %3d %s %s %8d %s %s\n",
		perms,
		stat.Nlink,
		usr.Username,
		grp.Name,
		info.Size(),
		info.ModTime().Format(timeFormat),
		getFilenameWithColor(entry),
	)

}

func getFilenameWithColor(entry os.DirEntry) string {
	if entry.IsDir() {
		return fmt.Sprintf("\033[1m\033[32m%s\033[0m", entry.Name())
	}
	return entry.Name()
}
