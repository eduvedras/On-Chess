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

func (b bishop) possibleMoves(board board) []position {
	positions := []position{}

	positions = append(positions, iterateDirection(b, upRight, board)...)
	positions = append(positions, iterateDirection(b, downRight, board)...)
	positions = append(positions, iterateDirection(b, upLeft, board)...)
	positions = append(positions, iterateDirection(b, downLeft, board)...)

	return positions
}

func (b bishop) toString() string {
	return fmt.Sprintf("b%v", b.getTeam())
}

func (b bishop) isEmpty() bool {
	return false
}
