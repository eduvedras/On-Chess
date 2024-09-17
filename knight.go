package main

import "fmt"

type knight struct {
	pos  position
	team team
}

func (k knight) getTeam() team {
	return k.team
}

func (k knight) getPos() position {
	return k.pos
}

func dirKnightMoves(piece chessPiece, dir direction, board board) []position {
	positions := []position{}
	var conditions []bool
	pos := piece.getPos()

	switch dir {
	case right:
		conditions = []bool{pos[0]+1 < size, pos[1]+2 < size, pos[0]-1 >= 0, pos[1]+2 < size}

		if conditions[0] && conditions[1] {
			posi := pos[0] + 1
			posj := pos[1] + 2

			curr := board[posi][posj]
			if curr.isEmpty() || curr.getTeam() != piece.getTeam() {
				positions = append(positions, position{posi, posj})
			}
		}

		if conditions[2] && conditions[3] {
			posi := pos[0] - 1
			posj := pos[1] + 2

			curr := board[posi][posj]
			if curr.isEmpty() || curr.getTeam() != piece.getTeam() {
				positions = append(positions, position{posi, posj})
			}
		}
	case left:
		conditions = []bool{pos[0]+1 < size, pos[1]-2 >= 0, pos[0]-1 >= 0, pos[1]-2 >= 0}

		if conditions[0] && conditions[1] {
			posi := pos[0] + 1
			posj := pos[1] - 2

			curr := board[posi][posj]
			if curr.isEmpty() || curr.getTeam() != piece.getTeam() {
				positions = append(positions, position{posi, posj})
			}
		}

		if conditions[2] && conditions[3] {
			posi := pos[0] - 1
			posj := pos[1] - 2

			curr := board[posi][posj]
			if curr.isEmpty() || curr.getTeam() != piece.getTeam() {
				positions = append(positions, position{posi, posj})
			}
		}
	case up:
		conditions = []bool{pos[0]+2 < size, pos[1]+1 < size, pos[0]+2 < size, pos[1]-1 >= 0}

		if conditions[0] && conditions[1] {
			posi := pos[0] + 2
			posj := pos[1] + 1

			curr := board[posi][posj]
			if curr.isEmpty() || curr.getTeam() != piece.getTeam() {
				positions = append(positions, position{posi, posj})
			}
		}

		if conditions[2] && conditions[3] {
			posi := pos[0] + 2
			posj := pos[1] - 1

			curr := board[posi][posj]
			if curr.isEmpty() || curr.getTeam() != piece.getTeam() {
				positions = append(positions, position{posi, posj})
			}
		}
	case down:
		conditions = []bool{pos[0]-2 >= 0, pos[1]+1 < size, pos[0]-2 >= 0, pos[1]-1 >= 0}

		if conditions[0] && conditions[1] {
			posi := pos[0] - 2
			posj := pos[1] + 1

			curr := board[posi][posj]
			if curr.isEmpty() || curr.getTeam() != piece.getTeam() {
				positions = append(positions, position{posi, posj})
			}
		}

		if conditions[2] && conditions[3] {
			posi := pos[0] - 2
			posj := pos[1] - 1

			curr := board[posi][posj]
			if curr.isEmpty() || curr.getTeam() != piece.getTeam() {
				positions = append(positions, position{posi, posj})
			}
		}
	}

	return positions
}

func (k knight) possibleMoves(board board) []position {
	positions := []position{}
	directions := []direction{right, left, up, down}

	for _, dir := range directions {
		positions = append(positions, dirKnightMoves(k, dir, board)...)
	}

	return positions
}

func (k knight) toString() string {
	return fmt.Sprintf("n%v", k.getTeam())
}

func (k knight) isEmpty() bool {
	return false
}
