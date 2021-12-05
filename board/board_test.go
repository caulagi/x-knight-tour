package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	assert.Equal(
		t,
		NewBoard(2, 0, 0),
		Board(Board{
			Tiles: [][]Tile{
				{
					Tile{Visit: 1, X: 0, Y: 0},
					Tile{Visit: 0, X: 0, Y: 1},
				},
				{
					{Visit: 0, X: 1, Y: 0},
					{Visit: 0, X: 1, Y: 1}},
			},
			PositionX: 0,
			PositionY: 0,
		},
		))
}

func TestHorizontalCandidates(t *testing.T) {
	b := NewBoard(10, 0, 0)
	assert.Equal(
		t,
		b.HorizontalCandidates(5, 5, []Tile{}),
		[]Tile([]Tile{
			{Visit: 0, X: 8, Y: 5},
			{Visit: 0, X: 2, Y: 5},
		}))
}

func TestVerticalCandidates(t *testing.T) {
	b := NewBoard(10, 0, 0)
	assert.Equal(
		t,
		b.VerticalCandidates(5, 5, []Tile{}),
		[]Tile([]Tile{
			{Visit: 0, X: 5, Y: 8},
			{Visit: 0, X: 5, Y: 2},
		}))
}

func TestDiagonalCandidates(t *testing.T) {
	b := NewBoard(10, 0, 0)
	assert.Equal(
		t,
		b.DiagonalCandidates(5, 5, []Tile{}),
		[]Tile([]Tile{
			{Visit: 0, X: 7, Y: 7},
			{Visit: 0, X: 3, Y: 3},
			{Visit: 0, X: 7, Y: 3},
			{Visit: 0, X: 3, Y: 7},
		}))
}
