package main

import (
	"fmt"
	"testing"
)

func TestRookPossibleMoves(t *testing.T) {
	cases := []struct {
		r   rook
		out []position
	}{
		{
			r:   rook{pos: position{2, 0}, team: white},
			out: []position{{2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7}, {3, 0}, {4, 0}, {5, 0}},
		},
		{
			r:   rook{pos: position{3, 4}, team: white},
			out: []position{{3, 5}, {3, 6}, {3, 7}, {3, 3}, {3, 2}, {3, 1}, {3, 0}, {4, 4}, {5, 4}, {2, 4}},
		},
		{
			r:   rook{pos: position{5, 7}, team: black},
			out: []position{{5, 6}, {5, 5}, {5, 4}, {5, 3}, {5, 2}, {5, 1}, {5, 0}, {4, 7}, {3, 7}, {2, 7}},
		},
		{
			r:   rook{pos: position{0, 0}, team: black},
			out: []position{},
		},
		{
			r:   rook{pos: position{7, 7}, team: white},
			out: []position{},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			board := initializeBoard()

			res := c.r.possibleMoves(board)
			for j := 0; j < len(res); j++ {
				if res[j] != c.out[j] {
					t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.r, c.r.possibleMoves(board)))
				}
			}
		})
	}
}
