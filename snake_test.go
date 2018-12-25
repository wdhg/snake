package snake

import (
	"fmt"
	"testing"
)

func TestShrinkBody(t *testing.T) {
	s := snake{vector{}, vector{}, vector{}, vector{}, vector{}}
	snakeLength := len(s)
	for expectedLength := snakeLength - 1; expectedLength >= 0; expectedLength-- {
		t.Run(fmt.Sprintf("%d", expectedLength), func(t *testing.T) {
			s.shrinkBody()
			if len(s) != expectedLength {
				t.Error("snake length not shrunk")
			}
		})
	}
}

func TestHasSnakeEatenSelf(t *testing.T) {
	tests := []struct {
		snake    []vector
		expected bool
	}{
		{
			[]vector{vector{0, 0}, vector{1, 0}, vector{1, 1}, vector{0, 1}},
			false,
		},
		{
			[]vector{vector{0, 0}, vector{1, 0}, vector{1, 1}, vector{0, 1}, vector{0, 0}},
			true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := hasSnakeEatenSelf(test.snake)
			if test.expected != output {
				t.Errorf("incorrect output, expected %t got %t", test.expected, output)
			}
		})
	}
}
