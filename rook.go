package main

import "fmt"

type rook struct {
	pos  position
	team team
}

type direction string

const (
	right     direction = "right"
	left      direction = "left"
	up        direction = "up"
	down      direction = "down"
	upRight   direction = "upRight"
	downRight direction = "downRight"
	upLeft    direction = "upLeft"
	downLeft  direction = "downLeft"
)

func iterateDirection(piece chessPiece, dir direction, board board) []position {
	positions := []position{}

	start := piece.getPos()

	switch dir {
	case up:
		for i := start[0] + 1; i < size; i++ {
			curr := board[i][start[1]]
			if !curr.isEmpty() {
				if curr.getTeam() != piece.getTeam() {
					positions = append(positions, position{i, start[1]})
				}
				break
			}
			positions = append(positions, position{i, start[1]})
		}
	case down:
		for i := start[0] - 1; i >= 0; i-- {
			curr := board[i][start[1]]
			if !curr.isEmpty() {
				if curr.getTeam() != piece.getTeam() {
					positions = append(positions, position{i, start[1]})
				}
				break
			}
			positions = append(positions, position{i, start[1]})
		}
	case right:
		for i := start[1] + 1; i < size; i++ {
			curr := board[start[0]][i]
			if !curr.isEmpty() {
				if curr.getTeam() != piece.getTeam() {
					positions = append(positions, position{start[0], i})
				}
				break
			}
			positions = append(positions, position{start[0], i})
		}
	case left:
		for i := start[1] - 1; i >= 0; i-- {
			curr := board[start[0]][i]
			if !curr.isEmpty() {
				if curr.getTeam() != piece.getTeam() {
					positions = append(positions, position{start[0], i})
				}
				break
			}
			positions = append(positions, position{start[0], i})
		}
	case upRight:
		i := start[0] + 1
		j := start[1] + 1

		for i < size && j < size {
			curr := board[i][j]
			if !curr.isEmpty() {
				if curr.getTeam() != piece.getTeam() {
					positions = append(positions, position{i, j})
				}
				break
			}
			positions = append(positions, position{i, j})
			i++
			j++
		}
	case downRight:
		i := start[0] - 1
		j := start[1] + 1

		for i >= 0 && j < size {
			curr := board[i][j]
			if !curr.isEmpty() {
				if curr.getTeam() != piece.getTeam() {
					positions = append(positions, position{i, j})
				}
				break
			}
			positions = append(positions, position{i, j})
			i--
			j++
		}
	case upLeft:
		i := start[0] + 1
		j := start[1] - 1

		for i < size && j >= 0 {
			curr := board[i][j]
			if !curr.isEmpty() {
				if curr.getTeam() != piece.getTeam() {
					positions = append(positions, position{i, j})
				}
				break
			}
			positions = append(positions, position{i, j})
			i++
			j--
		}
	case downLeft:
		i := start[0] - 1
		j := start[1] - 1

		for i >= 0 && j >= 0 {
			curr := board[i][j]
			if !curr.isEmpty() {
				if curr.getTeam() != piece.getTeam() {
					positions = append(positions, position{i, j})
				}
				break
			}
			positions = append(positions, position{i, j})
			i--
			j--
		}
	}

	return positions
}

func (r rook) getTeam() team {
	return r.team
}

func (r rook) getPos() position {
	return r.pos
}

func (r rook) possibleMoves(board board) []position {
	positions := []position{}

	positions = append(positions, iterateDirection(r, right, board)...)
	positions = append(positions, iterateDirection(r, left, board)...)
	positions = append(positions, iterateDirection(r, up, board)...)
	positions = append(positions, iterateDirection(r, down, board)...)

	return positions
}

func (r rook) toString() string {
	return fmt.Sprintf("r%v", r.getTeam())
}

func (r rook) isEmpty() bool {
	return false
}
