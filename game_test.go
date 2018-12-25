package snake

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetEmptyCells(t *testing.T) {
	tests := []struct {
		snake    []vector
		expected []vector
	}{
		{
			[]vector{vector{0, 0}, vector{1, 0}},
			[]vector{vector{0, 1}, vector{1, 1}},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := getEmptyCells(Game{
				size:  vector{2, 2},
				snake: test.snake,
			})
			if !reflect.DeepEqual(test.expected, output) {
				t.Errorf("incorrect output, expected %v got %v", test.expected, output)
			}
		})
	}
}

func TestIsSnakeOutOfBounds(t *testing.T) {
	tests := []struct {
		snake    []vector
		expected bool
	}{
		{[]vector{vector{0, 0}}, false},
		{[]vector{vector{0, -1}}, true},
		{[]vector{vector{-1, 0}}, true},
		{[]vector{vector{5, 0}}, true},
		{[]vector{vector{4, 0}}, false},
		{[]vector{vector{0, 5}}, true},
		{[]vector{vector{0, 4}}, false},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			game := Game{
				snake: test.snake,
				size:  vector{5, 5},
			}
			output := isSnakeOutOfBounds(game)
			if test.expected != output {
				t.Errorf("incorrect output, expected %t got %t", test.expected, output)
			}
		})
	}
}
