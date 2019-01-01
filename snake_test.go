package snake

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetGrid(t *testing.T) {
	game := Game{
		Snake: []vector{
			vector{0, 0},
			vector{1, 0},
			vector{1, 1},
			vector{1, 2},
		},
		Food:   vector{2, 3},
		bounds: vector{4, 4},
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
		expected   []vector
	}{
		{
			func(game *Game) {
				game.update(1, 0)
				game.update(1, 0)
				game.update(1, 0)
				game.update(0, 1)
			},
			[]vector{vector{3, 1}},
		},
		{
			func(game *Game) {
				game.update(1, 0)
				game.update(1, 0)
				game.update(0, 1)
				game.update(0, 1)
			},
			[]vector{vector{2, 2}, vector{2, 1}},
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
