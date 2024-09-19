package main

import (
	"bufio"
	"fmt"
	"os"
)

const size = 8

type board [size][size]chessPiece

func isPawn(piece chessPiece) bool {
	switch piece.(type) {
	case *pawn:
		return true
	default:
		return false
	}
}

func isKing(piece chessPiece) bool {
	switch piece.(type) {
	case *king:
		return true
	default:
		return false
	}
}

func handlePromotion(piece chessPiece) chessPiece {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Possible promotions for your pawn > rook knight bishop queen\n")
		fmt.Print("Choose your promotion > ")
		reader.Scan()

		input := reader.Text()

		switch input {
		case "rook":
			piece = &rook{pos: piece.getPos(), team: piece.getTeam()}
		case "knight":
			piece = &knight{pos: piece.getPos(), team: piece.getTeam()}
		case "bishop":
			piece = &bishop{pos: piece.getPos(), team: piece.getTeam()}
		case "queen":
			piece = &queen{pos: piece.getPos(), team: piece.getTeam()}
		default:
			fmt.Println("You need to choose one of the possible promotions!")
			continue
		}
		break
	}
	return piece
}

func updateBoard(board board, piece chessPiece, newPos position) (board, bool) {
	gameEnded := false

	board[piece.getPos()[0]][piece.getPos()[1]] = &empty{}

	if ((piece.getTeam() == white && newPos[0] == size-1) || (piece.getTeam() == black && newPos[0] == 0)) && isPawn(piece) {
		piece = handlePromotion(piece)
	}

	if isKing(board[newPos[0]][newPos[1]]) {
		gameEnded = true
	}

	piece.setPos(position{newPos[0], newPos[1]})

	board[newPos[0]][newPos[1]] = piece

	return board, gameEnded
}

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
				*curr = &pawn{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && (j == 0 || j == 7) {
				*curr = &rook{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && (j == 1 || j == 6) {
				*curr = &knight{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && (j == 2 || j == 5) {
				*curr = &bishop{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && j == 3 {
				*curr = &queen{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else if cond && j == 4 {
				*curr = &king{
					pos:  position{i, j},
					team: initTeam(i),
				}
			} else {
				*curr = &empty{}
			}
		}
	}
	return board
}
