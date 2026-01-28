package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type Food struct {
	Row int
	Col int
	Ch  rune
	Age time.Duration
}

func spawnFood(world [][]rune, count int) []Food {
	rows := len(world)
	cols := len(world[0])

	foods := make([]Food, 0, count)

	for len(foods) < count {
		r := random.Intn(rows)
		c := random.Intn(cols)

		if world[r][c] == ' ' {
			foods = append(foods, Food{
				Row: r,
				Col: c,
				Ch:  '🍎',
				Age: time.Second,
			})
		}
	}

	return foods
}

func eatFood(foods []Food, row, col int) ([]Food, bool) {
	for i := 0; i < len(foods); i++ {
		if foods[i].Row == row && foods[i].Col == col {
			foods = append(foods[:i], foods[i+1:]...)
			return foods, true
		}
	}
	return foods, false
}

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

func draw(world [][]rune, playerRow, playerCol int, foods []Food) {
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

	for _, food := range foods {
		termbox.SetCell(food.Col, food.Row, food.Ch, termbox.ColorYellow, termbox.ColorBlack)
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

	foods := spawnFood(world, 15)

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
	var score int
	var eaten bool

	// Loop Game
	for playing {

		select {
		// 🎮 INPUT (non-blocking)
		case evt := <-eventChan:
			switch evt.Type {
			case termbox.EventKey:
				var moved bool

				// Game Logic
				switch evt.Key {
				case termbox.KeyEsc:
					playing = false
					moved = false

				case termbox.KeyArrowUp:
					if playerRows > 0 && world[playerRows-1][playerCols] == ' ' {
						playerRows--
						moved = true
					}
				case termbox.KeyArrowDown:
					if playerRows < rows-1 && world[playerRows+1][playerCols] == ' ' {
						playerRows++
						moved = true
					}
				case termbox.KeyArrowLeft:
					if playerCols > 0 && world[playerRows][playerCols-1] == ' ' {
						playerCols--
						moved = true
					}
				case termbox.KeyArrowRight:
					if playerCols < cols-1 && world[playerRows][playerCols+1] == ' ' {
						playerCols++
						moved = true
					}

				default:
					if evt.Ch == 'q' || evt.Ch == 'Q' {
						playing = false
						moved = false
					}
				}

				if moved {
					foods, eaten = eatFood(foods, playerRows, playerCols)
					if eaten {
						score++
					}
				}

			case termbox.EventResize:
				cols, rows = evt.Width, evt.Height
				world = makeWorld(rows, cols, wallDensity)
				playerRows, playerCols = placePlayer(world)
				foods = spawnFood(world, 10)
			}

		// Tick / FPS control
		case <-ticker.C:
			// draw the world
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw(world, playerRows, playerCols, foods)

		}

		if len(foods) < 5 {
			foods = append(foods, spawnFood(world, 1)...)
		}

	}

}
