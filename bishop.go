package main

import "fmt"

type bishop struct {
	pos  position
	team team
}

func (b bishop) getTeam() team {
	return b.team
}

func (b bishop) getPos() position {
	return b.pos
}

func (b *bishop) setPos(pos position) {
	b.pos = pos
}

func (b bishop) possibleMoves(board board) []position {
	positions := []position{}
	directions := []direction{upRight, downRight, upLeft, downLeft}

	for _, dir := range directions {
		positions = append(positions, iterateDirection(&b, dir, board)...)
	}

	return positions
}

func (b bishop) toString() string {
	return fmt.Sprintf("b%v", b.getTeam())
}

func (b bishop) isEmpty() bool {
	return false
}
