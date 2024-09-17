package main

import (
	"fmt"
	"testing"
)

func TestKnightPossibleMoves(t *testing.T) {
	cases := []struct {
		k   knight
		out []position
	}{
		{
			k:   knight{pos: position{2, 0}, team: white},
			out: []position{{3, 2}, {4, 1}},
		},
		{
			k:   knight{pos: position{3, 4}, team: black},
			out: []position{{4, 6}, {2, 6}, {4, 2}, {2, 2}, {5, 5}, {5, 3}, {1, 5}, {1, 3}},
		},
		{
			k:   knight{pos: position{5, 7}, team: black},
			out: []position{{4, 5}, {3, 6}},
		},
		{
			k:   knight{pos: position{0, 0}, team: black},
			out: []position{{1, 2}, {2, 1}},
		},
		{
			k:   knight{pos: position{7, 7}, team: black},
			out: []position{{5, 6}},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			board := initializeBoard()

			res := c.k.possibleMoves(board)

			if len(res) != len(c.out) {
				t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.k, c.k.possibleMoves(board)))
			}

			for j := 0; j < len(c.out); j++ {
				if res[j] != c.out[j] {
					t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.k, c.k.possibleMoves(board)))
				}
			}
		})
	}
}
