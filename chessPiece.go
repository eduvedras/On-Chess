package main

type team string
type position [2]int

const (
	white team = "w"
	black team = "b"
	emp   team = " "
)

type chessPiece interface {
	getPos() position
	getTeam() team
	possibleMoves(board) []position
	toString() string
	isEmpty() bool
}

type empty struct{}

func (e empty) getTeam() team {
	return emp
}

func (e empty) getPos() position {
	return position{}
}

func (e empty) possibleMoves(_ board) []position {
	return []position{}
}

func (e empty) toString() string {
	return "  "
}

func (e empty) isEmpty() bool {
	return true
}
