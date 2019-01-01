// ackage snake implements a basic snake game
package snake

import (
	"math/rand"
)

// Cell represents the contents of a cell in the 2D game world
type Cell int

const (
	SnakeHeadCell Cell = iota
	SnakeBodyCell
	FoodCell
	EmptyCell
)

// Vector represents a 2D position
type Vector struct {
	X int
	Y int
}

// Game represents a single instance of a game of snake
type Game struct {
	// Snake stores the cells that the snake occupies
	Snake []Vector
	// Food represents the position of the target food
	Food   Vector
	bounds Vector
	prng   *rand.Rand
}

// NewGame initalises a new Game object
func NewGame(width, height int, seed int64) *Game {
	game := Game{
		Snake:  []Vector{Vector{0, 0}},
		bounds: Vector{width, height},
		prng:   rand.New(rand.NewSource(seed)),
	}
	game.spawnFood()
	return &game
}

// GetGrid returns the 2D game world of cells
func GetGrid(game Game) [][]Cell {
	grid := [][]Cell{}
	for y := 0; y < game.bounds.Y; y++ {
		row := []Cell{}
		for x := 0; x < game.bounds.X; x++ {
			row = append(row, getCellType(Vector{x, y}, game))
		}
		grid = append(grid, row)
	}
	return grid
}

// MoveDown moves the snake head down
func (g *Game) MoveDown() (alive bool) {
	return g.update(0, 1)
}

// MoveLeft moves the snake head left
func (g *Game) MoveLeft() (alive bool) {
	return g.update(-1, 0)
}

// MoveRight moves the snake head right
func (g *Game) MoveRight() (alive bool) {
	return g.update(1, 0)
}

// MoveUp moves the snake head up
func (g *Game) MoveUp() (alive bool) {
	return g.update(0, -1)
}

func (g *Game) update(dx, dy int) (alive bool) {
	// extend head forward
	head := g.Snake[0]
	head.X += dx
	head.Y += dy
	g.Snake = append([]Vector{head}, g.Snake...)
	// check if eaten food
	if head.X == g.Food.X && head.Y == g.Food.Y {
		// has eaten food, spawn new food
		g.spawnFood()
	} else {
		// hasn't eaten food, shrink rear
		g.Snake = g.Snake[:len(g.Snake)-1]
	}
	// snake is alive if inside bounds and it hasn't eaten itself
	return isInsideBounds(g.Snake[0], g.bounds) && !isVectorContainedIn(g.Snake[0], g.Snake[1:])
}

func isInsideBounds(vec Vector, bounds Vector) (alive bool) {
	return vec.X < 0 || vec.Y < 0 || bounds.X <= vec.X || bounds.Y <= vec.Y
}

func (g *Game) spawnFood() {
	emptyCells := []Vector{}
	for y := 0; y < g.bounds.Y; y++ {
		for x := 0; x < g.bounds.X; x++ {
			if !isVectorContainedIn(Vector{x, y}, g.Snake) {
				emptyCells = append(emptyCells, Vector{x, y})
			}
		}
	}
	g.Food = emptyCells[g.prng.Intn(len(emptyCells))]
}

func getCellType(vec Vector, game Game) Cell {
	if vec.X == game.Food.X && vec.Y == game.Food.Y {
		return FoodCell
	}
	if vec.X == game.Snake[0].X && vec.Y == game.Snake[0].Y {
		return SnakeHeadCell
	}
	if isVectorContainedIn(vec, game.Snake[1:]) {
		return SnakeBodyCell
	}
	return EmptyCell
}

func isVectorContainedIn(vec Vector, vs []Vector) bool {
	for _, v := range vs {
		if vec.X == v.X && vec.Y == v.Y {
			return true
		}
	}
	return false
}
