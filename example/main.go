package main

import (
	"fmt"

	"github.com/wdhg/snake"
)

func render(game *snake.Game) {
	grid := snake.GetGrid(*game)
	for _, row := range grid {
		for _, cell := range row {
			if cell == snake.SnakeHeadCell || cell == snake.SnakeBodyCell {
				fmt.Printf("# ")
			} else if cell == snake.FoodCell {
				fmt.Printf("@ ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	game := snake.NewGame(4, 4, 0)
	// Move snake around so it eats food 3 times
	game.MoveRight()
	render(game)
	game.MoveRight()
	render(game)
	game.MoveDown()
	render(game)
	game.MoveDown()
	render(game)
	game.MoveLeft()
	render(game)
	game.MoveUp()
	render(game)
	game.MoveUp()
	render(game)
	game.MoveRight()
	render(game)
	game.MoveRight()
	render(game)
	// Last move moves snake off right side of screen
	if !game.MoveRight() {
		fmt.Println("Game over!")
	}
}
