package main

import (
	"github.com/caulagi/x-knight-tour/board"
)

func main() {
	startingX, startingY, boardSize := userInput()
	b := board.NewBoard(boardSize, startingX, startingY)
	b.Solve(startingX, startingY, boardSize)
	b.Print()
}
