// package snake is a basic implementation of the snake game
package snake

type vector struct {
	x int
	y int
}

type snake []vector

// Direction represents the directions the snake can move
type Direction int

const (
	// Up represents the vector (0, -1)
	Up Direction = iota
	// Right represents the vector (1, 0)
	Right
	// Down represents the vector (0, 1)
	Down
	// Left represents the vector(-1, 0)
	Left
)

func (s *snake) extendHeadInDirection(dir Direction) {
	head := (*s)[0]
	switch dir {
	case Up:
		head.y++
	case Right:
		head.x++
	case Down:
		head.y--
	case Left:
		head.x--
	}
	*s = append([]vector{head}, *s...)
}

func (s *snake) shrinkBody() {
	// cuts off rear end of snake
	*s = (*s)[0 : len(*s)-1]
}

func snakeContainsCell(x, y int, s snake) bool {
	for _, cell := range s {
		if x == cell.x && y == cell.y {
			return true
		}
	}
	return false
}

func hasSnakeEatenSelf(s snake) bool {
	// checks if head is within the body of the snake
	return snakeContainsCell(s[0].x, s[0].y, s[1:])
}
