package main

import "fmt"

const size = 8

type board [size][size]chessPiece

func printBoard(board board) {
	fmt.Print("\n ---- ---- ---- ---- ---- ---- ---- ----\n")
	for i := size - 1; i >= 0; i-- {
		for j := 0; j < size; j++ {
			curr := board[i][j]
			if j == size-1 {
				fmt.Printf("| %v |", curr.toString())
			} else {
				fmt.Printf("| %v ", curr.toString())
			}
		}
		fmt.Print("\n ---- ---- ---- ---- ---- ---- ---- ----\n")
	}
}

func initTeam(i int) team {
	if i == 0 || i == 1 {
		return white
	}
	return black
}

func initializeBoard() board {
	board := board{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			curr := &board[i][j]

			cond := i == 0 || i == 7
			if i == 1 || i == 6 {
				*curr = pawn{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && (j == 0 || j == 7) {
				*curr = rook{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && (j == 1 || j == 6) {
				*curr = knight{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && (j == 2 || j == 5) {
				*curr = bishop{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && j == 3 {
				*curr = queen{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && j == 4 {
				*curr = king{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else {
				*curr = empty{}
			}
		}
	}
	return board
}
