package board

import "fmt"

type Tile struct {
	Visit int
	X     int
	Y     int
}

type Board struct {
	Tiles     [][]Tile
	PositionX int
	PositionY int
}

func NewBoard(size, startX, startY int) Board {
	b := Board{PositionX: startX, PositionY: startY}
	for i := 0; i < size; i++ {
		b.Tiles = append(b.Tiles, []Tile{})
		for j := 0; j < size; j++ {
			b.Tiles[i] = append(b.Tiles[i], Tile{Visit: 0, X: i, Y: j})
		}
	}
	b.Tiles[startX][startY].Visit = 1
	return b
}

func (b *Board) CurrentCandidates() []Tile {
	return b.Candidates(b.PositionX, b.PositionY)
}

func (b *Board) horizontalCandidates(x, y int, potentialMoves []Tile) []Tile {
	if x+3 <= len(b.Tiles)-1 {
		potentialMoves = append(potentialMoves, b.Tiles[x+3][y])
	}
	if x-3 >= 0 {
		potentialMoves = append(potentialMoves, b.Tiles[x-3][y])
	}
	return potentialMoves
}

func (b *Board) verticalCandidates(x, y int, potentialMoves []Tile) []Tile {
	if y+3 <= len(b.Tiles)-1 {
		potentialMoves = append(potentialMoves, b.Tiles[x][y+3])
	}
	if y-3 >= 0 {
		potentialMoves = append(potentialMoves, b.Tiles[x][y-3])
	}
	return potentialMoves
}

func (b *Board) DiagonalCandidates(x, y int, potentialMoves []Tile) []Tile {
	if x+2 <= len(b.Tiles)-1 && y+2 <= len(b.Tiles)-1 {
		potentialMoves = append(potentialMoves, b.Tiles[x+2][y+2])
	}
	if x-2 >= 0 && y-2 >= 0 {
		potentialMoves = append(potentialMoves, b.Tiles[x-2][y-2])
	}
	if x+2 <= len(b.Tiles)-1 && y-2 >= 0 {
		potentialMoves = append(potentialMoves, b.Tiles[x+2][y-2])
	}
	if x-2 >= 0 && y+2 <= len(b.Tiles)-1 {
		potentialMoves = append(potentialMoves, b.Tiles[x-2][y+2])
	}
	return potentialMoves
}

func (b *Board) Candidates(x, y int) []Tile {

	potentialMoves := []Tile{}

	if x >= len(b.Tiles) || y >= len(b.Tiles) {
		return potentialMoves
	}

	potentialMoves = b.horizontalCandidates(x, y, potentialMoves)
	potentialMoves = b.verticalCandidates(x, y, potentialMoves)
	potentialMoves = b.DiagonalCandidates(x, y, potentialMoves)

	moves := []Tile{}
	//Remove all potential moves that have been visited before
	for _, move := range potentialMoves {
		if move.Visit == 0 {
			moves = append(moves, move)
		}
	}
	return moves
}

func (b *Board) Print() {
	for i, row := range b.Tiles {
		for j := range row {
			fmt.Printf(" %3d", b.Tiles[i][j].Visit)
		}
		fmt.Println()
	}
}

func (b *Board) MoveToTile(tile Tile) {
	tile.Visit = 1
	b.PositionX = tile.X
	b.PositionY = tile.Y
}
