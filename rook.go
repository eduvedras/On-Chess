package main

import "fmt"

type rook struct {
	pos  position
	team team
}

type direction string

const (
	right direction = "right"
	left  direction = "left"
	up    direction = "up"
	down  direction = "down"
)

func iterateRowOrCol(start position, dir direction, board board) []position {
	positions := []position{}

	switch dir {
	case up:
		for i := start[0] + 1; i < size; i++ {
			if !board[i][start[1]].isEmpty() {
				break
			}
			positions = append(positions, position{i, start[1]})
		}
	case down:
		for i := start[0] - 1; i >= 0; i-- {
			if !board[i][start[1]].isEmpty() {
				break
			}
			positions = append(positions, position{i, start[1]})
		}
	case right:
		for i := start[1] + 1; i < size; i++ {
			if !board[start[0]][i].isEmpty() {
				break
			}
			positions = append(positions, position{start[0], i})
		}
	case left:
		for i := start[1] - 1; i >= 0; i-- {
			if !board[start[0]][i].isEmpty() {
				break
			}
			positions = append(positions, position{start[0], i})
		}
	}

	return positions
}

func (r rook) getTeam() team {
	return r.team
}

func (r rook) possibleMoves(board board) []position {
	positions := []position{}

	positions = append(positions, iterateRowOrCol(r.pos, right, board)...)
	positions = append(positions, iterateRowOrCol(r.pos, left, board)...)
	positions = append(positions, iterateRowOrCol(r.pos, up, board)...)
	positions = append(positions, iterateRowOrCol(r.pos, down, board)...)

	return positions
}

func (r rook) toString() string {
	return fmt.Sprintf("r%v", r.getTeam())
}

func (r rook) isEmpty() bool {
	return false
}
