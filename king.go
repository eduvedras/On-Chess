package main

import "fmt"

type king struct {
	pos  position
	team team
}

func (k king) getTeam() team {
	return k.team
}

func (k king) getPos() position {
	return k.pos
}

func (k *king) setPos(pos position) {
	k.pos = pos
}

func (k king) possibleMoves(board board) []position {
	positions := []position{}
	directions := []direction{right, left, up, down, upRight, downRight, upLeft, downLeft}

	for _, dir := range directions {
		moves := iterateDirection(&k, dir, board)
		if len(moves) > 0 {
			positions = append(positions, moves[0])
		}
	}

	return positions
}

func (k king) toString() string {
	return fmt.Sprintf("k%v", k.getTeam())
}

func (k king) isEmpty() bool {
	return false
}
