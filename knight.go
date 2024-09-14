package main

import "fmt"

type knight struct {
	pos  position
	team team
}

func (k knight) getTeam() team {
	return k.team
}

func (k knight) possibleMoves(board board) []position {
	return []position{{1, 2}, {3, 4}}
}

func (k knight) toString() string {
	return fmt.Sprintf("n%v", k.getTeam())
}
