package snake

import "math/rand"

type Cell int

const (
	EmptyCell Cell = iota
	SnakeCell
	FoodCell
)

// Game represents a single snake game
type Game struct {
	snake
	size vector
	food vector
	prng *rand.Rand
}

// NewGame initialises a new Game
func NewGame(width, height int, seed int64) *Game {
	game := Game{
		snake: []vector{vector{0, 0}},
		size:  vector{width, height},
		prng:  rand.New(rand.NewSource(seed)),
	}
	game.spawnFood()
	return &game
}

func (g *Game) GetSnakeHead() (int, int) {
	return g.snake[0].x, g.snake[0].y
}

func (g *Game) GetFood() (int, int) {
	return g.food.x, g.food.y
}

func (g *Game) GetCells() [][]Cell {
	cells := [][]Cell{}
	for y := 0; y < g.size.y; y++ {
		row := []Cell{}
		for x := 0; x < g.size.x; x++ {
			if g.food.x == x && g.food.y == y {
				row = append(row, FoodCell)
			} else if snakeContainsCell(x, y, g.snake) {
				row = append(row, SnakeCell)
			} else {
				row = append(row, EmptyCell)
			}
		}
		cells = append(cells, row)
	}
	return cells
}

// Update updates the game state by one move and returns whether the is snake alive
func (g *Game) Update(dir Direction) (alive bool) {
	g.snake.extendHeadInDirection(dir)
	if hasEatenFood(*g) {
		g.spawnFood()
	} else {
		g.snake.shrinkBody()
	}
	return isSnakeOutOfBounds(*g) || hasSnakeEatenSelf(g.snake)
}

func getEmptyCells(g Game) []vector {
	empty := []vector{}
	for x := 0; x < g.size.x; x++ {
		for y := 0; y < g.size.y; y++ {
			if !snakeContainsCell(x, y, g.snake) {
				empty = append(empty, vector{x, y})
			}
		}
	}
	return empty
}

func hasEatenFood(g Game) bool {
	head := g.snake[0]
	if head.x == g.food.x && head.y == g.food.y {
		return true
	}
	return false
}

func isSnakeOutOfBounds(g Game) bool {
	head := g.snake[0]
	return head.x < 0 || head.y < 0 || g.size.x <= head.x || g.size.y <= head.y
}

func (g *Game) spawnFood() {
	empty := getEmptyCells(*g)
	g.food = empty[g.prng.Intn(len(empty))]
}
