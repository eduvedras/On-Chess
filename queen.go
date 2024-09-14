package main

import "fmt"

type queen struct {
	pos  position
	team team
}

func (q queen) getTeam() team {
	return q.team
}

func (q queen) possibleMoves(board board) []position {
	return []position{{1, 2}, {3, 4}}
}

func (q queen) toString() string {
	return fmt.Sprintf("q%v", q.getTeam())
}

func (q queen) isEmpty() bool {
	return false
}
