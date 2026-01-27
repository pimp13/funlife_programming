package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()

	cols, rows := termbox.Size()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	msg := fmt.Sprintf("Cols: %d | Rows: %d", cols, rows)

	x := (cols - len(msg)) / 2
	y := rows / 2

	for i, ch := range msg {
		termbox.SetCell(x+i, y, ch, termbox.ColorGreen, termbox.ColorBlack)
	}
	termbox.Flush()

	for {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc || ev.Ch == 'q' || ev.Ch == 'Q' {
				break
			}
		}
	}

	// termbox.SetCell(10, 10, '@', termbox.ColorGreen, termbox.ColorBlack)

}
