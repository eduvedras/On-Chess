package main

import (
	"fmt"
	"testing"
)

func TestUpdateBoard(t *testing.T) {
	cases := []struct {
		p        chessPiece
		newPos   position
		oldPos   position
		outEnded bool
	}{
		{
			p:        &pawn{pos: position{1, 0}, team: white},
			newPos:   position{3, 0},
			oldPos:   position{1, 0},
			outEnded: false,
		},
		{
			p:        &queen{pos: position{3, 0}, team: white},
			newPos:   position{6, 0},
			oldPos:   position{3, 0},
			outEnded: false,
		},
		{
			p:        &rook{pos: position{7, 3}, team: white},
			newPos:   position{7, 4},
			oldPos:   position{7, 3},
			outEnded: true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			board := initializeBoard()

			outBoard, ended := updateBoard(board, c.p, c.newPos)

			if !outBoard[c.oldPos[0]][c.oldPos[1]].isEmpty() {
				t.Errorf(fmt.Sprintf("Expected position %v to be empty", c.oldPos))
			}

			if outBoard[c.newPos[0]][c.newPos[1]] != c.p {
				t.Errorf(fmt.Sprintf("Expected position %v to have %v but got %v", c.newPos, c.p, outBoard[c.newPos[0]][c.newPos[1]]))
			}

			if ended != c.outEnded {
				t.Errorf(fmt.Sprintf("Expected ended to be %v but got %v", c.outEnded, ended))
			}
		})
	}
}
