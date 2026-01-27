package main

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

func makeWorld(rows, cols int) [][]rune {
	world := make([][]rune, rows)

	for h := 0; h < rows; h++ {
		world[h] = make([]rune, cols)

		for w := 0; w < cols; w++ {
			if rand.Float64() > 0.1 {
				world[h][w] = '.'
			} else {
				world[h][w] = ' '
			}
		}
	}

	return world
}

func main() {
	termbox.Init()
	defer termbox.Close()

	// width: cols
	// height: rows
	cols, rows := termbox.Size()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Make the game world
	world := makeWorld(rows, cols)

	// Draw the world
	for h := 0; h < rows; h++ {
		for w := 0; w < cols; w++ {
			termbox.SetCell(
				w,
				h,
				world[h][w],
				termbox.ColorGreen,
				termbox.ColorBlack,
			)
		}
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

}
