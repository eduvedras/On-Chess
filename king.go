package main

import "fmt"

type king struct {
	pos  position
	team team
}

func (k king) getTeam() team {
	return k.team
}

func (k king) possibleMoves(board board) []position {
	return []position{{1, 2}, {3, 4}}
}

func (k king) toString() string {
	return fmt.Sprintf("k%v", k.getTeam())
}

func (k king) isEmpty() bool {
	return false
}
