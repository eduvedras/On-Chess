package main

import (
	"fmt"
	"testing"
)

func TestPawnPossibleMoves(t *testing.T) {
	cases := []struct {
		p   pawn
		out []position
	}{
		{
			p:   pawn{pos: position{1, 0}, team: white},
			out: []position{{3, 0}, {2, 0}},
		},
		{
			p:   pawn{pos: position{1, 3}, team: white},
			out: []position{{3, 3}, {2, 3}},
		},
		{
			p:   pawn{pos: position{1, 7}, team: white},
			out: []position{{3, 7}, {2, 7}},
		},
		{
			p:   pawn{pos: position{5, 0}, team: white},
			out: []position{{6, 1}},
		},
		{
			p:   pawn{pos: position{5, 3}, team: white},
			out: []position{{6, 2}, {6, 4}},
		},
		{
			p:   pawn{pos: position{5, 7}, team: white},
			out: []position{{6, 6}},
		},
		{
			p:   pawn{pos: position{3, 3}, team: white},
			out: []position{{4, 3}},
		},
		{
			p:   pawn{pos: position{6, 0}, team: black},
			out: []position{{4, 0}, {5, 0}},
		},
		{
			p:   pawn{pos: position{6, 3}, team: black},
			out: []position{{4, 3}, {5, 3}},
		},
		{
			p:   pawn{pos: position{6, 7}, team: black},
			out: []position{{4, 7}, {5, 7}},
		},
		{
			p:   pawn{pos: position{2, 0}, team: black},
			out: []position{{1, 1}},
		},
		{
			p:   pawn{pos: position{2, 3}, team: black},
			out: []position{{1, 2}, {1, 4}},
		},
		{
			p:   pawn{pos: position{2, 7}, team: black},
			out: []position{{1, 6}},
		},
		{
			p:   pawn{pos: position{3, 3}, team: black},
			out: []position{{2, 3}},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			board := initializeBoard()

			res := c.p.possibleMoves(board)

			if len(res) != len(c.out) {
				t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.p, c.p.possibleMoves(board)))
			}

			for j := 0; j < len(c.out); j++ {
				if res[j] != c.out[j] {
					t.Errorf(fmt.Sprintf("Expected %v for input %v but got %v", c.out, c.p, c.p.possibleMoves(board)))
				}
			}
		})
	}
}
