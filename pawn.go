package main

import "fmt"

type pawn struct {
	pos  position
	team team
}

func (p pawn) getTeam() team {
	return p.team
}

func (p pawn) possibleMoves(board board) []position {
	return []position{{1, 2}, {3, 4}}
}

func (p pawn) toString() string {
	return fmt.Sprintf("p%v", p.getTeam())
}
