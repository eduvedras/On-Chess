package main

import (
	"fmt"
	"testing"
)

func TestQueenPossibleMoves(t *testing.T) {
	cases := []struct {
		q   queen
		out []position
	}{
		{
			q:   queen{pos: position{2, 0}, team: white},
			out: []position{{2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {3, 1}, {4, 2}, {5, 3}, {6, 4}},
		},
		{
			q:   queen{pos: position{3, 4}, team: white},
			out: []position{{3, 5}, {3, 6}, {3, 7}, {3, 3}, {3, 2}, {3, 1}, {3, 0}, {4, 4}, {5, 4}, {6, 4}, {2, 4}, {4, 5}, {5, 6}, {6, 7}, {2, 5}, {4, 3}, {5, 2}, {6, 1}, {2, 3}},
		},
		{
			q:   queen{pos: position{5, 7}, team: black},
			out: []position{{5, 6}, {5, 5}, {5, 4}, {5, 3}, {5, 2}, {5, 1}, {5, 0}, {4, 7}, {3, 7}, {2, 7}, {1, 7}, {4, 6}, {3, 5}, {2, 4}, {1, 3}},
		},
		{
			q:   queen{pos: position{0, 0}, team: black},
			out: []position{{0, 1}, {1, 0}, {1, 1}},
		},
		{
			q:   queen{pos: position{7, 7}, team: black},
			out: []position{},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			board := initializeBoard()

			res := c.q.possibleMoves(board)

			if len(res) != len(c.out) {
				t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.q, c.q.possibleMoves(board)))
			}

			for j := 0; j < len(c.out); j++ {
				if res[j] != c.out[j] {
					t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.q, c.q.possibleMoves(board)))
				}
			}
		})
	}
}
