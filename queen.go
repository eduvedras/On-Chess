package main

import "fmt"

type queen struct {
	pos  position
	team team
}

func (q queen) getTeam() team {
	return q.team
}

func (q queen) getPos() position {
	return q.pos
}

func (q *queen) setPos(pos position) {
	q.pos = pos
}

func (q queen) possibleMoves(board board) []position {
	positions := []position{}
	directions := []direction{right, left, up, down, upRight, downRight, upLeft, downLeft}

	for _, dir := range directions {
		positions = append(positions, iterateDirection(&q, dir, board)...)
	}

	return positions
}

func (q queen) toString() string {
	return fmt.Sprintf("q%v", q.getTeam())
}

func (q queen) isEmpty() bool {
	return false
}
