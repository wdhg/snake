package snake

import "math/rand"

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