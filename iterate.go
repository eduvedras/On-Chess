package main

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

func addPosition(row int, col int, board board, piece chessPiece) ([]position, bool) {
	positions := []position{}
	curr := board[row][col]

	if !curr.isEmpty() {
		if curr.getTeam() != piece.getTeam() {
			positions = append(positions, position{row, col})
		}
		return positions, true
	}
	positions = append(positions, position{row, col})

	return positions, false
}

func iterateDirection(piece chessPiece, dir direction, board board) []position {
	positions := []position{}

	start := piece.getPos()

	switch dir {
	case up:
		for i := start[0] + 1; i < size; i++ {
			pos, br := addPosition(i, start[1], board, piece)
			positions = append(positions, pos...)
			if br {
				break
			}
		}
	case down:
		for i := start[0] - 1; i >= 0; i-- {
			pos, br := addPosition(i, start[1], board, piece)
			positions = append(positions, pos...)
			if br {
				break
			}
		}
	case right:
		for i := start[1] + 1; i < size; i++ {
			pos, br := addPosition(start[0], i, board, piece)
			positions = append(positions, pos...)
			if br {
				break
			}
		}
	case left:
		for i := start[1] - 1; i >= 0; i-- {
			pos, br := addPosition(start[0], i, board, piece)
			positions = append(positions, pos...)
			if br {
				break
			}
		}
	case upRight:
		i := start[0] + 1
		j := start[1] + 1

		for i < size && j < size {
			pos, br := addPosition(i, j, board, piece)
			positions = append(positions, pos...)
			if br {
				break
			}
			i++
			j++
		}
	case downRight:
		i := start[0] - 1
		j := start[1] + 1

		for i >= 0 && j < size {
			pos, br := addPosition(i, j, board, piece)
			positions = append(positions, pos...)
			if br {
				break
			}
			i--
			j++
		}
	case upLeft:
		i := start[0] + 1
		j := start[1] - 1

		for i < size && j >= 0 {
			pos, br := addPosition(i, j, board, piece)
			positions = append(positions, pos...)
			if br {
				break
			}
			i++
			j--
		}
	case downLeft:
		i := start[0] - 1
		j := start[1] - 1

		for i >= 0 && j >= 0 {
			pos, br := addPosition(i, j, board, piece)
			positions = append(positions, pos...)
			if br {
				break
			}
			i--
			j--
		}
	}

	return positions
}
