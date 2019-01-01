package snake

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetGrid(t *testing.T) {
	game := Game{
		Snake: []Vector{
			Vector{0, 0},
			Vector{1, 0},
			Vector{1, 1},
			Vector{1, 2},
		},
		Food:   Vector{2, 3},
		bounds: Vector{4, 4},
	}
	expected := [][]Cell{
		[]Cell{SnakeHeadCell, SnakeBodyCell, EmptyCell, EmptyCell},
		[]Cell{EmptyCell, SnakeBodyCell, EmptyCell, EmptyCell},
		[]Cell{EmptyCell, SnakeBodyCell, EmptyCell, EmptyCell},
		[]Cell{EmptyCell, EmptyCell, FoodCell, EmptyCell},
	}
	output := GetGrid(game)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("incorrect output, expected %v got %v", expected, output)
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		operations func(*Game)
		expected   []Vector
	}{
		{
			func(game *Game) {
				game.update(1, 0)
				game.update(1, 0)
				game.update(1, 0)
				game.update(0, 1)
			},
			[]Vector{Vector{3, 1}},
		},
		{
			func(game *Game) {
				game.update(1, 0)
				game.update(1, 0)
				game.update(0, 1)
				game.update(0, 1)
			},
			[]Vector{Vector{2, 2}, Vector{2, 1}},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			game := NewGame(4, 4, 0)
			test.operations(game)
			if !reflect.DeepEqual(test.expected, game.Snake) {
				t.Errorf("incorrect final snake, expected %v got %v", test.expected, game.Snake)
			}
		})
	}
}
