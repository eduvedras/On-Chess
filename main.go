package main

import "fmt"

func main() {
	board := initializeBoard()
	printBoard(board)
	fmt.Println(board[0][0].possibleMoves(board))
	fmt.Println(board[7][0].possibleMoves(board))

	fmt.Println(board[0][7].possibleMoves(board))
	fmt.Println(board[7][7].possibleMoves(board))

}
