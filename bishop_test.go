package main

import (
	"fmt"
	"testing"
)

func TestBishopPossibleMoves(t *testing.T) {
	cases := []struct {
		b   bishop
		out []position
	}{
		{
			b:   bishop{pos: position{2, 0}, team: white},
			out: []position{{3, 1}, {4, 2}, {5, 3}, {6, 4}},
		},
		{
			b:   bishop{pos: position{3, 4}, team: white},
			out: []position{{4, 5}, {5, 6}, {6, 7}, {2, 5}, {4, 3}, {5, 2}, {6, 1}, {2, 3}},
		},
		{
			b:   bishop{pos: position{5, 7}, team: black},
			out: []position{{4, 6}, {3, 5}, {2, 4}, {1, 3}},
		},
		{
			b:   bishop{pos: position{0, 2}, team: black},
			out: []position{{1, 3}, {1, 1}},
		},
		{
			b:   bishop{pos: position{7, 7}, team: black},
			out: []position{},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			board := initializeBoard()

			res := c.b.possibleMoves(board)

			if len(res) != len(c.out) {
				t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.b, c.b.possibleMoves(board)))
			}

			for j := 0; j < len(c.out); j++ {
				if res[j] != c.out[j] {
					t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.b, c.b.possibleMoves(board)))
				}
			}
		})
	}
}
