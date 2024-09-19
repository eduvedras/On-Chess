package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type inputError struct{}

func (e inputError) Error() string {
	return "Wrong input format. The input should be XX where the first X is a number between 0-7 that represents the row number and the second X is a number between 0-7 that represents the column."
}

func sanitizeInput(input string) (int, int, error) {
	if len(input) != 2 {
		return 0, 0, inputError{}
	}

	inputRune := []rune(input)

	posi := int(input[0] - '0')
	if !unicode.IsDigit(inputRune[0]) || posi >= size {
		return 0, 0, inputError{}
	}

	posj := int(input[1] - '0')
	if !unicode.IsDigit(inputRune[1]) || posj >= size {
		return 0, 0, inputError{}
	}

	return posi, posj, nil
}

func main() {
	board := initializeBoard()
	player := white
	reader := bufio.NewScanner(os.Stdin)

	printBoard(board)

	for {
		fmt.Printf("Player %v turn > ", player)
		reader.Scan()

		input := reader.Text()

		posi, posj, err := sanitizeInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		piece := board[posi][posj]

		if piece.isEmpty() || piece.getTeam() != player {
			fmt.Printf("You need to choose a position with one of the %v pieces!\n", player)
			continue
		}

		possMoves := piece.possibleMoves(board)

		fmt.Print("Possible moves > ")
		for _, move := range possMoves {
			fmt.Printf("%v%v ", move[0], move[1])
		}
		fmt.Println("back")

		fmt.Print("Choose your move > ")

		reader.Scan()

		newposi, newposj, err := sanitizeInput(reader.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		newPos := position{newposi, newposj}
		isValidPos := false
		ended := false
		for _, pos := range possMoves {
			if newPos == pos {
				isValidPos = true
				board, ended = updateBoard(board, piece, newPos)
				//write tests
				//refactor

				if player == white {
					player = black
				} else if player == black {
					player = white
				}
			}
		}

		if !isValidPos {
			fmt.Println("Move is not possible, please choose one of the possible moves!")
			continue
		}

		printBoard(board)

		if ended {
			fmt.Printf("Game ended! Player %v won!!\n", piece.getTeam())
			break
		}
	}

}
