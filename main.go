package main

import "fmt"

func main() {
	board := initializeBoard()
	printBoard(board)
	fmt.Println(board[1][0].possibleMoves(board))
	fmt.Println(board[1][3].possibleMoves(board))
	fmt.Println(board[1][7].possibleMoves(board))

	fmt.Println(board[6][0].possibleMoves(board))
	fmt.Println(board[6][3].possibleMoves(board))
	fmt.Println(board[6][7].possibleMoves(board))
}
