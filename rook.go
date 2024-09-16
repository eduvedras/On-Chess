package main

import "fmt"

type rook struct {
	pos  position
	team team
}

func (r rook) getTeam() team {
	return r.team
}

func (r rook) getPos() position {
	return r.pos
}

func (r rook) possibleMoves(board board) []position {
	positions := []position{}

	positions = append(positions, iterateDirection(r, right, board)...)
	positions = append(positions, iterateDirection(r, left, board)...)
	positions = append(positions, iterateDirection(r, up, board)...)
	positions = append(positions, iterateDirection(r, down, board)...)

	return positions
}

func (r rook) toString() string {
	return fmt.Sprintf("r%v", r.getTeam())
}

func (r rook) isEmpty() bool {
	return false
}
