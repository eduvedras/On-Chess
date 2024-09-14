package main

import "fmt"

func main() {
	board := initializeBoard()
	printBoard(board)
	fmt.Println(board[0][0].getTeam())
	fmt.Println(board[0][0].possibleMoves(board))
}
