package main

import (
	"fmt"
	"testing"
)

func TestKingPossibleMoves(t *testing.T) {
	cases := []struct {
		k   king
		out []position
	}{
		{
			k:   king{pos: position{2, 0}, team: white},
			out: []position{{2, 1}, {3, 0}, {3, 1}},
		},
		{
			k:   king{pos: position{3, 4}, team: white},
			out: []position{{3, 5}, {3, 3}, {4, 4}, {2, 4}, {4, 5}, {2, 5}, {4, 3}, {2, 3}},
		},
		{
			k:   king{pos: position{5, 7}, team: black},
			out: []position{{5, 6}, {4, 7}, {4, 6}},
		},
		{
			k:   king{pos: position{0, 0}, team: black},
			out: []position{{0, 1}, {1, 0}, {1, 1}},
		},
		{
			k:   king{pos: position{7, 7}, team: black},
			out: []position{},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			board := initializeBoard()

			res := c.k.possibleMoves(board)
			for j := 0; j < len(res); j++ {
				if res[j] != c.out[j] {
					t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.k, c.k.possibleMoves(board)))
				}
			}
		})
	}
}
