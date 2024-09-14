package main

import "fmt"

type bishop struct {
	pos  position
	team team
}

func (b bishop) getTeam() team {
	return b.team
}

func (b bishop) possibleMoves(board board) []position {
	return []position{{1, 2}, {3, 4}}
}

func (b bishop) toString() string {
	return fmt.Sprintf("b%v", b.getTeam())
}

func (b bishop) isEmpty() bool {
	return false
}
