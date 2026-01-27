package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func makeWorld(rows, cols int, density float64) [][]rune {
	world := make([][]rune, rows)

	for h := 0; h < rows; h++ {
		world[h] = make([]rune, cols)

		for w := 0; w < cols; w++ {
			if random.Float64() < density {
				world[h][w] = '.'
			} else {
				world[h][w] = ' '
			}
		}
	}

	return world
}

func draw(world [][]rune, playerRow, playerCol int) {
	rows := len(world)
	if rows == 0 {
		return
	}
	cols := len(world[0])

	for h := 0; h < rows; h++ {
		for w := 0; w < cols; w++ {
			char := world[h][w]
			color := termbox.ColorLightCyan
			if h == playerRow && w == playerCol {
				char = '🛸'
				color = termbox.ColorLightRed
			}

			termbox.SetCell(w, h, char, color, termbox.ColorBlack)
		}
	}
	termbox.Flush()
}

func placePlayer(world [][]rune) (int, int) {
	rows := len(world)
	if rows == 0 {
		return 0, 0
	}
	cols := len(world[0])

	for {
		r := random.Intn(rows)
		c := random.Intn(cols)

		if world[r][c] == ' ' {
			return r, c
		}
	}
}

func main() {
	termbox.Init()
	defer termbox.Close()

	// width: cols
	// height: rows
	cols, rows := termbox.Size()

	// Make the game world
	const wallDensity = 0.1
	world := makeWorld(rows, cols, wallDensity)

	// place player
	playerRows, playerCols := placePlayer(world)

	eventChan := make(chan termbox.Event)
	go func() {
		for {
			ev := termbox.PollEvent()
			eventChan <- ev
		}
	}()
	playing := true
	const fps = 10
	// For Stable the Game (baraye sabeet bodan soraat bazi)
	ticker := time.NewTicker(time.Second / fps)
	defer ticker.Stop()

	// Loop Game
	for playing {
		select {
		// 🎮 INPUT (non-blocking)
		case evt := <-eventChan:
			switch evt.Type {
			case termbox.EventKey:
				switch evt.Key {
				case termbox.KeyEsc:
					playing = false

					// Game Logic
				case termbox.KeyArrowUp:
					if playerRows > 0 && world[playerRows-1][playerCols] == ' ' {
						playerRows--
					}
				case termbox.KeyArrowDown:
					if playerRows < rows-1 && world[playerRows+1][playerCols] == ' ' {
						playerRows++
					}
				case termbox.KeyArrowLeft:
					if playerCols > 0 && world[playerRows][playerCols-1] == ' ' {
						playerCols--
					}
				case termbox.KeyArrowRight:
					if playerCols < cols-1 && world[playerRows][playerCols+1] == ' ' {
						playerCols++
					}

				default:
					if evt.Ch == 'q' || evt.Ch == 'Q' {
						playing = false
					}
				}

			case termbox.EventResize:
				cols, rows = evt.Width, evt.Height
				world = makeWorld(rows, cols, wallDensity)
				playerRows, playerCols = placePlayer(world)
			}

		// Tick / FPS control
		case <-ticker.C:
			// draw the world
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw(world, playerRows, playerCols)

		}

	}

}
