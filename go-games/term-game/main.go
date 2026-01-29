package main

import (
	"fmt"
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

type Enemy struct {
	Row int
	Col int
	Ch  rune
}

func randomChoice[T comparable](slice []T) T {
	return slice[random.Intn(len(slice))]
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

func spawnEnemy(world [][]rune, count int) []Enemy {
	rows := len(world)
	cols := len(world[0])

	enemies := make([]Enemy, 0, count)
	for len(enemies) < count {
		r := random.Intn(rows)
		c := random.Intn(cols)
		if world[r][c] == ' ' {
			enemies = append(enemies, Enemy{
				Row: r,
				Col: c,
				Ch:  '👾',
			})
		}
	}

	return enemies
}

func enemyLogic(enemies []Enemy, playerRow, playerCol int) ([]Enemy, bool) {
	for i := 0; i < len(enemies); i++ {

		// Move random enemy
		if random.Float64() > 0.2 {

			if enemies[i].Row > playerRow {
				enemies[i].Row--
			} else if enemies[i].Col > playerCol {
				enemies[i].Col--
			} else if enemies[i].Row < playerRow {
				enemies[i].Row++
			} else if enemies[i].Col < playerCol {
				enemies[i].Col++
			}

			// enemies[i].Row += randomChoice([]int{0, 1, -1})
			// enemies[i].Col += randomChoice([]int{0, 1, -1})
		}

		// Died player
		if enemies[i].Row == playerRow && enemies[i].Col == playerCol {
			enemies = append(enemies[:i], enemies[i+1:]...)
			return enemies, true
		}

	}
	return enemies, false
}

func moveEnemy(enemies []Enemy) {
	for i := 0; i < len(enemies); i++ {
		// Move random enemy
		if random.Float64() > 0.7 {
			enemies[i].Row += randomChoice([]int{0, 1, -1})
			enemies[i].Col += randomChoice([]int{0, 1, -1})
		}
	}
}

func eatFood(foods []Food, row, col int) ([]Food, bool) {
	for i := 0; i < len(foods); i++ {
		if foods[i].Row == row && foods[i].Col == col {
			foods = append(foods[:i], foods[i+1:]...)
			return foods, true
		}
	}
	// TODO: Vaghti food khoorde shod bayad ye spawnFood beshe dobare
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

func drawText(x, y int, text string, fg, bg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}
}

func draw(world [][]rune, playerRow, playerCol int, foods []Food, score int, enemies []Enemy) {
	rows := len(world)
	if rows == 0 {
		return
	}
	cols := len(world[0])

	// +1 mikonim height player va foods ro ke 1-vahed bishter beshe ta score jaa beshe
	hud := fmt.Sprintf("Score: %d", score)
	drawText(0, 0, hud, termbox.ColorWhite|termbox.AttrBold, termbox.ColorBlue)

	for h := 0; h < rows; h++ {
		for w := 0; w < cols; w++ {
			char := world[h][w]
			color := termbox.ColorLightCyan
			if h == playerRow && w == playerCol {
				char = '🛸'
				color = termbox.ColorLightRed
			}

			termbox.SetCell(w, h+1, char, color, termbox.ColorBlack)
		}
	}

	for _, food := range foods {
		termbox.SetCell(food.Col, food.Row+1, food.Ch, termbox.ColorYellow, termbox.ColorBlack)
	}

	for _, enemy := range enemies {
		termbox.SetCell(enemy.Col, enemy.Row+1, enemy.Ch, termbox.ColorDefault, termbox.ColorBlack)
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
	enemies := spawnEnemy(world, 12)

	eventChan := make(chan termbox.Event)
	go func() {
		for {
			ev := termbox.PollEvent()
			eventChan <- ev
		}
	}()
	var playing bool = true
	const fps = 10
	// For Stable the Game (baraye sabeet bodan soraat bazi)
	ticker := time.NewTicker(time.Second / fps)
	defer ticker.Stop()
	var score int = 1
	var eaten bool
	var gameOver bool

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
						score += 10
					}
					enemies, gameOver = enemyLogic(enemies, playerRows, playerCols)
					if gameOver {
						termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
						// drawText(cols, rows+5, "You Died", termbox.ColorWhite|termbox.AttrBold, termbox.ColorBlue)
						time.Sleep(time.Second)
						playing = false
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
			draw(world, playerRows, playerCols, foods, score, enemies)

		}

		if len(foods) < 5 {
			foods = append(foods, spawnFood(world, 1)...)
		}
		if score <= 0 {
			playing = false
		}

	}

}
