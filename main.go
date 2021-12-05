package main

import (
	"fmt"
	"math"

	"github.com/caulagi/x-knight-tour/board"
)

func main() {
	startingX, startingY, boardSize := userInput()
	visitCount := 1
	b := board.NewBoard(boardSize, startingX, startingY)

	for {
		candidates := b.CurrentCandidates()
		least := math.MaxInt8
		leastTile := board.Tile{}
		for _, candidate := range candidates {
			nextMoves := b.Candidates(candidate.X, candidate.Y)
			if len(nextMoves) < least {
				least = len(nextMoves)
				leastTile = candidate
			}
		}
		visitCount += 1
		b.Tiles[leastTile.X][leastTile.Y].Visit = visitCount
		b.MoveToTile(leastTile)

		if b.PositionX == startingX && b.PositionY == startingY {
			fmt.Println("Failed to find a path; try again?")
			break
		}

		if visitCount == boardSize*boardSize {
			break
		}
	}
	b.Print()
}
