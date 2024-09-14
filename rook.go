package main

import "fmt"

type rook struct {
	pos  position
	team team
}

func (r rook) getTeam() team {
	return r.team
}

func (r rook) possibleMoves(board board) []position {
	return []position{{1, 2}, {3, 4}}
}

func (r rook) toString() string {
	return fmt.Sprintf("r%v", r.getTeam())
}

func (r rook) isEmpty() bool {
	return false
}
