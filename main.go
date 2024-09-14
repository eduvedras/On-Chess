package main

import "fmt"

type chessPiece string

const (
	Pawn   chessPiece = "p"
	Bishop chessPiece = "b"
	Knight chessPiece = "n"
	Rook   chessPiece = "r"
	Queen  chessPiece = "q"
	King   chessPiece = "k"
	Empty  chessPiece = " "
)

const size = 8

type board [8][8]chessPiece

func main() {
	board := initializeBoard()
	printBoard(board)
}

func printBoard(board board) {
	fmt.Print("\n --- --- --- --- --- --- --- ---\n")
	for i := size - 1; i >= 0; i-- {
		for j := 0; j < size; j++ {
			curr := board[i][j]
			if j == size-1 {
				fmt.Printf("| %v |", curr)
			} else {
				fmt.Printf("| %v ", curr)
			}
		}
		fmt.Print("\n --- --- --- --- --- --- --- ---\n")
	}
}

func initializeBoard() board {
	board := board{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			curr := &board[i][j]

			cond := i == 0 || i == 7
			if i == 1 || i == 6 {
				*curr = Pawn
			} else if cond && (j == 0 || j == 7) {
				*curr = Rook
			} else if cond && (j == 1 || j == 6) {
				*curr = Knight
			} else if cond && (j == 2 || j == 5) {
				*curr = Bishop
			} else if cond && j == 3 {
				*curr = Queen
			} else if cond && j == 4 {
				*curr = King
			} else {
				*curr = Empty
			}
		}
	}
	return board
}
