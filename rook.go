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
	directions := []direction{right, left, up, down}

	for _, dir := range directions {
		positions = append(positions, iterateDirection(r, dir, board)...)
	}

	return positions
}

func (r rook) toString() string {
	return fmt.Sprintf("r%v", r.getTeam())
}

func (r rook) isEmpty() bool {
	return false
}
