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

func TestSolve(t *testing.T) {
	b := NewBoard(10, 0, 0)
	b.Solve(0, 0, 10)
	assert.Equal(
		t,
		b.GetTiles(),
		[][]Tile{
			{
				{Visit: 1, X: 0, Y: 0},
				{Visit: 62, X: 0, Y: 1},
				{Visit: 40, X: 0, Y: 2},
				{Visit: 16, X: 0, Y: 3},
				{Visit: 61, X: 0, Y: 4},
				{Visit: 39, X: 0, Y: 5},
				{Visit: 75, X: 0, Y: 6},
				{Visit: 49, X: 0, Y: 7},
				{Visit: 38, X: 0, Y: 8},
				{Visit: 74, X: 0, Y: 9},
			},
			{
				{Visit: 54, X: 1, Y: 0},
				{Visit: 18, X: 1, Y: 1},
				{Visit: 3, X: 1, Y: 2},
				{Visit: 53, X: 1, Y: 3},
				{Visit: 58, X: 1, Y: 4},
				{Visit: 4, X: 1, Y: 5},
				{Visit: 52, X: 1, Y: 6},
				{Visit: 72, X: 1, Y: 7},
				{Visit: 5, X: 1, Y: 8},
				{Visit: 51, X: 1, Y: 9},
			},
			{
				{Visit: 41, X: 2, Y: 0},
				{Visit: 15, X: 2, Y: 1},
				{Visit: 60, X: 2, Y: 2},
				{Visit: 67, X: 2, Y: 3},
				{Visit: 94, X: 2, Y: 4},
				{Visit: 79, X: 2, Y: 5},
				{Visit: 66, X: 2, Y: 6},
				{Visit: 95, X: 2, Y: 7},
				{Visit: 78, X: 2, Y: 8},
				{Visit: 48, X: 2, Y: 9},
			},
			{
				{Visit: 2, X: 3, Y: 0},
				{Visit: 63, X: 3, Y: 1},
				{Visit: 57, X: 3, Y: 2},
				{Visit: 17, X: 3, Y: 3},
				{Visit: 64, X: 3, Y: 4},
				{Visit: 69, X: 3, Y: 5},
				{Visit: 76, X: 3, Y: 6},
				{Visit: 50, X: 3, Y: 7},
				{Visit: 37, X: 3, Y: 8},
				{Visit: 73, X: 3, Y: 9},
			},
			{
				{Visit: 55, X: 4, Y: 0},
				{Visit: 19, X: 4, Y: 1},
				{Visit: 91, X: 4, Y: 2},
				{Visit: 82, X: 4, Y: 3},
				{Visit: 59, X: 4, Y: 4},
				{Visit: 90, X: 4, Y: 5},
				{Visit: 99, X: 4, Y: 6},
				{Visit: 71, X: 4, Y: 7},
				{Visit: 6, X: 4, Y: 8},
				{Visit: 98, X: 4, Y: 9},
			},
			{
				{Visit: 42, X: 5, Y: 0},
				{Visit: 14, X: 5, Y: 1},
				{Visit: 29, X: 5, Y: 2},
				{Visit: 68, X: 5, Y: 3},
				{Visit: 93, X: 5, Y: 4},
				{Visit: 80, X: 5, Y: 5},
				{Visit: 65, X: 5, Y: 6},
				{Visit: 96, X: 5, Y: 7},
				{Visit: 77, X: 5, Y: 8},
				{Visit: 47, X: 5, Y: 9},
			},
			{
				{Visit: 23, X: 6, Y: 0},
				{Visit: 83, X: 6, Y: 1},
				{Visit: 56, X: 6, Y: 2},
				{Visit: 24, X: 6, Y: 3},
				{Visit: 86, X: 6, Y: 4},
				{Visit: 70, X: 6, Y: 5},
				{Visit: 25, X: 6, Y: 6},
				{Visit: 87, X: 6, Y: 7},
				{Visit: 36, X: 6, Y: 8},
				{Visit: 10, X: 6, Y: 9},
			},
			{
				{Visit: 30, X: 7, Y: 0},
				{Visit: 20, X: 7, Y: 1},
				{Visit: 92, X: 7, Y: 2},
				{Visit: 81, X: 7, Y: 3},
				{Visit: 32, X: 7, Y: 4},
				{Visit: 89, X: 7, Y: 5},
				{Visit: 100, X: 7, Y: 6},
				{Visit: 33, X: 7, Y: 7},
				{Visit: 7, X: 7, Y: 8},
				{Visit: 97, X: 7, Y: 9},
			},
			{
				{Visit: 43, X: 8, Y: 0},
				{Visit: 13, X: 8, Y: 1},
				{Visit: 28, X: 8, Y: 2},
				{Visit: 44, X: 8, Y: 3},
				{Visit: 12, X: 8, Y: 4},
				{Visit: 27, X: 8, Y: 5},
				{Visit: 45, X: 8, Y: 6},
				{Visit: 11, X: 8, Y: 7},
				{Visit: 26, X: 8, Y: 8},
				{Visit: 46, X: 8, Y: 9},
			},
			{
				{Visit: 22, X: 9, Y: 0},
				{Visit: 84, X: 9, Y: 1},
				{Visit: 31, X: 9, Y: 2},
				{Visit: 21, X: 9, Y: 3},
				{Visit: 85, X: 9, Y: 4},
				{Visit: 34, X: 9, Y: 5},
				{Visit: 8, X: 9, Y: 6},
				{Visit: 88, X: 9, Y: 7},
				{Visit: 35, X: 9, Y: 8},
				{Visit: 9, X: 9, Y: 9},
			}})
}
