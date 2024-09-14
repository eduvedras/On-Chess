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
	positions := []position{}
	initLine := 0
	incrementer := 0

	if p.getTeam() == white {
		initLine = 1
		incrementer = +1

	} else {
		initLine = 6
		incrementer = -1
	}

	fronti := position{p.pos[0] + incrementer, p.pos[1]}
	front2i := position{p.pos[0] + incrementer*2, p.pos[1]}
	ldiag := position{p.pos[0] + incrementer, p.pos[1] - 1}
	rdiag := position{p.pos[0] + incrementer, p.pos[1] + 1}

	if p.pos[0] == initLine && board[front2i[0]][front2i[1]].isEmpty() {
		positions = append(positions, front2i)
	}

	if board[fronti[0]][fronti[1]].isEmpty() {
		positions = append(positions, fronti)
	}

	if ldiag[1] >= 0 && !board[ldiag[0]][ldiag[1]].isEmpty() {
		positions = append(positions, ldiag)
	}

	if rdiag[1] <= 7 && !board[rdiag[0]][rdiag[1]].isEmpty() {
		positions = append(positions, rdiag)
	}

	return positions
}

func (p pawn) toString() string {
	return fmt.Sprintf("p%v", p.getTeam())
}

func (p pawn) isEmpty() bool {
	return false
}
